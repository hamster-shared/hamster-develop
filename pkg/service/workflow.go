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
	"path"
	"strconv"
	"strings"
	"text/template"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	engine "github.com/hamster-shared/aline-engine"
	jober "github.com/hamster-shared/aline-engine/job"
	"github.com/hamster-shared/aline-engine/logger"
	"github.com/hamster-shared/aline-engine/model"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	"github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/utils"
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
		} else if project.FrameType == 2 {
			image = "https://develop-images.api.hamsternet.io/react.png"
		} else if project.FrameType == 3 {
			image = "https://static.devops.hamsternet.io/ipfs/QmW8DNyCUrvDHaG4a4aKjkDNTbYDy9kwFxhFno2nKmgTKt"
		} else {
			image = "https://static.devops.hamsternet.io/ipfs/QmPsa61VtwQH3ixzZys7EF9VG1zV7LQHDYjEYBfZpnmPDy"
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
				if project.DeployType == int(consts.IPFS) {
					packageDeploy.DeployInfo = deploy.Cid
				}
				packageDeploy.ProjectId = project.Id
				packageDeploy.WorkflowId = workflowDetail.WorkflowId
				packageDeploy.WorkflowDetailId = workflowDetail.Id
				packageDeploy.PackageId = data.Id
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

	if len(jobDetail.Artifactorys) == 0 {
		return
	}

	for _, arti := range jobDetail.Artifactorys {
		// 如果是 starknet 合约
		if strings.HasSuffix(arti.Url, "starknet.output.json") {
			err := w.syncContractStarknet(projectId, workflowId, workflowDetail, arti)
			if err != nil {
				logger.Errorf("sync contract starknet failed: %s", err.Error())
			}
			continue
		}

		// 如果以 .mv 或者 .bcs 结尾，认为是 aptos 合约，退出循环，在另一个函数中处理
		if strings.HasSuffix(arti.Url, ".mv") || strings.HasSuffix(arti.Url, ".bcs") {
			err := w.syncContractAptos(projectId, workflowId, workflowDetail, jobDetail.Artifactorys)
			if err != nil {
				logger.Errorf("sync contract aptos failed: %s", err.Error())
			}
			return
		}

		// zip 文件不处理
		if strings.HasSuffix(arti.Url, ".zip") {
			continue
		}

		// 其他作为 evm 合约处理
		w.syncContractEvm(projectId, workflowId, workflowDetail, arti)
	}
}

func (w *WorkflowService) saveContractToDatabase(contract *db.Contract) error {
	err := w.db.Save(contract).Error
	if err != nil {
		logger.Errorf("save contract to database failed: %s", err.Error())
		return err
	}
	logger.Trace("save contract to database success: ", contract.Name)
	return nil
}

func (w *WorkflowService) getEvmAbiInfoAndByteCode(arti model.Artifactory) (abiInfo string, byteCode string, err error) {
	data, _ := os.ReadFile(arti.Url)
	m := make(map[string]any)

	err = json.Unmarshal(data, &m)
	if err != nil {
		logger.Errorf("unmarshal contract abi failed: %s", err.Error())
		return "", "", err
	}

	abiByte, err := json.Marshal(m["abi"])
	if err != nil {
		logger.Errorf("marshal contract abi failed: %s", err.Error())
		return "", "", err
	}
	abiInfo = string(abiByte)

	byteCode, ok := m["bytecode"].(string)
	if !ok {
		logger.Errorf("contract bytecode is not string")
		return "", "", err
	}
	return abiInfo, byteCode, nil
}

func (w *WorkflowService) syncContractEvm(projectId uuid.UUID, workflowId uint, workflowDetail db.WorkflowDetail, arti model.Artifactory) error {
	abiInfo, byteCode, err := w.getEvmAbiInfoAndByteCode(arti)
	if err != nil {
		return err
	}

	contract := db.Contract{
		ProjectId:        projectId,
		WorkflowId:       workflowId,
		WorkflowDetailId: workflowDetail.Id,
		Name:             strings.TrimSuffix(arti.Name, path.Ext(arti.Name)),
		Version:          fmt.Sprintf("%d", workflowDetail.ExecNumber),
		BuildTime:        workflowDetail.CreateTime,
		AbiInfo:          abiInfo,
		ByteCode:         byteCode,
		CreateTime:       time.Now(),
		Type:             uint(consts.Evm),
		Status:           consts.STATUS_SUCCESS,
	}
	return w.saveContractToDatabase(&contract)
}

func (w *WorkflowService) getAptosMvAndByteCode(artis []model.Artifactory) (mv string, byteCode string, err error) {
	for _, arti := range artis {
		// 此处逻辑存疑
		// 以 .bcs 结尾，认为是 byteCode
		if strings.HasSuffix(arti.Url, ".bcs") {
			byteCode, err = utils.FileToHexString(arti.Url)
			if err != nil {
				logger.Errorf("hex string failed: %s", err.Error())
				return "", "", err
			}
			continue
		}
		// 以 .mv 结尾，认为是 abi
		if strings.HasSuffix(arti.Url, ".bcs") {
			mv, err = utils.FileToHexString(arti.Url)
			if err != nil {
				logger.Errorf("hex string failed: %s", err.Error())
				return "", "", err
			}
			continue
		}
		logger.Warnf("aptos contract file name is not end with .bcs or .mv: %s", arti.Url)
	}
	return mv, byteCode, nil
}

func (w *WorkflowService) syncContractAptos(projectId uuid.UUID, workflowId uint, workflowDetail db.WorkflowDetail, artis []model.Artifactory) error {
	mv, byteCode, err := w.getAptosMvAndByteCode(artis)
	if err != nil {
		return err
	}

	contract := db.Contract{
		ProjectId:        projectId,
		WorkflowId:       workflowId,
		WorkflowDetailId: workflowDetail.Id,
		Name:             strings.TrimSuffix(artis[0].Name, path.Ext(artis[0].Name)),
		Version:          fmt.Sprintf("%d", workflowDetail.ExecNumber),
		BuildTime:        workflowDetail.CreateTime,
		AbiInfo:          "",
		ByteCode:         byteCode,
		AptosMv:          mv,
		CreateTime:       time.Now(),
		Type:             uint(consts.Aptos),
		Status:           consts.STATUS_SUCCESS,
	}

	return w.saveContractToDatabase(&contract)
}

func (w *WorkflowService) getStarknetAbiInfoAndByteCode(artiUrl string) (abiInfo string, byteCode string, err error) {
	data, err := os.ReadFile(artiUrl)
	if err != nil {
		return "", "", err
	}
	m := make(map[string]any)
	err = json.Unmarshal(data, &m)
	if err != nil {
		return "", "", err
	}
	abiBytes, err := json.Marshal(m["abi"])
	if err != nil {
		return "", "", err
	}
	abiInfo = string(abiBytes)
	contractService := application.GetBean[*ContractService]("contractService")
	_, classHash, err := contractService.DoStarknetDeclare(data)
	if err != nil {
		logger.Errorf("starknet contract class hash failed: %s", err.Error())
		return "", "", err
	}
	logger.Trace("starknet contract class hash: ", classHash)
	byteCode = classHash
	return
}

func (w *WorkflowService) syncContractStarknet(projectId uuid.UUID, workflowId uint, workflowDetail db.WorkflowDetail, arti model.Artifactory) error {
	abiInfo, byteCode, err := w.getStarknetAbiInfoAndByteCode(arti.Url)
	if err != nil {
		return err
	}

	contract := db.Contract{
		ProjectId:        projectId,
		WorkflowId:       workflowId,
		WorkflowDetailId: workflowDetail.Id,
		Name:             strings.TrimSuffix(arti.Name, path.Ext(arti.Name)),
		Version:          fmt.Sprintf("%d", workflowDetail.ExecNumber),
		BuildTime:        workflowDetail.CreateTime,
		AbiInfo:          abiInfo,
		ByteCode:         byteCode,
		CreateTime:       time.Now(),
		Type:             uint(consts.StarkWare),
		Status:           consts.STATUS_SUCCESS,
	}

	return w.saveContractToDatabase(&contract)
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
			// return
		}
		begin.Commit()
	}

}

func (w *WorkflowService) ExecProjectCheckWorkflow(projectId uuid.UUID, user vo.UserAuth) error {
	_, err := w.ExecProjectWorkflow(projectId, user, 1, nil)
	return err
}

func (w *WorkflowService) ExecProjectBuildWorkflow(projectId uuid.UUID, user vo.UserAuth) (vo.DeployResultVo, error) {
	var project db.Project
	err := w.db.Model(db.Project{}).Where("id = ?", projectId.String()).First(&project).Error
	if err != nil {
		logger.Info("project is not exit ")
		return vo.DeployResultVo{}, err
	}
	params := make(map[string]string)
	if project.Type == uint(consts.FRONTEND) && project.DeployType == int(consts.CONTAINER) {
		image := fmt.Sprintf("%s/%s-%d:%d", consts.DockerHubName, strings.ToLower(user.Username), user.Id, time.Now().Unix())
		params["imageName"] = image
		if project.FrameType == 1 || project.FrameType == 2 {
			params["runBuild"] = "true"
		}
	} else {
		params = nil
	}
	data, err := w.ExecProjectWorkflow(projectId, user, 2, params)
	return data, err
}

func (w *WorkflowService) ExecProjectDeployWorkflow(projectId uuid.UUID, buildWorkflowId, buildWorkflowDetailId int, user vo.UserAuth) (vo.DeployResultVo, error) {
	buildWorkflowKey := w.GetWorkflowKey(projectId.String(), uint(buildWorkflowId))

	workflowDetail, err := w.GetWorkflowDetail(buildWorkflowId, buildWorkflowDetailId)
	if err != nil {
		logger.Errorf("workflow : %s", err)
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

func (w *WorkflowService) ExecContainerDeploy(projectId uuid.UUID, buildWorkflowId, buildWorkflowDetailId int, user vo.UserAuth, deployParam parameter.K8sDeployParam) (vo.DeployResultVo, error) {
	var project db.Project
	err := w.db.Model(db.Project{}).Where("id = ?", projectId.String()).First(&project).Error
	if err != nil {
		logger.Info("project is not exit ")
		return vo.DeployResultVo{}, err
	}
	buildWorkflowKey := w.GetWorkflowKey(projectId.String(), uint(buildWorkflowId))

	workflowDetail, err := w.GetWorkflowDetail(buildWorkflowId, buildWorkflowDetailId)
	if err != nil {
		logger.Errorf("GetWorkflowDetail err: %s", err.Error())
		return vo.DeployResultVo{}, err
	}
	buildJobDetail, err := w.engine.GetJobHistory(buildWorkflowKey, int(workflowDetail.ExecNumber))
	if err != nil {
		logger.Errorf("GetJobHistory err: %s", err.Error())
		return vo.DeployResultVo{}, err
	}
	if len(buildJobDetail.ActionResult.BuildData) == 0 {
		return vo.DeployResultVo{}, errors.New("No Image")
	}
	var containers []corev1.Container
	var ports []corev1.ContainerPort
	port := corev1.ContainerPort{
		ContainerPort: deployParam.ContainerPort,
	}
	resources := corev1.ResourceRequirements{
		Limits: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("500m"),
			corev1.ResourceMemory: resource.MustParse("500Mi"),
		},
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("50m"),
			corev1.ResourceMemory: resource.MustParse("50Mi"),
		},
	}
	projectName := strings.Replace(project.Name, "_", "-", -1)
	ports = append(ports, port)
	container1 := corev1.Container{
		Name:      fmt.Sprintf("%s-%s", strings.ToLower(user.Username), strings.ToLower(projectName)),
		Image:     buildJobDetail.ActionResult.BuildData[0].ImageName,
		Ports:     ports,
		Resources: resources,
	}
	containers = append(containers, container1)
	containerStr, err := json.Marshal(containers)
	if err != nil {
		logger.Info("containers json marshal failed ")
		return vo.DeployResultVo{}, err
	}
	var servicePorts []parameter.ServicePort
	servicePort := parameter.ServicePort{
		Protocol:   deployParam.ServiceProtocol,
		Port:       deployParam.ServicePort,
		TargetPort: deployParam.ServiceTargetPort,
	}
	servicePorts = append(servicePorts, servicePort)
	serviceStr, err := json.Marshal(servicePorts)
	if err != nil {
		logger.Info("services json marshal failed ")
		return vo.DeployResultVo{}, err
	}
	params := make(map[string]string)
	params["namespace"] = strings.ToLower(user.Username)
	params["projectName"] = strings.ToLower(projectName)
	params["servicePorts"] = string(serviceStr)
	params["containers"] = string(containerStr)
	params["gateway"] = consts.Gateway
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
			return &detail, err
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

func getTemplate(project *vo.ProjectDetailVo, workflowType consts.WorkflowType) string {
	filePath := "templates/truffle-build.yml"
	if project.Type == uint(consts.CONTRACT) {
		if workflowType == consts.Check {
			filePath = "templates/truffle_check.yml"
		} else if workflowType == consts.Build {
			if project.FrameType == uint(consts.StarkWare) {
				filePath = "templates/stark-ware-build.yml"
			} else {
				filePath = "templates/truffle-build.yml"
			}
		}
	} else if project.Type == uint(consts.FRONTEND) {
		if workflowType == consts.Check {
			filePath = "templates/frontend-check.yml"
		} else if workflowType == consts.Build {
			if project.DeployType == int(consts.IPFS) {
				filePath = "templates/frontend-build.yml"
			} else {
				filePath = "templates/frontend-image-build.yml"
			}
		} else if workflowType == consts.Deploy {
			if project.DeployType == int(consts.IPFS) {
				filePath = "templates/frontend-deploy.yml"
			} else {
				filePath = "templates/frontend-k8s-deploy.yml"
			}
		}
	}
	return filePath
}

func (w *WorkflowService) TemplateParse(name string, project *vo.ProjectDetailVo, workflowType consts.WorkflowType) (string, error) {
	if project == nil {
		return "", errors.New("project is nil")
	}
	filePath := getTemplate(project, workflowType)
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
	if project.Type == uint(consts.FRONTEND) && project.DeployType == int(consts.CONTAINER) && workflowType == consts.Build {
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

func (w *WorkflowService) CheckRunningJob() {

	var workflowList []db.WorkflowDetail
	err := w.db.Model(db.WorkflowDetail{}).Where("status = ?", vo.WORKFLOW_STATUS_RUNNING).Find(&workflowList).Error
	if err != nil {
		return
	}

	stopList := make([]db.WorkflowDetail, 0)
	for _, flow := range workflowList {
		workflowKey := w.GetWorkflowKey(flow.ProjectId.String(), uint(flow.WorkflowId))
		jobDetail, err := w.engine.GetJobHistory(workflowKey, int(flow.ExecNumber))

		if err != nil {
			stopList = append(stopList, flow)
			continue
		}

		if jobDetail.Status == model.STATUS_RUNNING {
			// check it is really running
			status, err := w.engine.GetCurrentJobStatus(workflowKey, int(flow.ExecNumber))
			if err != nil || status != model.STATUS_RUNNING {
				_ = w.engine.TerminalJob(workflowKey, int(flow.ExecNumber))
				stopList = append(stopList, flow)
			}
		}
	}

	for _, flow := range stopList {
		workflowKey := w.GetWorkflowKey(flow.ProjectId.String(), flow.WorkflowId)
		jobDetail, _ := w.engine.GetJobHistory(workflowKey, int(flow.ExecNumber))
		flow.Status = vo.WORKFLOW_STATUS_CANCEL
		if jobDetail != nil {
			stageInfo, err := json.Marshal(jobDetail.Stages)
			if err == nil {
				flow.StageInfo = string(stageInfo)
			}
			flow.Duration = jobDetail.Duration
		}
		flow.UpdateTime = time.Now()
		_ = w.db.Save(flow).Error
	}
}
