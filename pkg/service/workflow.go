package service

import (
	"bytes"
	"database/sql"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"text/template"
	"time"

	engine "github.com/hamster-shared/aline-engine"
	jober "github.com/hamster-shared/aline-engine/job"
	"github.com/hamster-shared/aline-engine/logger"
	"github.com/hamster-shared/aline-engine/model"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	"github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/jinzhu/copier"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

//go:embed templates
var temp embed.FS

type WorkflowService struct {
	db     *gorm.DB
	engine engine.Engine
}

func NewWorkflowService() *WorkflowService {
	workflowService := &WorkflowService{
		db:     application.GetBean[*gorm.DB]("db"),
		engine: application.GetBean[engine.Engine]("engine"),
	}

	go workflowService.engine.RegisterStatusChangeHook(workflowService.SyncStatus)

	return workflowService
}

func (w *WorkflowService) SyncStatus(message model.StatusChangeMessage) {
	logger.Debugf("SyncStatus: %v", message)

	_, workflowId, err := GetProjectIdAndWorkflowIdByWorkflowKey(message.JobName)
	if err != nil {
		return
	}

	jobDetail, err := w.engine.GetJobHistory(message.JobName, message.JobId)
	if err != nil {
		return
	}

	var workflowDetail db.WorkflowDetail

	w.db.Where(&db.WorkflowDetail{
		WorkflowId: uint(workflowId),
		ExecNumber: uint(jobDetail.Id),
	}).First(&workflowDetail)

	if workflowDetail.Id == 0 {
		return
	}

	workflowDetail.Status = uint(jobDetail.Status)
	stageInfo, err := json.Marshal(jobDetail.Stages)
	if err != nil {
		return
	}
	workflowDetail.StageInfo = string(stageInfo)
	workflowDetail.UpdateTime = time.Now()
	workflowDetail.CodeInfo, err = w.engine.GetCodeInfo(message.JobName, message.JobId)
	if err != nil {
		logger.Warnf("get code info failed: %v", err)
	}
	workflowDetail.Duration = jobDetail.Duration

	tx := w.db.Save(&workflowDetail)
	tx.Commit()

	w.SyncContract(message, workflowDetail)
	w.SyncReport(message, workflowDetail)
	w.SyncFrontendPackage(message, workflowDetail)
}

func (w *WorkflowService) SyncFrontendPackage(message model.StatusChangeMessage, workflowDetail db.WorkflowDetail) {
	projectIdStr, _, err := GetProjectIdAndWorkflowIdByWorkflowKey(message.JobName)
	if err != nil {
		return
	}
	projectId, err := uuid.FromString(projectIdStr)
	if err != nil {
		log.Println("UUID from string failed: ", err.Error())
		return
	}
	var projectData db.Project
	err = w.db.Model(db.Project{}).Where("id = ?", projectId).First(&projectData).Error
	if err != nil {
		log.Println("find project by id failed: ", err.Error())
		return
	}

	if uint(consts.FRONTEND) != projectData.Type {
		return
	}
	jobDetail, err := w.engine.GetJobHistory(message.JobName, message.JobId)
	if err != nil {
		return
	}
	if uint(consts.Build) == workflowDetail.Type {
		w.syncFrontendBuild(jobDetail, workflowDetail, projectData)
	} else if uint(consts.Deploy) == workflowDetail.Type {
		w.syncFrontendDeploy(jobDetail, workflowDetail, projectData)
	}
}

func (w *WorkflowService) syncFrontendBuild(detail *model.JobDetail, workflowDetail db.WorkflowDetail, project db.Project) {
	if len(detail.ActionResult.Artifactorys) > 0 {
		for range detail.ActionResult.Artifactorys {
			frontendPackage := db.FrontendPackage{
				ProjectId:        workflowDetail.ProjectId,
				WorkflowId:       workflowDetail.WorkflowId,
				WorkflowDetailId: workflowDetail.Id,
				Name:             project.Name,
				Version:          fmt.Sprintf("%d", workflowDetail.ExecNumber),
				Branch:           workflowDetail.CodeInfo,
				BuildTime:        workflowDetail.CreateTime,
				CreateTime:       time.Now(),
			}
			err := w.db.Save(&frontendPackage).Error
			if err != nil {
				log.Println("save frontend package failed: ", err.Error())
			}
		}
	}
}

func (w *WorkflowService) syncFrontendDeploy(detail *model.JobDetail, workflowDetail db.WorkflowDetail, project db.Project) {

	if len(detail.ActionResult.Deploys) > 0 {
		buildWorkflowDetailIdStr := detail.Parameter["buildWorkflowDetailId"]
		if buildWorkflowDetailIdStr == "" {
			return
		}
		buildWorkflowDetailId, err := strconv.Atoi(buildWorkflowDetailIdStr)
		if err != nil {
			return
		}
		var image string
		if project.FrameType == 1 {
			image = "https://develop-images.api.hamsternet.io/vue.png"
		} else {
			image = "https://develop-images.api.hamsternet.io/react.png"
		}
		for _, deploy := range detail.ActionResult.Deploys {
			var data db.FrontendPackage
			err := w.db.Model(db.FrontendPackage{}).Where("workflow_detail_id = ?", buildWorkflowDetailId).First(&data).Error
			if err == nil {
				data.Domain = deploy.Url
				err := w.db.Save(&data).Error
				if err != nil {
					log.Println("save frontend package failed: ", err.Error())
				}
				var packageDeploy db.FrontendDeploy
				packageDeploy.ProjectId = project.Id
				packageDeploy.WorkflowId = workflowDetail.WorkflowId
				packageDeploy.WorkflowDetailId = workflowDetail.Id
				packageDeploy.PackageId = data.Id
				packageDeploy.DeployInfo = deploy.Cid
				packageDeploy.Domain = deploy.Url
				packageDeploy.Version = data.Version
				packageDeploy.DeployTime = sql.NullTime{Time: time.Now(), Valid: true}
				packageDeploy.Name = project.Name
				packageDeploy.Branch = data.Branch
				packageDeploy.CreateTime = time.Now()
				packageDeploy.Image = image
				err = w.db.Save(&packageDeploy).Error
				if err != nil {
					log.Println("save frontend deploy failed: ", err.Error())
				}

			}
		}
	}
}

func (w *WorkflowService) SyncContract(message model.StatusChangeMessage, workflowDetail db.WorkflowDetail) {
	projectIdStr, workflowId, err := GetProjectIdAndWorkflowIdByWorkflowKey(message.JobName)
	if err != nil {
		return
	}
	projectId, err := uuid.FromString(projectIdStr)
	if err != nil {
		log.Println("UUID from string failed: ", err.Error())
		return
	}
	jobDetail, err := w.engine.GetJobHistory(message.JobName, message.JobId)
	if err != nil {
		return
	}

	if len(jobDetail.Artifactorys) > 0 {

		for _, arti := range jobDetail.Artifactorys {

			// 判断是不是 starknet 合约
			isStarknetContract := false
			starknetContractClassHash := ""
			if strings.HasSuffix(arti.Url, "starknet.output.json") {
				isStarknetContract = true
				classHash, err := starkClassHash(arti.Url)
				if err != nil {
					logger.Errorf("starknet contract class hash failed: %s", err.Error())
					continue
				}
				starknetContractClassHash = classHash
				logger.Trace("starknet contract class hash: ", starknetContractClassHash)
			}

			data, _ := os.ReadFile(arti.Url)
			m := make(map[string]any)

			err := json.Unmarshal(data, &m)
			if err != nil {
				logger.Errorf("unmarshal contract abi failed: %s", err.Error())
				continue
			}

			var abi string
			if !isStarknetContract {
				abiByte, err := json.Marshal(m["abi"])
				if err != nil {
					logger.Errorf("marshal contract abi failed: %s", err.Error())
					continue
				}
				abi = string(abiByte)
			} else {
				abi = string(data)
			}

			var bytecodeData string
			if !isStarknetContract {
				var ok bool
				bytecodeData, ok = m["bytecode"].(string)
				if !ok {
					logger.Errorf("contract bytecode is not string")
					continue
				}
			} else {
				bytecodeData = starknetContractClassHash
			}

			var contractType uint
			if !isStarknetContract {
				contractType = consts.Evm
			} else {
				contractType = consts.StarkWare
			}

			contract := db.Contract{
				ProjectId:        projectId,
				WorkflowId:       workflowId,
				WorkflowDetailId: workflowDetail.Id,
				Name:             strings.TrimSuffix(arti.Name, path.Ext(arti.Name)),
				Version:          fmt.Sprintf("%d", workflowDetail.ExecNumber),
				BuildTime:        workflowDetail.CreateTime,
				AbiInfo:          abi,
				ByteCode:         bytecodeData,
				CreateTime:       time.Now(),
				Type:             uint(contractType),
				Status:           consts.STATUS_SUCCESS,
			}
			err = w.db.Save(&contract).Error

			if err != nil {
				logger.Errorf("save contract to database failed: %s", err.Error())
				continue
			}
			logger.Trace("save contract to database success: ", contract.Name)

			// declare classHash
			if isStarknetContract {
				contractService := application.GetBean[*ContractService]("contractService")
				go func() {
					_, err := contractService.DoStarknetDeclare([]byte(contract.AbiInfo))
					if err != nil {
						logger.Trace("declare starknet abi error:", err.Error())
					}
				}()
			}

		}

	}
}

func (w *WorkflowService) SyncReport(message model.StatusChangeMessage, workflowDetail db.WorkflowDetail) {
	if !strings.Contains(message.JobName, "_") {
		return
	}
	projectIdStr, workflowId, err := GetProjectIdAndWorkflowIdByWorkflowKey(message.JobName)
	if err != nil {
		return
	}
	projectId, err := uuid.FromString(projectIdStr)
	if err != nil {
		log.Println("UUID from string failed: ", err.Error())
		return
	}
	workflowExecNumber := message.JobId

	if message.Status == model.STATUS_SUCCESS {
		//TODO.... 实现同步报告
		fmt.Println(projectId, workflowId, workflowExecNumber)
		jobDetail, err := w.engine.GetJobHistory(message.JobName, message.JobId)
		if err != nil {
			logger.Errorf("Get job history fail, jobName: %s, jobId: %d", message.JobName, message.JobId)
			return
		}
		logger.Tracef("Get job history success, jobName: %s, jobId: %d", message.JobName, message.JobId)
		logger.Tracef("len jobDetail.Reports: %d", len(jobDetail.Reports))
		logger.Tracef("jobDetail file path: %s", jober.GetJobDetailFilePath(message.JobName, message.JobId))
		var reportList []db.Report
		begin := w.db.Begin()
		for _, report := range jobDetail.Reports {

			// contract check
			if report.Type == 2 {
				if report.Url == "" {
					continue
				}
				file, err := os.ReadFile(report.Url)
				if err != nil {
					logger.Errorf("Check result path is err")
					return
				}
				var contractCheckResultList []model.ContractCheckResult[json.RawMessage]
				err = json.Unmarshal(file, &contractCheckResultList)
				if err != nil {
					logger.Errorf("Check result get fail")
				}
				for _, contractCheckResult := range contractCheckResultList {
					marshal, err := json.Marshal(contractCheckResult.Context)
					if err != nil {
						logger.Errorf("Check context conversion failed")
					}
					report := db.Report{
						ProjectId:        projectId,
						WorkflowId:       workflowId,
						WorkflowDetailId: workflowDetail.Id,
						Name:             contractCheckResult.Name,
						Type:             uint(consts.Check),
						CheckTool:        contractCheckResult.Tool,
						// CheckVersion:     contractCheckResult.SolcVersion,
						Result:     contractCheckResult.Result,
						CheckTime:  time.Now(),
						ReportFile: string(marshal),
						CreateTime: time.Now(),
					}
					reportList = append(reportList, report)
				}
			}
			// openai report
			if report.Type == 3 {
				report := db.Report{
					ProjectId:        projectId,
					WorkflowId:       workflowId,
					WorkflowDetailId: workflowDetail.Id,
					Name:             "AI Analysis Report",
					Type:             uint(consts.Check),
					CheckTool:        "OpenAI",
					Result:           "success",
					CheckTime:        time.Now(),
					// ReportFile:       string(report.Content),
					CreateTime: time.Now(),
				}
				reportList = append(reportList, report)
			}
		}
		logger.Tracef("len(reportList): %d ", len(reportList))
		err = begin.Save(&reportList).Error
		if err != nil {
			logger.Errorf("Save report fail, err is %s", err.Error())
			return
		}
		begin.Commit()
	}

}

func (w *WorkflowService) ExecProjectCheckWorkflow(projectId uuid.UUID, user vo.UserAuth) error {
	_, err := w.ExecProjectWorkflow(projectId, user, 1, nil)
	return err
}

func (w *WorkflowService) ExecProjectBuildWorkflow(projectId uuid.UUID, user vo.UserAuth) error {
	_, err := w.ExecProjectWorkflow(projectId, user, 2, nil)
	return err
}

func (w *WorkflowService) ExecProjectDeployWorkflow(projectId uuid.UUID, buildWorkflowId, buildWorkflowDetailId int, user vo.UserAuth) (vo.DeployResultVo, error) {
	buildWorkflowKey := w.GetWorkflowKey(projectId.String(), uint(buildWorkflowId))

	workflowDetail, err := w.GetWorkflowDetail(buildWorkflowId, buildWorkflowDetailId)
	if err != nil {
		logger.Info("workflow ")
		return vo.DeployResultVo{}, err
	}
	buildJobDetail, err := w.engine.GetJobHistory(buildWorkflowKey, int(workflowDetail.ExecNumber))
	if err != nil {
		return vo.DeployResultVo{}, err
	}

	if len(buildJobDetail.ActionResult.Artifactorys) == 0 {
		return vo.DeployResultVo{}, errors.New("No Artifacts")
	}

	params := make(map[string]string)
	params["baseDir"] = "dist"
	params["ArtifactUrl"] = "file://" + buildJobDetail.Artifactorys[0].Url
	params["buildWorkflowDetailId"] = strconv.Itoa(buildWorkflowDetailId)
	return w.ExecProjectWorkflow(projectId, user, 3, params)
}

func (w *WorkflowService) ExecProjectWorkflow(projectId uuid.UUID, user vo.UserAuth, workflowType uint, params map[string]string) (vo.DeployResultVo, error) {

	// query project workflow

	var workflow db.Workflow
	var deployResult vo.DeployResultVo
	w.db.Where(&db.Workflow{
		ProjectId: projectId,
		Type:      workflowType,
	}).First(&workflow)

	if &workflow == nil {
		return deployResult, errors.New("no check workflow in the project ")
	}

	workflowKey := w.GetWorkflowKey(projectId.String(), workflow.Id)

	logger.Tracef("workflow key is %s", workflowKey)
	job, err := w.engine.GetJob(workflowKey)
	if err != nil {
		logger.Tracef("job is not exist, create job: %s", workflowKey)
		var jobModel model.Job
		err := yaml.Unmarshal([]byte((workflow.ExecFile)), &jobModel)
		if err != nil {
			logger.Errorf("Unmarshal job fail, err is %s", err.Error())
			return deployResult, err
		}
		if jobModel.Name != workflowKey {
			jobModel.Name = workflowKey
			execFile, _ := yaml.Marshal(jobModel)
			workflow.ExecFile = string(execFile)
		}

		err = w.engine.CreateJob(workflowKey, workflow.ExecFile)
		if err != nil {
			return deployResult, err
		}
		job, err = w.engine.GetJob(workflowKey)
		if err != nil {
			logger.Errorf("Get job fail, err is %s", err.Error())
			return deployResult, err
		}
		logger.Tracef("create job success, job name is %s", job.Name)
	}

	if params != nil {
		if job.Parameter == nil {
			job.Parameter = params
		} else {
			for k, v := range params {
				job.Parameter[k] = v
			}
		}
		err := w.engine.SaveJobParams(job.Name, params)
		if err != nil {
			return deployResult, err
		}
	}

	detail, err := w.engine.ExecuteJob(workflowKey)

	if err != nil {
		logger.Errorf("Create job detail fail, err is %s", err.Error())
		return deployResult, err
	}
	stageInfo, err := json.Marshal(detail.Stages)
	if err != nil {
		logger.Errorf("Marshal stage info fail, err is %s", err.Error())
		return deployResult, err
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
		//Status:      uint(detail.Status),
		Status:     1,
		StartTime:  detail.StartTime,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	err = w.db.Transaction(func(tx *gorm.DB) error {
		tx.Save(&dbDetail)
		return nil
	})

	if err != nil {
		logger.Errorf("Save workflow detail fail, err is %s", err.Error())
		return deployResult, err
	}
	deployResult.WorkflowId = workflow.Id
	deployResult.DetailId = dbDetail.Id
	logger.Tracef("create job detail success, job detail id is %d", detail.Id)
	return deployResult, nil
}

func (w *WorkflowService) GetWorkflowList(projectId string, workflowType, page, size int) (*vo.WorkflowPage, error) {
	var total int64
	var data vo.WorkflowPage
	var workflowData []vo.WorkflowVo
	var workflowList []db.WorkflowDetail
	tx := w.db.Model(db.WorkflowDetail{}).Where("project_id = ?", projectId)
	if workflowType != 0 {
		tx = tx.Where("type = ? ", workflowType)
	}
	result := tx.Offset((page - 1) * size).Limit(size).Find(&workflowList).Offset(-1).Limit(-1).Count(&total)
	if result.Error != nil {
		return &data, result.Error
	}
	if len(workflowList) > 0 {
		for _, datum := range workflowList {
			var resData vo.WorkflowVo
			_ = copier.Copy(&resData, &datum)
			resData.DetailId = datum.Id
			resData.Id = datum.WorkflowId
			workflowData = append(workflowData, resData)
		}
	}
	data.Data = workflowData
	data.Total = int(total)
	data.Page = page
	data.PageSize = size
	return &data, nil
}

func (w *WorkflowService) GetWorkflowDetail(workflowId, workflowDetailId int) (*vo.WorkflowDetailVo, error) {
	var workflowDetail db.WorkflowDetail
	var detail vo.WorkflowDetailVo
	res := w.db.Model(db.WorkflowDetail{}).Where("workflow_id = ? and id = ?", workflowId, workflowDetailId).First(&workflowDetail)
	if res.Error != nil {
		return &detail, res.Error
	}

	_ = copier.Copy(&detail, &workflowDetail)
	if workflowDetail.Status == vo.WORKFLOW_STATUS_RUNNING {
		workflowKey := w.GetWorkflowKey(workflowDetail.ProjectId.String(), workflowDetail.WorkflowId)
		jobDetail, err := w.engine.GetJobHistory(workflowKey, int(workflowDetail.ExecNumber))
		if err != nil {
			logger.Warnf("get job history fail, err is %s", err.Error())
		}
		data, err := json.Marshal(jobDetail.Stages)
		if err == nil {
			detail.StageInfo = string(data)
			detail.Duration = jobDetail.Duration
		}
	}
	return &detail, nil
}

func (w *WorkflowService) QueryWorkflowDetail(workflowId, workflowDetailId int) (*db.WorkflowDetail, error) {
	var workflowDetail db.WorkflowDetail
	res := w.db.Model(db.WorkflowDetail{}).Where("workflow_id = ? and id = ?", workflowId, workflowDetailId).First(&workflowDetail)
	if res.Error != nil {
		return &workflowDetail, res.Error
	}
	return &workflowDetail, nil
}

func (w *WorkflowService) QueryWorkflow(workflowId int) (*db.Workflow, error) {
	var workflow db.Workflow
	res := w.db.Model(db.Workflow{}).Where("id = ?", workflowId).First(&workflow)
	if res.Error != nil {
		return &workflow, res.Error
	}
	return &workflow, nil
}

func (w *WorkflowService) GetWorkflowKey(projectId string, workflowId uint) string {
	return fmt.Sprintf("%s_%d", projectId, workflowId)
}

func GetProjectIdAndWorkflowIdByWorkflowKey(projectKey string) (string, uint, error) {
	projectId := strings.Split(projectKey, "_")[0]
	workflowId, err := strconv.Atoi(strings.Split(projectKey, "_")[1])
	if err != nil {
		return projectId, 0, err
	}
	return projectId, uint(workflowId), err

}

func (w *WorkflowService) SaveWorkflow(saveData parameter.SaveWorkflowParam) (db.Workflow, error) {
	var workflow db.Workflow
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

func (w *WorkflowService) UpdateWorkflow(data db.Workflow) error {
	res := w.db.Save(&data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func getTemplate(projectType, projectFrameType uint, workflowType consts.WorkflowType) string {
	filePath := "templates/truffle-build.yml"
	if projectType == uint(consts.CONTRACT) {
		if workflowType == consts.Check {
			filePath = "templates/truffle_check.yml"
		} else if workflowType == consts.Build {
			if projectFrameType == uint(consts.StarkWare) {
				filePath = "templates/stark-ware-build.yml"
			} else {
				filePath = "templates/truffle-build.yml"
			}
		}
	} else if projectType == uint(consts.FRONTEND) {
		if workflowType == consts.Check {
			filePath = "templates/frontend-check.yml"
		} else if workflowType == consts.Build {
			filePath = "templates/frontend-build.yml"
		} else if workflowType == consts.Deploy {
			filePath = "templates/frontend-deploy.yml"
		}
	}
	return filePath
}

func (w *WorkflowService) TemplateParse(name string, project *vo.ProjectDetailVo, workflowType consts.WorkflowType) (string, error) {
	if project == nil {
		return "", errors.New("project is nil")
	}
	filePath := getTemplate(project.Type, uint(project.FrameType), workflowType)
	content, err := temp.ReadFile(filePath)
	if err != nil {
		log.Println("read template file failed ", err.Error())
		return "", err
	}
	fileContent := string(content)

	tmpl := template.New("test")
	if workflowType == consts.Deploy {
		tmpl = tmpl.Delims("[[", "]]")
	}

	tmpl, err = tmpl.Parse(fileContent)

	if err != nil {
		log.Println("template parse failed ", err.Error())
		return "", err
	}

	templateData := parameter.TemplateCheck{
		Name:          name,
		RepositoryUrl: project.RepositoryUrl,
	}
	var input bytes.Buffer
	err = tmpl.Execute(&input, templateData)
	if err != nil {
		log.Println("failed to write parameters to the template ", err)
		return "", err
	}
	return input.String(), nil
}

func (w *WorkflowService) DeleteWorkflow(workflowId, detailId int) error {
	err := w.db.Debug().Where("id = ? and workflow_id = ?", detailId, workflowId).Delete(&db.WorkflowDetail{}).Error
	if err != nil {
		return err
	}
	return nil
}

func starkClassHash(filename string) (string, error) {
	cmdStr := "starkli class-hash " + filename
	cmd := exec.Command("/bin/bash", "-c", cmdStr)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	classHash := strings.TrimSpace(out.String())
	return classHash, nil
}
