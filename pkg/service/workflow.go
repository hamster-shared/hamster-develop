package service

import (
	"bytes"
	"database/sql"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hamster-shared/a-line/engine"
	"github.com/hamster-shared/a-line/engine/model"
	"github.com/hamster-shared/a-line/pkg/application"
	"github.com/hamster-shared/a-line/pkg/consts"
	"github.com/hamster-shared/a-line/pkg/db"
	db2 "github.com/hamster-shared/a-line/pkg/db"
	"github.com/hamster-shared/a-line/pkg/parameter"
	"github.com/hamster-shared/a-line/pkg/vo"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"log"
	"strings"
	"text/template"
	"time"
)

//go:embed templates
var temp embed.FS

type WorkflowService struct {
	db     *gorm.DB
	engine *engine.Engine
}

func NewWorkflowService() *WorkflowService {
	workflowService := &WorkflowService{
		db:     application.GetBean[*gorm.DB]("db"),
		engine: application.GetBean[*engine.Engine]("engine"),
	}

	go workflowService.engine.RegisterStatusChangeHook(workflowService.SyncStatus)

	return workflowService
}

func (w *WorkflowService) SyncStatus(message model.StatusChangeMessage) {
	w.SyncContract(message)
	w.SyncReport(message)
}

func (w *WorkflowService) SyncContract(message model.StatusChangeMessage) {

}

func (w *WorkflowService) SyncReport(message model.StatusChangeMessage) {
	if !strings.Contains(message.JobName, "_") {
		return
	}
	projectId := strings.Split(message.JobName, "_")[0]
	workflowId := strings.Split(message.JobName, "_")[1]
	workflowExecNumber := message.JobId

	if message.Status == model.STATUS_SUCCESS {
		//TODO.... 实现同步报告
		fmt.Println(projectId, workflowId, workflowExecNumber)
	}

}

func (w *WorkflowService) ExecProjectCheckWorkflow(projectId uint, user vo.UserAuth) error {
	return w.ExecProjectWorkflow(projectId, user, 1)
}

func (w *WorkflowService) ExecProjectBuildWorkflow(projectId uint, user vo.UserAuth) error {
	return w.ExecProjectWorkflow(projectId, user, 2)
}

func (w *WorkflowService) ExecProjectWorkflow(projectId uint, user vo.UserAuth, workflowType uint) error {

	// query project workflow

	var workflow db.Workflow

	w.db.Where(&db.Workflow{
		ProjectId: projectId,
		Type:      workflowType,
	}).First(&workflow)

	if &workflow == nil {
		return errors.New("no check workflow in the project ")
	}

	workflowKey := GetWorkflowKey(projectId, workflow.Id)

	detail, err := w.engine.ExecuteJob(workflowKey)

	if err != nil {
		return err
	}
	stageInfo, err := json.Marshal(detail.Stages)
	if err != nil {
		return err
	}

	dbDetail := db.WorkflowDetail{
		WorkflowId:  workflow.Id,
		ExecNumber:  uint(detail.Id),
		StageInfo:   string(stageInfo),
		TriggerUser: user.Username,
		TriggerMode: 1,
		CodeInfo:    "",
		Status:      uint(detail.Status),
		StartTime: sql.NullTime{
			Time:  detail.StartTime,
			Valid: true,
		},
	}

	err = w.db.Transaction(func(tx *gorm.DB) error {
		tx.Save(&dbDetail)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (w *WorkflowService) GetWorkflowList(projectId, workflowType, page, size int) (*vo.WorkflowPage, error) {
	var total int64
	var data vo.WorkflowPage
	var workflowData []vo.WorkflowVo
	var workflowList []db2.Workflow
	tx := w.db.Model(db2.Workflow{}).Where("project_id = ?", projectId)
	if workflowType != 0 {
		tx = tx.Where("type = ? ", workflowType)
	}
	result := tx.Offset((page - 1) * size).Limit(size).Find(&workflowList).Count(&total)
	if result.Error != nil {
		return &data, result.Error
	}
	copier.Copy(&workflowData, &workflowList)
	if len(workflowData) > 0 {
		for _, datum := range workflowData {
			var detailData db2.WorkflowDetail
			res := w.db.Model(db2.WorkflowDetail{}).Where("workflow_id = ? and id = ?", datum.Id, datum.LastExecId).First(&detailData)
			if res.Error != nil {
				copier.Copy(&datum, &detailData)
			}
		}
	}
	data.Data = workflowData
	data.Total = int(total)
	data.Page = page
	data.PageSize = size
	return &data, nil
}

func (w *WorkflowService) GetWorkflowDetail(workflowId, workflowDetailId int) (*db2.WorkflowDetail, error) {
	var workflowDetail db2.WorkflowDetail
	res := w.db.Model(db2.WorkflowDetail{}).Where("workflow_id = ? and id = ?", workflowId, workflowDetailId).First(&workflowDetail)
	if res.Error != nil {
		return &workflowDetail, res.Error
	}
	return &workflowDetail, nil
}

func GetWorkflowKey(projectId uint, workflowId uint) string {
	return fmt.Sprintf("%d_%d", projectId, workflowId)
}

func (w *WorkflowService) SaveWorkflow(saveData parameter.SaveWorkflowParam) (uint, error) {
	var workflow db2.Workflow
	workflow.Type = uint(saveData.Type)
	workflow.CreateTime = time.Now()
	workflow.ProjectId = saveData.ProjectId
	workflow.ExecFile = saveData.ExecFile
	workflow.LastExecId = saveData.LastExecId
	res := w.db.Create(&workflow)
	if res.Error != nil {
		return 0, res.Error
	}
	return workflow.Id, nil
}

func (w *WorkflowService) TemplateParse(name, url string, workflowType consts.WorkflowType) (string, error) {
	filePath := "templates/truffle_check.yml"
	if workflowType == consts.Check {
		filePath = "templates/truffle-build.yml"
	}
	content, err := temp.ReadFile(filePath)
	if err != nil {
		log.Println("read template file failed ", err.Error())
		return "", err
	}
	fileContent := string(content)
	tmpl, err := template.New("test").Parse(fileContent)
	if err != nil {
		log.Println("template parse failed ", err.Error())
		return "", err
	}
	templateData := parameter.TemplateCheck{
		Name:          name,
		RepositoryUrl: url,
	}
	var input bytes.Buffer
	err = tmpl.Execute(&input, templateData)
	if err != nil {
		log.Println("failed to write parameters to the template ", err)
		return "", err
	}
	return input.String(), nil
}
