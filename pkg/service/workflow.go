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
	err := w.db.Where(&db.Workflow{
		ProjectId: projectId,
		Type:      workflowType,
	}).First(&workflow).Error
	if err != nil {
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
	if workflow.Tool != "" {
		err = w.engine.CreateJob(workflowKey, workflow.ExecFile)
		if err != nil {
			return deployResult, err
		}
	}
	met, token := setMetaScanToken(workflow)
	if met {
		params["scanToken"] = token
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
		workflow, err = w.SaveWorkflow(settingData)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	checkKey := w.GetWorkflowKey(settingData.ProjectId.String(), workflow.Id)
	file, err := w.TemplateParseV2(checkKey, settingData.Tool, projectData)
	if err != nil {
		return err
	}
	workflow.ExecFile = file
	toolType := hasCommonElements(settingData.Tool, consts.MetaScanTool)
	if toolType {
		workflow.ToolType = 1
	} else {
		workflow.ToolType = 0
	}
	workflow.Tool = strings.Join(settingData.Tool, ",")
	w.UpdateWorkflow(workflow)
	return nil
}

func (w *WorkflowService) WorkflowSettingCheck(projectId string, workflowType consts.WorkflowType) map[string][]string {
	var workflow db.Workflow
	var data []string
	result := make(map[string][]string)
	err := w.db.Model(db.Workflow{}).Where("project_id=? and type=?", projectId, workflowType).First(&workflow).Error
	if err == gorm.ErrRecordNotFound {
		return result
	}
	data = strings.Split(workflow.Tool, ",")
	if len(data) > 0 {
		for _, i2 := range data {
			switch i2 {
			case "MetaTrust (SA)", "MetaTrust (SP)", "Mythril":
				result["Security Analysis"] = append(result["Security Analysis"], i2)
			case "MetaTrust (OSA)":
				result["Open Source Analysis"] = append(result["Open Source Analysis"], i2)
			case "Solhint", "MetaTrust (CQ)":
				result["Code Quality Analysis"] = append(result["Code Quality Analysis"], i2)
			case "eth-gas-reporter":
				result["Gas Usage Analysis"] = append(result["Gas Usage Analysis"], i2)
			default:
				result["Other Analysis"] = append(result["Other Analysis"], i2)
			}
		}
	}
	return result
}

func (w *WorkflowService) UpdateWorkflow(data db.Workflow) error {
	res := w.db.Save(&data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func getCheckTemplate() string {
	filePath := "templates/metascan-check.yml"
	return filePath
}

func getTemplate(project *vo.ProjectDetailVo, workflowType consts.WorkflowType) string {
	filePath := "templates/truffle-build.yml"
	if project.Type == uint(consts.CONTRACT) {
		if workflowType == consts.Check {

			if project.FrameType == consts.Sui {
				filePath = "templates/sui-check.yml"
			} else if project.FrameType == consts.Aptos {
				filePath = "templates/aptos-check.yml"
			} else {
				filePath = "templates/truffle_check.yml"
			}
		} else if workflowType == consts.Build {
			if project.FrameType == consts.StarkWare {
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

func (w *WorkflowService) TemplateParseV2(name string, tool []string, project *vo.ProjectDetailVo) (string, error) {
	if project == nil {
		return "", errors.New("project is nil")
	}
	filePath := getCheckTemplate()
	content, err := temp.ReadFile(filePath)
	if err != nil {
		log.Println("read template file failed ", err.Error())
		return "", err
	}
	fileContent := string(content)
	tmpl := template.New("test")
	tmpl = tmpl.Delims("[[", "]]")
	var checkType []string
	metaCheck := hasCommonElements(tool, consts.MetaScanTool)
	if metaCheck {
		checkType = append(checkType, "CheckMetaScan")
	}
	truffleCheck := hasCommonElements(tool, consts.TruffleCheckTool)
	if truffleCheck {
		checkType = append(checkType, "Truffle Check")
	}
	templateData := parameter.MetaScanCheck{
		Name:          name,
		CheckType:     checkType,
		Tool:          tool,
		RepositoryUrl: project.RepositoryUrl,
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
	if workflow.ToolType == 1 {

	}
	switch workflow.ToolType {
	case 1:
		metaScanFlag = true
		token = utils.MetaScanHttpRequestToken()
		log.Println(fmt.Sprintf("flag is %t,token is :%s", metaScanFlag, token))
	default:
		metaScanFlag = false
		token = ""
	}
	log.Println(fmt.Sprintf("flag is %t,token is :%s", metaScanFlag, token))
	return metaScanFlag, token
}

func hasCommonElements(arr1, arr2 []string) bool {
	for _, elem1 := range arr1 {
		for _, elem2 := range arr2 {
			if elem1 == elem2 {
				return true
			}
		}
	}
	return false
}
