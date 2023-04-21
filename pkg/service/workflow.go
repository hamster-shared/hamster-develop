package service

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	engine "github.com/hamster-shared/aline-engine"
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

func (w *WorkflowService) ExecProjectCheckWorkflow(projectId uuid.UUID, user vo.UserAuth) error {
	var project db.Project
	err := w.db.Model(db.Project{}).Where("id = ?", projectId.String()).First(&project).Error
	if err != nil {
		logger.Info("project is not exit ")
		return err
	}
	params := make(map[string]string)
	if project.Type == uint(consts.CONTRACT) && project.FrameType == consts.Evm {
		params["projectName"] = fmt.Sprintf("%s/%s", user.Username, project.Name)
		params["projectUrl"] = project.RepositoryUrl
	}
	_, err = w.ExecProjectWorkflow(projectId, user, 1, params)
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
	//params["gateway"] = consts.Gateway
	params["gateway"] = os.Getenv("GATEWAY")
	params["buildWorkflowDetailId"] = strconv.Itoa(buildWorkflowDetailId)
	return w.ExecProjectWorkflow(projectId, user, 3, params)
}

func (w *WorkflowService) ExecProjectBuildWorkflowAptos(projectID uuid.UUID, user vo.UserAuth) (vo.DeployResultVo, error) {
	var project db.Project
	err := w.db.Model(db.Project{}).Where("id = ?", projectID.String()).First(&project).Error
	if err != nil {
		logger.Info("project is not exit ")
		return vo.DeployResultVo{}, err
	}
	params, err := utils.GetKeyValuesFromString(project.Params)
	if err != nil {
		logger.Errorf("project params is not valid %s", err)
		return vo.DeployResultVo{}, err
	}

	aptos_param := ""
	for k, v := range params {
		aptos_param += fmt.Sprintf("%s=%s,", k, v)
	}
	params["aptos_param"] = aptos_param
	return w.ExecProjectWorkflow(projectID, user, 2, params)
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
			logger.Errorf("job file is %s", workflow.ExecFile)
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
	met, token := setMetaScanToken(workflow)
	if met {
		params["metaScanToken"] = token
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

	// 从数据库获取最新的执行次数
	var workflowDetail db.WorkflowDetail
	var execNumber uint
	if w.db.Where(&db.WorkflowDetail{WorkflowId: workflow.Id}).Order("exec_number desc").First(&workflowDetail).Error == nil {
		execNumber = workflowDetail.ExecNumber
	} else {
		execNumber = 0
	}

	var detail *model.JobDetail
	var dbDetail db.WorkflowDetail
	// 重试 10 次
	for i := 0; i < 10; i++ {
		detail, err = w.engine.CreateJobDetail(workflowKey, int(execNumber)+1+i)
		if err != nil {
			logger.Errorf("Create job detail fail, err is %s", err.Error())
			return deployResult, err
		}
		var stageInfo []byte
		stageInfo, err = json.Marshal(detail.Stages)
		if err != nil {
			logger.Errorf("Marshal stage info fail, err is %s", err.Error())
			return deployResult, err
		}

		dbDetail = db.WorkflowDetail{
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
			return tx.Save(&dbDetail).Error
		})

		if err != nil {
			logger.Warnf("Save workflow detail fail, err is %s, retry counter: %d", err.Error(), i)
		} else {
			logger.Infof("create job detail success, job detail id is %d", detail.Id)
			break
		}
	}
	// 重试 10 次后仍然失败，返回错误
	if err != nil {
		return deployResult, err
	}
	deployResult.WorkflowId = workflow.Id
	deployResult.DetailId = dbDetail.Id
	err = w.engine.ExecuteJobDetail(workflowKey, detail.Id)
	if err != nil {
		logger.Errorf("execute job detail fail, err is %s", err.Error())
		return deployResult, err
	}
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
	workflow.ToolType = saveData.ToolType
	workflow.Tool = saveData.Tool
	res := w.db.Save(&workflow)
	if res.Error != nil {
		return workflow, res.Error
	}
	return workflow, nil
}

func (w *WorkflowService) SettingWorkflow(settingData parameter.SaveWorkflowParam, projectData *vo.ProjectDetailVo) error {
	var workflow db.Workflow
	err := w.db.Model(db.Workflow{}).Where("project_id=? and type=?", settingData.ProjectId, consts.Check).First(&workflow).Error
	if err == gorm.ErrRecordNotFound {
		workflowCheckRes, err := w.SaveWorkflow(settingData)
		if err != nil {
			return err
		}
		checkKey := w.GetWorkflowKey(settingData.ProjectId.String(), workflowCheckRes.Id)
		file, err := w.TemplateParseV2(checkKey, settingData.Tool, projectData)
		if err != nil {
			return err
		}
		workflowCheckRes.ExecFile = file
		w.UpdateWorkflow(workflowCheckRes)
		return nil
	}
	return errors.New("workflow already set")
}

func (w *WorkflowService) WorkflowSettingCheck(projectId string, workflowType consts.WorkflowType) bool {
	var workflow db.Workflow
	err := w.db.Model(db.Workflow{}).Where("project_id=? and type=?", projectId, workflowType).First(&workflow).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	if workflow.ExecFile != "" {
		return true
	}
	return true
}

func (w *WorkflowService) UpdateWorkflow(data db.Workflow) error {
	res := w.db.Save(&data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func getCheckTemplate(tool string) string {
	var filePath string
	switch tool {
	case "MetaTrust (SA)", "MetaTrust (SP)", "MetaTrust (OSA)", "MetaTrust (CQ)":
		filePath = "templates/metascan-check.yml"
	case "Mythril", "Solhint", "eth-gas-reporter", "AI":
		filePath = "templates/truffle_check.yml"
	default:
		filePath = ""
	}
	return filePath
}

func getTemplate(project *vo.ProjectDetailVo, workflowType consts.WorkflowType) string {
	filePath := "templates/truffle-build.yml"
	if project.Type == uint(consts.CONTRACT) {
		if workflowType == consts.Check {
			if project.FrameType == consts.Sui {
				filePath = "templates/sui-check.yml"
			} else {
				filePath = "templates/truffle_check.yml"
			}
		} else if workflowType == consts.Build {
			if project.FrameType == uint(consts.StarkWare) {
				filePath = "templates/stark-ware-build.yml"
			} else if project.FrameType == consts.Aptos {
				filePath = "templates/aptos-build.yml"
			} else if project.FrameType == consts.Sui {
				filePath = "templates/sui-build.yml"
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
				if project.FrameType == 1 || project.FrameType == 2 {
					filePath = "templates/frontend-image-build.yml"
				} else {
					filePath = "templates/frontend-node-image-build.yml"
				}
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

func (w *WorkflowService) TemplateParseV2(name, tool string, project *vo.ProjectDetailVo) (string, error) {
	if project == nil {
		return "", errors.New("project is nil")
	}
	filePath := getCheckTemplate(tool)
	content, err := temp.ReadFile(filePath)
	if err != nil {
		log.Println("read template file failed ", err.Error())
		return "", err
	}
	fileContent := string(content)
	tmpl := template.New("test")
	var templateData interface{}
	switch tool {
	case "MetaTrust (SA)", "MetaTrust (SP)", "MetaTrust (OSA)", "MetaTrust (CQ)":
		tmpl = tmpl.Delims("[[", "]]")
		templateData = parameter.MetaScanCheck{
			Name: name,
			Tool: tool,
		}
	default:
		tmpl = tmpl.Delims("{{", "}}")
		templateData = parameter.TemplateCheck{
			Name:          name,
			RepositoryUrl: project.RepositoryUrl,
		}
	}
	tmpl, err = tmpl.Parse(fileContent)
	if err != nil {
		log.Println("template parse failed ", err.Error())
		return "", err
	}
	var input bytes.Buffer
	err = tmpl.Execute(&input, templateData)
	if err != nil {
		log.Println("failed to write parameters to the template ", err)
		return "", err
	}
	return input.String(), nil
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
	if project.Type == uint(consts.CONTRACT) {
		if workflowType == consts.Build && (project.FrameType == consts.Aptos || project.FrameType == consts.Sui) {
			tmpl = tmpl.Delims("[[", "]]")
		}
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

func setMetaScanToken(workflow db.Workflow) (bool, string) {
	token := ""
	metaScanFlag := false
	switch workflow.Tool {
	case "MetaTrust (SA)", "MetaTrust (SP)", "MetaTrust (OSA)", "MetaTrust (CQ)":
		metaScanFlag = true
		token = metaScanHttpRequestToken()
	case "Mythril", "Solhint", "eth-gas-reporter", "AI":
		token = ""
		metaScanFlag = false
	default:
		metaScanFlag = false
		token = ""
	}
	return metaScanFlag, token
}

func metaScanHttpRequestToken() string {
	url := "https://account.metatrust.io/realms/mt/protocol/openid-connect/token"
	token := struct {
		AccessToken      string `json:"access_token"`
		ExpiresIn        int64  `json:"expires_in"`
		RefreshExpiresIn int64  `json:"refresh_expires_in"`
		RefreshToken     string `json:"refresh_token"`
		TokenType        string `json:"token_type"`
		NotBeforePolicy  int    `json:"not-before-policy"`
		SessionState     string `json:"session_state"`
		Scope            string `json:"scope"`
	}{}
	res, err := utils.NewHttp().NewRequest().SetFormData(map[string]string{
		"grant_type": "password",
		"username":   "tom@hamsternet.io",
		"password":   "pysded-hismoh-3Dagcy",
		"client_id":  "webapp",
	}).SetResult(&token).SetHeader("Content-Type", "application/x-www-form-urlencoded").Post(url)
	if res.StatusCode() != 200 {
		return ""
	}
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
}

func GetProjectList() {
	url := "https://app.metatrust.io/api/project"
	result := struct {
		Data MetaProjectsData `json:"data"`
	}{}
	res, err := utils.NewHttp().NewRequest().SetQueryParams(map[string]string{
		"title": "ddsss",
	}).SetResult(&result).
		SetHeaders(map[string]string{
			"Authorization":  "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJiYXdqN2JZQjdHX2MtVDJXNmFiQkIwMHZld2xoaHZLVVNfSXJUTDFBdUs4In0.eyJleHAiOjE2ODE4OTc3NjAsImlhdCI6MTY4MTg5NTk2MCwianRpIjoiZmFmNDAwMzQtNWI5Yi00ZThmLTk0MDgtOWI3YzNiOGM5OWU5IiwiaXNzIjoiaHR0cHM6Ly9hY2NvdW50Lm1ldGF0cnVzdC5pby9yZWFsbXMvbXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiMjEzNjdlMGQtYWQ0NC00YTMwLWI4OWUtMDRmNDM2NWE4ZmM3IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2ViYXBwIiwic2Vzc2lvbl9zdGF0ZSI6IjQ5ZWYyYmRmLWYzMWItNDE3YS1hMTIzLTJlNTcxM2ZmYWQzNiIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOlsiaHR0cHM6Ly9hcHAubWV0YXRydXN0LmlvIiwiaHR0cHM6Ly9tZXRhdHJ1c3QuaW8iLCJodHRwczovL3d3dy5tZXRhdHJ1c3QuaW8iXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtbXQiLCJvZmZsaW5lX2FjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwiZGVsZXRlLWFjY291bnQiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJzaWQiOiI0OWVmMmJkZi1mMzFiLTQxN2EtYTEyMy0yZTU3MTNmZmFkMzYiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicHJlZmVycmVkX3VzZXJuYW1lIjoidG9tQGhhbXN0ZXJuZXQuaW8iLCJlbWFpbCI6InRvbUBoYW1zdGVybmV0LmlvIiwidXNlcm5hbWUiOiJ0b21AaGFtc3Rlcm5ldC5pbyJ9.wkgmwRviR32HnW0LAIMU2NfyoOZ3xHY2SGTFxE6uYYnsKsqwQBx3TUnZRX3g55yB8s296ydEVUFc6cFdnYtCtqihMRZEdTGLSfR3Nz39VGkMmuiA5rGfJDLmUZ2pXePIWfGFAYwjMsm2ArzUXnjphcu25d3eTMCN2iE3t8hOqIBZxmgF88uJpAPQ_tgdhh7PfGBtjFhdNapp94DnwurCwZKVTgr5K8s2Q68hUK5P2onacGXfsE0FwfTR0ePNBFWyfi72qzyVieWmu6bCHL57c4LiG8Aj6UmIJ3rGut-7DN5wt_INht6Np_MoaMMSGYIzxjiCvHbS6jYNudBU565ktA",
			"X-MetaScan-Org": "1078238259684835328",
		}).Get(url)
	if err != nil {
		log.Println("---------------")
		log.Println(err)
		log.Println("---------------")
		log.Println("创建失败")
		return
	}
	if res.StatusCode() == 401 {
		log.Println("权限认证失败")
		return
	}
	if res.StatusCode() != 200 {
		log.Println(res)
		log.Println("创建失败")
		return
	}
	log.Println(result)
}

type MetaScanProject struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type MetaProjectsData struct {
	Total      int               `json:"total"`
	TotalPages int               `json:"totalPages"`
	Items      []MetaScanProject `json:"items"`
}

func CreateMetaScanProject() {
	url := "https://app.metatrust.io/api/project"
	createData := struct {
		Title           string `json:"title"`
		RepoUrl         string `json:"repoUrl"`
		IntegrationType string `json:"integrationType"`
	}{
		Title:           "jiangzhihui/0420",
		RepoUrl:         "hamster-template/Token",
		IntegrationType: "github",
	}
	result := struct {
		Message string          `json:"message"`
		Data    MetaScanProject `json:"data"`
		Code    int64           `json:"code"`
	}{}
	res, err := utils.NewHttp().NewRequest().SetBody(&createData).
		SetHeaders(map[string]string{
			"Authorization":  "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJiYXdqN2JZQjdHX2MtVDJXNmFiQkIwMHZld2xoaHZLVVNfSXJUTDFBdUs4In0.eyJleHAiOjE2ODE5NzM1MzYsImlhdCI6MTY4MTk3MTczNiwianRpIjoiZDQwYTFjYjQtNGVjNC00NjllLWJhYTYtZjU4YmU0ZTY2NTMyIiwiaXNzIjoiaHR0cHM6Ly9hY2NvdW50Lm1ldGF0cnVzdC5pby9yZWFsbXMvbXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiMjEzNjdlMGQtYWQ0NC00YTMwLWI4OWUtMDRmNDM2NWE4ZmM3IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2ViYXBwIiwic2Vzc2lvbl9zdGF0ZSI6ImY0ZjY2MzYzLWI1MTMtNDNmMC05OGVhLTMxMjEwNmY5NDIxNSIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOlsiaHR0cHM6Ly9hcHAubWV0YXRydXN0LmlvIiwiaHR0cHM6Ly9tZXRhdHJ1c3QuaW8iLCJodHRwczovL3d3dy5tZXRhdHJ1c3QuaW8iXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtbXQiLCJvZmZsaW5lX2FjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwiZGVsZXRlLWFjY291bnQiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJzaWQiOiJmNGY2NjM2My1iNTEzLTQzZjAtOThlYS0zMTIxMDZmOTQyMTUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicHJlZmVycmVkX3VzZXJuYW1lIjoidG9tQGhhbXN0ZXJuZXQuaW8iLCJlbWFpbCI6InRvbUBoYW1zdGVybmV0LmlvIiwidXNlcm5hbWUiOiJ0b21AaGFtc3Rlcm5ldC5pbyJ9.jDrHK3uFqdrfIvx5gjhf4_Gg4S_bhnXV2QdeWR0AMbtevBRL3U2V4AxvooL4IBWBYeac3Q4-Hk4mfRH6keG26nU3B1PItnkJteoJ_VJRuc2Qg96MnfVr6S_wLnTD5OrzkZhm2rhvxXsofuU5aWxi8PlVvXKnUGF9upICCJaSjo6vBbBwMmE55s71JRI-WxwD8c2taRS-LvGiceofE2e6_jzMzAXxOrbqvQ44jZeypbsjknN8M6D__82N_9sa8hdi6DRbXpGot79wZI9i1zhXG_dI3h2VM-vcGL0O_I7yJqJY_aVmtDhb-gTZBlJTmw0TK_Fh6m2fnSbE-FWERHzwyQ",
			"X-MetaScan-Org": "1078238259684835328",
			"Content-Type":   "application/json",
		}).SetResult(&result).Post(url)
	log.Println("++++++++++++++++")
	log.Println(res.StatusCode())
	log.Println("++++++++++++++++")
	if err != nil {
		log.Println("---------------")
		log.Println(err)
		log.Println("---------------")
		log.Println("创建失败")
		return
	}
	if res.StatusCode() == 401 {
		log.Println("权限认证失败")
		return
	}
	if res.StatusCode() != 200 {
		log.Println(res)
		log.Println("创建失败")
		return
	}
	log.Println(res.Error())
	log.Println(result)
	log.Println("创建成功")
}

func StartScanTask() {
	url := "https://app.metatrust.io/api/scan/task"
	var types []string
	types = append(types, "STATIC")
	repoData := Repo{
		Branch:     "",
		CommitHash: "",
	}
	scanData := TaskScan{
		SubPath:      "",
		Mode:         "",
		IgnoredPaths: "node_modules,test,tests,mock",
	}
	envData := TaskEnv{
		Node:           "",
		Solc:           "",
		PackageManage:  "",
		CompileCommand: "",
		Variables:      "",
	}
	bodyData := struct {
		EngineTypes []string `json:"engine_types"`
		Repo        Repo     `json:"repo"`
		Scan        TaskScan `json:"scan"`
		Env         TaskEnv  `json:"env"`
	}{
		EngineTypes: types,
		Repo:        repoData,
		Scan:        scanData,
		Env:         envData,
	}
	var result StartTaskRes
	log.Println(bodyData.EngineTypes)
	res, err := utils.NewHttp().NewRequest().SetQueryParams(map[string]string{
		"action":    "start-scan",
		"projectId": "1098614733587611648",
	}).SetBody(&bodyData).SetHeaders(map[string]string{
		"Authorization":  "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJiYXdqN2JZQjdHX2MtVDJXNmFiQkIwMHZld2xoaHZLVVNfSXJUTDFBdUs4In0.eyJleHAiOjE2ODE5NzU2MDYsImlhdCI6MTY4MTk3MzgwNiwianRpIjoiYzQ4NDEyY2UtYjk1NS00MjU4LThkOWItNjZhNmRkOTc3MTllIiwiaXNzIjoiaHR0cHM6Ly9hY2NvdW50Lm1ldGF0cnVzdC5pby9yZWFsbXMvbXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiMjEzNjdlMGQtYWQ0NC00YTMwLWI4OWUtMDRmNDM2NWE4ZmM3IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2ViYXBwIiwic2Vzc2lvbl9zdGF0ZSI6IjQ0YmExOTM5LTljYWYtNDEwMC1hMjlmLTM1YmY3YWM0NWJmYyIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOlsiaHR0cHM6Ly9hcHAubWV0YXRydXN0LmlvIiwiaHR0cHM6Ly9tZXRhdHJ1c3QuaW8iLCJodHRwczovL3d3dy5tZXRhdHJ1c3QuaW8iXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtbXQiLCJvZmZsaW5lX2FjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwiZGVsZXRlLWFjY291bnQiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJzaWQiOiI0NGJhMTkzOS05Y2FmLTQxMDAtYTI5Zi0zNWJmN2FjNDViZmMiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicHJlZmVycmVkX3VzZXJuYW1lIjoidG9tQGhhbXN0ZXJuZXQuaW8iLCJlbWFpbCI6InRvbUBoYW1zdGVybmV0LmlvIiwidXNlcm5hbWUiOiJ0b21AaGFtc3Rlcm5ldC5pbyJ9.BPunpK_dUysbfrs3ISwyiP3soVz9tNU5G9GInCuLzlko2g3DGkNOsgm0lW69H_ykN-gqAmTvZEdB-4EE4hYM6KcJvaQ90ZmwpAMoI77NYI2YsWeJeMWBGw41No9jHKeTrUnWhMhuNflSdplqpM5ybCV0wWu2gk-OFdpdlbhd-utJrl7lawAs-CPMTO2Pv-JyvgrAyBFdD0B23g3O7wwaT5IPxdHomt4SVFYfmH0h1yeRaKyt-7szBfccFNkHRGBurGn_iFfbL1tvNq58-V2GSVK_LwSLCt1nQkhwCak6HWUbfu83Yemecvz4UEdlodf2buiEI7Die4OHLmxClsUpug",
		"X-MetaScan-Org": "1078238259684835328",
		"Content-Type":   "application/json",
	}).SetResult(&result).Post(url)
	log.Println(res.StatusCode())
	if err != nil {
		log.Println(err.Error())
		log.Println("启动检查失败")
		return
	}
	if res.StatusCode() != 200 {
		log.Println(res.Error())
		log.Println(res)
		log.Println("调用失败")
		return
	}
	log.Println(result)
	log.Println("启动任务成功")
}

type Repo struct {
	Branch     string `json:"branch"`
	CommitHash string `json:"commit_hash"`
}
type StartTaskRes struct {
	Data TaskData `json:"data"`
}

type TaskData struct {
	Id          string       `json:"id"`
	TaskState   string       `json:"taskState"`
	EngineTasks []BaseEntity `json:"engineTasks"`
}

type BaseEntity struct {
	Id string `json:"id"`
}

type TaskScan struct {
	SubPath      string `json:"sub_path"`
	Mode         string `json:"mode"`
	IgnoredPaths string `json:"ignored_paths"`
}

type TaskEnv struct {
	Node           string `json:"node"`
	Solc           string `json:"solc"`
	PackageManage  string `json:"package_manage"`
	CompileCommand string `json:"compile_command"`
	Variables      string `json:"variables"`
}

func QueryTaskStatus() {
	url := "https://app.metatrust.io/api/scan/state"
	result := struct {
		Data TaskStatusRes `json:"data"`
	}{}
	res, err := utils.NewHttp().NewRequest().SetQueryParam("taskId", "1098610494056431618").SetHeaders(map[string]string{
		"Authorization":  "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJiYXdqN2JZQjdHX2MtVDJXNmFiQkIwMHZld2xoaHZLVVNfSXJUTDFBdUs4In0.eyJleHAiOjE2ODE5NzE2NjEsImlhdCI6MTY4MTk2OTg2MSwianRpIjoiZmNkYmU0ZTAtYjg1Ni00NzEyLWFmZjAtZDgxYjgxNjkwNjNmIiwiaXNzIjoiaHR0cHM6Ly9hY2NvdW50Lm1ldGF0cnVzdC5pby9yZWFsbXMvbXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiMjEzNjdlMGQtYWQ0NC00YTMwLWI4OWUtMDRmNDM2NWE4ZmM3IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2ViYXBwIiwic2Vzc2lvbl9zdGF0ZSI6ImYwYjNhZmE3LTkxZDYtNDg2ZC05YzY0LTAzMGEzMTY5YzkzYyIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOlsiaHR0cHM6Ly9hcHAubWV0YXRydXN0LmlvIiwiaHR0cHM6Ly9tZXRhdHJ1c3QuaW8iLCJodHRwczovL3d3dy5tZXRhdHJ1c3QuaW8iXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtbXQiLCJvZmZsaW5lX2FjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwiZGVsZXRlLWFjY291bnQiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJzaWQiOiJmMGIzYWZhNy05MWQ2LTQ4NmQtOWM2NC0wMzBhMzE2OWM5M2MiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicHJlZmVycmVkX3VzZXJuYW1lIjoidG9tQGhhbXN0ZXJuZXQuaW8iLCJlbWFpbCI6InRvbUBoYW1zdGVybmV0LmlvIiwidXNlcm5hbWUiOiJ0b21AaGFtc3Rlcm5ldC5pbyJ9.cArWpbG1csPLdIQoSRhtD8nH4nhsLPXaT4Jp58Nl1HwjK2griXRA-jwN2X6Y5gXaVE2OV1sdCeN0ldm2IxuqGs6CFxB6ehi3w_1bRzpVDYy3AmdrRiECXve3v0ltQTeSu8gtzM6eDaPGo0up-b3NuyppOxsSwsOHbmP2x-Y_yDfZ9Ia381wQf1KjzsiXHnhkaWyC6KetRgfz2Q2geoCMh5O9YqwQ-clM8LWeVHo5bMNkewvU8z-72ZAMckp7pAI0FanVX2BPE2jzRJ3n8ccIIfT5uoxBMiblCbeXF0UZFukVAk6Zse-xAt5AHtoM2iPFzVSZYBMvLgOhD59lLtR9Yg",
		"X-MetaScan-Org": "1078238259684835328",
	}).SetResult(&result).Get(url)
	log.Println(res.StatusCode())
	if err != nil {
		log.Println(err.Error())
		log.Println("查询任务状态失败")
		return
	}
	if res.StatusCode() == 401 {
		log.Println("没有权限")
		return
	}
	if res.StatusCode() != 200 {
		log.Println(res.Error())
		log.Println(res)
		log.Println("查询任务状态失败")
		return
	}
	log.Println(result.Data.State)
}

type TaskStatusRes struct {
	State string `json:"state"`
}

func GetEngineTaskSummary() {
	url := "https://app.metatrust.io/api/scan/engineTask/{engineTaskId}"
	result := struct {
		Data SummaryData `json:"data"`
	}{}
	res, err := utils.NewHttp().NewRequest().SetPathParam("engineTaskId", "1098616244203945984").SetHeaders(map[string]string{
		"Authorization":  "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJiYXdqN2JZQjdHX2MtVDJXNmFiQkIwMHZld2xoaHZLVVNfSXJUTDFBdUs4In0.eyJleHAiOjE2ODE5NzcxMTEsImlhdCI6MTY4MTk3NTMxMSwianRpIjoiODI4MmE4YzctYWQ4Ny00ZTEyLTk2YTktOGUyMWE4NjdhOTAzIiwiaXNzIjoiaHR0cHM6Ly9hY2NvdW50Lm1ldGF0cnVzdC5pby9yZWFsbXMvbXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiMjEzNjdlMGQtYWQ0NC00YTMwLWI4OWUtMDRmNDM2NWE4ZmM3IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2ViYXBwIiwic2Vzc2lvbl9zdGF0ZSI6IjE5NTZjYjY1LWE3M2UtNDk1My1iOTNiLTgwNmQ0YTUyNDc5NiIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOlsiaHR0cHM6Ly9hcHAubWV0YXRydXN0LmlvIiwiaHR0cHM6Ly9tZXRhdHJ1c3QuaW8iLCJodHRwczovL3d3dy5tZXRhdHJ1c3QuaW8iXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtbXQiLCJvZmZsaW5lX2FjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwiZGVsZXRlLWFjY291bnQiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJzaWQiOiIxOTU2Y2I2NS1hNzNlLTQ5NTMtYjkzYi04MDZkNGE1MjQ3OTYiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicHJlZmVycmVkX3VzZXJuYW1lIjoidG9tQGhhbXN0ZXJuZXQuaW8iLCJlbWFpbCI6InRvbUBoYW1zdGVybmV0LmlvIiwidXNlcm5hbWUiOiJ0b21AaGFtc3Rlcm5ldC5pbyJ9.kGm_sYXbNtGz8TmQqA2GcTNuglRFRrP5hyJ_OsV2OaHJr_MfxMtsArWHMCB95zfBkLaiTl2ZQ-HwDWyYNwz9Vxxg1vDHIYEgIOLUEfXTrVMzqKiKF-R0U2sggkoBoLFlSAfaS5M7k7mDMprQvgVVkTGduc_crdYLk7YNW0NW5KhvujCg-8ksIf_2HZ-UBI7TnrsE5noYb7lOiXZunsDGAB_CXNqDCguYa4U5h8BPtbczgfsZmwc2l5SxR5u7rYyjJMmV1Rk_RB5XfPZMyRGWtfDD_j9jw2GbEBdXxfZ3fLgSw4xFkZ2jAV-V4u9KfKkKIMc5PSq9deyGu5Cy-AOx9Q",
		"X-MetaScan-Org": "1078238259684835328",
	}).SetResult(&result).Get(url)
	log.Println(res.StatusCode())
	if err != nil {
		log.Println(err.Error())
		log.Println("查询任务状态失败")
		return
	}
	if res.StatusCode() == 401 {
		log.Println("没有权限")
		return
	}
	if res.StatusCode() != 200 {
		log.Println(res.Error())
		log.Println(res)
		log.Println("查询任务状态失败")
		return
	}
	log.Println(result)

}

type SummaryData struct {
	ResultOverview ResultOverview `json:"resultOverview"`
}

type ResultOverview struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Impact  Impact `json:"impact"`
}

type Impact struct {
	Critical      int `json:"CRITICAL"`
	Low           int `json:"LOW"`
	High          int `json:"HIGH"`
	Medium        int `json:"MEDIUM"`
	Informational int `json:"INFORMATIONAL"`
}

func GetTaskResult() {
	url := "https://app.metatrust.io/api/scan/history/engine/{engineTaskId}/result"
	result := struct {
		Data TaskResult `json:"data"`
	}{}
	res, err := utils.NewHttp().NewRequest().SetPathParam("engineTaskId", "1098923611277754368").SetHeaders(map[string]string{
		"Authorization":  "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJiYXdqN2JZQjdHX2MtVDJXNmFiQkIwMHZld2xoaHZLVVNfSXJUTDFBdUs4In0.eyJleHAiOjE2ODIwNzIxMjMsImlhdCI6MTY4MjA3MDMyMywianRpIjoiOTA4NjE5NzgtNjgyZC00MDJhLTkxYTQtZWNlNDZiZDJkZDdlIiwiaXNzIjoiaHR0cHM6Ly9hY2NvdW50Lm1ldGF0cnVzdC5pby9yZWFsbXMvbXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiMjEzNjdlMGQtYWQ0NC00YTMwLWI4OWUtMDRmNDM2NWE4ZmM3IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2ViYXBwIiwic2Vzc2lvbl9zdGF0ZSI6IjEyNjg4ZjkxLTFhNDQtNGIwZS1iMjQzLTdjNDg3OTIxZjE2NSIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOlsiaHR0cHM6Ly9hcHAubWV0YXRydXN0LmlvIiwiaHR0cHM6Ly9tZXRhdHJ1c3QuaW8iLCJodHRwczovL3d3dy5tZXRhdHJ1c3QuaW8iXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtbXQiLCJvZmZsaW5lX2FjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwiZGVsZXRlLWFjY291bnQiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJzaWQiOiIxMjY4OGY5MS0xYTQ0LTRiMGUtYjI0My03YzQ4NzkyMWYxNjUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicHJlZmVycmVkX3VzZXJuYW1lIjoidG9tQGhhbXN0ZXJuZXQuaW8iLCJlbWFpbCI6InRvbUBoYW1zdGVybmV0LmlvIiwidXNlcm5hbWUiOiJ0b21AaGFtc3Rlcm5ldC5pbyJ9.r_jaYmHjlHHDN2e93pHF9OKhUNpdZuZv5lUOrjlEWGtY0VsR2KIVu0SZVow0ygB6BatmKo10gdZliFGBl5mqbYjPhcvpmc8QRNXXJt2E80k9wc4gL1wtUWkds3wrBVDNpQ4PoxOvAIupPOKPLeA6R1OrnGsFgZBXy34ybc8gcTUGjNeuuWHTs6efdFhkFs7kX0LE1FnN6827LfL-Igi5XMVKcTpeJZhMTr-Mb4yGsZCtXZt_MSIlkvcbBE44jgNRB4eaCEGCbiagHjPe5ZFejZ8Q-Hf8gjkRxRx4x3uBxAHyjJgbrdhwilV4RALJT0w8AMrzPyJoG2JrtrSsGK2tDw",
		"X-MetaScan-Org": "1098616244203945984",
	}).SetResult(&result).Get(url)
	log.Println(res.StatusCode())
	if err != nil {
		log.Println(err.Error())
		log.Println("查询任务状态失败")
		return
	}
	if res.StatusCode() == 401 {
		log.Println("没有权限")
		return
	}
	if res.StatusCode() != 200 {
		log.Println(res.Error())
		log.Println(res)
		log.Println("查询任务状态失败")
		return
	}
	var resultData SecurityAnalyzerResponse
	if err = json.Unmarshal([]byte(result.Data.Result), &resultData); err != nil {
		log.Println("json 格式化有问题")
	}
	log.Println(resultData.FileMapping)

}

type TaskResult struct {
	Title  string `json:"title"`
	Result string `json:"result"`
}

func GetFile() {
	url := "https://app.metatrust.io/api/scan/history/vulnerability-files/8f82c7ad-cacd-418a-bcd2-cbb4218c3f86/Functions.sol"
	res, err := utils.NewHttp().NewRequest().SetHeaders(map[string]string{
		"Authorization":  "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJiYXdqN2JZQjdHX2MtVDJXNmFiQkIwMHZld2xoaHZLVVNfSXJUTDFBdUs4In0.eyJleHAiOjE2ODIwNzIxMjMsImlhdCI6MTY4MjA3MDMyMywianRpIjoiOTA4NjE5NzgtNjgyZC00MDJhLTkxYTQtZWNlNDZiZDJkZDdlIiwiaXNzIjoiaHR0cHM6Ly9hY2NvdW50Lm1ldGF0cnVzdC5pby9yZWFsbXMvbXQiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiMjEzNjdlMGQtYWQ0NC00YTMwLWI4OWUtMDRmNDM2NWE4ZmM3IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoid2ViYXBwIiwic2Vzc2lvbl9zdGF0ZSI6IjEyNjg4ZjkxLTFhNDQtNGIwZS1iMjQzLTdjNDg3OTIxZjE2NSIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOlsiaHR0cHM6Ly9hcHAubWV0YXRydXN0LmlvIiwiaHR0cHM6Ly9tZXRhdHJ1c3QuaW8iLCJodHRwczovL3d3dy5tZXRhdHJ1c3QuaW8iXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtbXQiLCJvZmZsaW5lX2FjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwiZGVsZXRlLWFjY291bnQiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJzaWQiOiIxMjY4OGY5MS0xYTQ0LTRiMGUtYjI0My03YzQ4NzkyMWYxNjUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicHJlZmVycmVkX3VzZXJuYW1lIjoidG9tQGhhbXN0ZXJuZXQuaW8iLCJlbWFpbCI6InRvbUBoYW1zdGVybmV0LmlvIiwidXNlcm5hbWUiOiJ0b21AaGFtc3Rlcm5ldC5pbyJ9.r_jaYmHjlHHDN2e93pHF9OKhUNpdZuZv5lUOrjlEWGtY0VsR2KIVu0SZVow0ygB6BatmKo10gdZliFGBl5mqbYjPhcvpmc8QRNXXJt2E80k9wc4gL1wtUWkds3wrBVDNpQ4PoxOvAIupPOKPLeA6R1OrnGsFgZBXy34ybc8gcTUGjNeuuWHTs6efdFhkFs7kX0LE1FnN6827LfL-Igi5XMVKcTpeJZhMTr-Mb4yGsZCtXZt_MSIlkvcbBE44jgNRB4eaCEGCbiagHjPe5ZFejZ8Q-Hf8gjkRxRx4x3uBxAHyjJgbrdhwilV4RALJT0w8AMrzPyJoG2JrtrSsGK2tDw",
		"X-MetaScan-Org": "1098616244203945984",
	}).Get(url)
	if err != nil {
		log.Println("获取失败")
		return
	}
	log.Println(res.StatusCode())
	log.Println(string(res.Body()))
}
