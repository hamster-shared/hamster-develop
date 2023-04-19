package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	jober "github.com/hamster-shared/aline-engine/job"
	"github.com/hamster-shared/aline-engine/logger"
	"github.com/hamster-shared/aline-engine/model"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	"github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/utils"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	uuid "github.com/iris-contrib/go.uuid"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func (w *WorkflowService) SyncStatus(message model.StatusChangeMessage) {
	logger.Debugf("SyncStatus: %v", message)

	_, workflowId, err := GetProjectIdAndWorkflowIdByWorkflowKey(message.JobName)
	if err != nil {
		return
	}

	jobDetail, err := w.engine.GetJobHistory(message.JobName, message.JobId)
	if err != nil {
		logger.Errorf("get job history failed: %v", err)
		return
	}

	var workflowDetail db.WorkflowDetail

	w.db.Where(&db.WorkflowDetail{
		WorkflowId: uint(workflowId),
		ExecNumber: uint(jobDetail.Id),
	}).First(&workflowDetail)

	if workflowDetail.Id == 0 {
		logger.Errorf("workflowDetail.Id is 0")
		return
	}

	workflowDetail.Status = uint(jobDetail.Status)
	stageInfo, err := json.Marshal(jobDetail.Stages)
	if err != nil {
		logger.Errorf("stage info json marshal failed: %v", err)
		return
	}
	workflowDetail.StageInfo = string(stageInfo)
	workflowDetail.UpdateTime = time.Now()
	workflowDetail.CodeInfo, err = w.engine.GetCodeInfo(message.JobName, message.JobId)
	if err != nil {
		logger.Warnf("get code info failed: %v", err)
	}
	workflowDetail.Duration = jobDetail.Duration

	if workflowDetail.Status != uint(message.Status) {
		// 如果 detail 的状态和 message 的状态不一致，可能是因为 detail 是从文件读取的，读取时还没有保存最新的状态，以 message 的状态为准
		logger.Warnf("workflowDetail.Status(%d) != message.Status(%d), use message.Status", workflowDetail.Status, message.Status)
		workflowDetail.Status = uint(message.Status)
	}

	// retry 3 times
	for i := 0; i < 3; i++ {
		err := w.db.Model(&workflowDetail).Select("*").Updates(workflowDetail).Error
		if err != nil {
			logger.Errorf("save workflow detail status to database failed: %s", err)
		} else {
			logger.Infof("save workflow detail status to database success")
			break
		}
	}

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
		logger.Errorf("find project by id failed: %s", err.Error())
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

	projectService := application.GetBean[IProjectService]("projectService")
	project, err := projectService.GetProjectById(projectIdStr)

	if err != nil {
		logger.Errorf("get project fail : %s", projectIdStr)
		return
	}

	switch project.FrameType {
	case consts.StarkWare:
		err = w.syncContractStarknet(projectId, workflowId, workflowDetail, jobDetail.Artifactorys)
		return
	case consts.Aptos:
		err = w.syncContractAptos(projectId, workflowId, workflowDetail, jobDetail.Artifactorys)
		return
	case consts.Ton:
		return
	case consts.Sui:
		contractName := getSuiModuleName(project)
		err = w.syncContractSui(projectId, workflowId, workflowDetail, jobDetail.Artifactorys, contractName)
		return
	default:
		for _, arti := range jobDetail.Artifactorys {
			err = w.syncContractEvm(projectId, workflowId, workflowDetail, arti)
		}
	}

	if err != nil {
		logger.Errorf("sync contract error : %v", err)
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
					ReportFile:       string(report.Content),
					CreateTime:       time.Now(),
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

func (w *WorkflowService) syncContractStarknet(projectId uuid.UUID, workflowId uint, workflowDetail db.WorkflowDetail, artis []model.Artifactory) error {

	var arti model.Artifactory
	for _, a := range artis {
		// 如果是 starknet 合约
		if strings.HasSuffix(a.Url, "starknet.output.json") {
			arti = a
			continue
		}
	}
	if arti.Url == "" {
		return nil
	}

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

	// logger.Tracef("aptos contract: %+v", contract)
	return w.saveContractToDatabase(&contract)
}

func getSuiModuleName(project *db.Project) string {
	DEFAULT_RESULT := "SuiModule"

	if project == nil {
		return DEFAULT_RESULT
	}
	short := strings.TrimPrefix(project.RepositoryUrl, "https://github.com/")
	short = strings.TrimSuffix(short, ".git")
	moveUrl := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/Move.toml", short, project.Branch)

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	response, err := client.Get(moveUrl)
	if err != nil {
		return DEFAULT_RESULT
	}

	data, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	var config vo.Config
	_, err = toml.Decode(string(data), &config)

	if err != nil {
		return DEFAULT_RESULT
	}
	return config.Package.Name
}

func (w *WorkflowService) syncContractSui(projectId uuid.UUID, workflowId uint, workflowDetail db.WorkflowDetail, artis []model.Artifactory, contractName string) error {

	var bytecode string
	for _, arti := range artis {
		if arti.Name == "bytecode.json" {
			data, err := os.ReadFile(arti.Url)
			if err != nil {
				logger.Errorf("read bytecode error : %v", err)
			}
			bytecode = string(data)
		}
	}

	contract := db.Contract{
		ProjectId:        projectId,
		WorkflowId:       workflowId,
		WorkflowDetailId: workflowDetail.Id,
		Name:             contractName,
		Version:          fmt.Sprintf("%d", workflowDetail.ExecNumber),
		BuildTime:        workflowDetail.CreateTime,
		AbiInfo:          "",
		ByteCode:         bytecode,
		AptosMv:          "",
		CreateTime:       time.Now(),
		Type:             uint(consts.Sui),
		Status:           consts.STATUS_SUCCESS,
	}

	// logger.Tracef("aptos contract: %+v", contract)
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
		// 以 .bcs 结尾，认为是 byteCode
		if strings.HasSuffix(arti.Url, ".bcs") {
			byteCode, err = utils.FileToHexString(arti.Url)
			if err != nil {
				logger.Errorf("hex string failed: %s", err.Error())
				return "", "", err
			}
			continue
		}
		if strings.HasSuffix(arti.Url, ".mv") {
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
