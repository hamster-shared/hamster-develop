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
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"log"
	"os"
	"path"
	"strconv"
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

	_, workflowId, err := GetProjectIdAndWorkflowIdByWorkflowKey(message.JobName)
	if err != nil {
		return
	}

	jobDetail := w.engine.GetJobHistory(message.JobName, message.JobId)

	var workflowDetail db.WorkflowDetail

	w.db.Where(&db.WorkflowDetail{
		WorkflowId: uint(workflowId),
		ExecNumber: uint(jobDetail.Id),
	}).First(&workflowDetail)

	workflowDetail.Status = uint(jobDetail.Status)
	stageInfo, err := json.Marshal(jobDetail.Stages)
	if err != nil {
		return
	}
	workflowDetail.StageInfo = string(stageInfo)
	workflowDetail.UpdateTime = time.Now()
	workflowDetail.CodeInfo = w.engine.GetCodeInfo(message.JobName, message.JobId)

	tx := w.db.Save(workflowDetail)
	tx.Commit()

	w.SyncContract(message, workflowDetail)
	w.SyncReport(message)
}

func (w *WorkflowService) SyncContract(message model.StatusChangeMessage, workflowDetail db.WorkflowDetail) {
	projectId, workflowId, err := GetProjectIdAndWorkflowIdByWorkflowKey(message.JobName)
	if err != nil {
		return
	}
	jobDetail := w.engine.GetJobHistory(message.JobName, message.JobId)

	if len(jobDetail.Artifactorys) > 0 {

		for _, arti := range jobDetail.Artifactorys {

			data, _ := os.ReadFile(arti.Url)
			m := make(map[string]any)

			err := json.Unmarshal(data, &m)
			if err != nil {
				continue
			}
			abi, err := json.Marshal(m["abi"])
			bytecodeData, ok := m["bytecode"]
			if !ok {
				continue
			}

			contract := db.Contract{
				ProjectId:        projectId,
				WorkflowId:       workflowId,
				WorkflowDetailId: workflowDetail.Id,
				Name:             strings.TrimSuffix(arti.Name, path.Ext(arti.Name)),
				Version:          fmt.Sprintf("#%d", workflowDetail.ExecNumber),
				Network:          "",
				BuildTime:        workflowDetail.CreateTime,
				AbiInfo:          string(abi),
				ByteCode:         bytecodeData.(string),
				CreateTime:       time.Now(),
			}
			err = w.db.Save(&contract).Error
			fmt.Println(err)
		}

	}
}

func (w *WorkflowService) SyncReport(message model.StatusChangeMessage) {
	if !strings.Contains(message.JobName, "_") {
		return
	}
	projectId, workflowId, err := GetProjectIdAndWorkflowIdByWorkflowKey(message.JobName)
	if err != nil {
		return
	}
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

	workflowKey := w.GetWorkflowKey(projectId, workflow.Id)

	job := w.engine.GetJob(workflowKey)
	if job == nil {
		var jobModel model.Job
		err := yaml.Unmarshal([]byte((workflow.ExecFile)), &jobModel)
		if jobModel.Name != workflowKey {
			jobModel.Name = workflowKey
			execFile, _ := yaml.Marshal(jobModel)
			workflow.ExecFile = string(execFile)
		}

		err = w.engine.CreateJob(workflowKey, workflow.ExecFile)
		if err != nil {
			return err
		}
		job = w.engine.GetJob(workflowKey)
	}

	detail, err := w.engine.ExecuteJob(workflowKey)

	if err != nil {
		return err
	}
	stageInfo, err := json.Marshal(detail.Stages)
	if err != nil {
		return err
	}

	dbDetail := db.WorkflowDetail{
		Type:        workflowType,
		ProjectId:   projectId,
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
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
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

func (w *WorkflowService) GetWorkflowDetail(workflowId, workflowDetailId int) (*vo.WorkflowDetailVo, error) {
	var workflowDetail db2.WorkflowDetail
	var detail vo.WorkflowDetailVo
	res := w.db.Model(db2.WorkflowDetail{}).Where("workflow_id = ? and id = ?", workflowId, workflowDetailId).First(&workflowDetail)
	if res.Error != nil {
		return &detail, res.Error
	}
	copier.Copy(&detail, &workflowDetail)
	return &detail, nil
}

func (w *WorkflowService) QueryWorkflowDetail(workflowId, workflowDetailId int) (*db2.WorkflowDetail, error) {
	var workflowDetail db2.WorkflowDetail
	res := w.db.Model(db2.WorkflowDetail{}).Where("workflow_id = ? and id = ?", workflowId, workflowDetailId).First(&workflowDetail)
	if res.Error != nil {
		return &workflowDetail, res.Error
	}
	return &workflowDetail, nil
}

func (w *WorkflowService) QueryWorkflow(workflowId int) (*db2.Workflow, error) {
	var workflow db2.Workflow
	res := w.db.Model(db2.Workflow{}).Where("id = ?", workflowId).First(&workflow)
	if res.Error != nil {
		return &workflow, res.Error
	}
	return &workflow, nil
}

func (w *WorkflowService) GetWorkflowKey(projectId uint, workflowId uint) string {
	return fmt.Sprintf("%d_%d", projectId, workflowId)
}

func GetProjectIdAndWorkflowIdByWorkflowKey(projectKey string) (uint, uint, error) {
	projectId, err := strconv.Atoi(strings.Split(projectKey, "_")[0])
	if err != nil {
		return 0, 0, err
	}
	workflowId, err := strconv.Atoi(strings.Split(projectKey, "_")[1])
	return uint(projectId), uint(workflowId), err

}

func (w *WorkflowService) SaveWorkflow(saveData parameter.SaveWorkflowParam) (db2.Workflow, error) {
	var workflow db2.Workflow
	workflow.Type = uint(saveData.Type)
	workflow.CreateTime = time.Now()
	workflow.UpdateTime = time.Now()
	workflow.ProjectId = saveData.ProjectId
	workflow.ExecFile = saveData.ExecFile
	workflow.LastExecId = saveData.LastExecId
	res := w.db.Save(&workflow)
	if res.Error != nil {
		return workflow, res.Error
	}
	return workflow, nil
}

func (w *WorkflowService) UpdateWorkflow(data db2.Workflow) error {
	res := w.db.Save(&data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (w *WorkflowService) TemplateParse(name, url string, workflowType consts.WorkflowType) (string, error) {
	filePath := "templates/truffle-build.yml"
	if workflowType == consts.Check {
		filePath = "templates/truffle_check.yml"
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
