package service

import (
	"bufio"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

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
	"github.com/mohaijiang/agent-go/candid"
	"gorm.io/gorm"
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
	if uint(consts.FRONTEND) == projectData.Type || uint(consts.BLOCKCHAIN) == projectData.Type {
		jobDetail, err := w.engine.GetJobHistory(message.JobName, message.JobId)
		if err != nil {
			logger.Errorf("get job history failed: %s", err)
			return
		}
		if uint(consts.Build) == workflowDetail.Type {
			w.syncFrontendBuild(jobDetail, workflowDetail, projectData)
		} else if uint(consts.Deploy) == workflowDetail.Type {
			w.syncFrontendDeploy(jobDetail, workflowDetail, projectData)
		}
	}
}

func (w *WorkflowService) syncFrontendBuild(detail *model.JobDetail, workflowDetail db.WorkflowDetail, project db.Project) {
	if len(detail.ActionResult.Artifactorys) > 0 {
		projectName := project.Name
		if project.Type == uint(consts.BLOCKCHAIN) {
			projectName = fmt.Sprintf("%s_node_polkadot", project.Name)
		} else if project.Type == uint(consts.FRONTEND) && project.DeployType == int(consts.INTERNET_COMPUTER) {
			if project.FrameType == 1 {
				projectName = fmt.Sprintf("%s_%s_ic", project.Name, "vuejs")
			} else if project.FrameType == 2 {
				projectName = fmt.Sprintf("%s_%s_ic", project.Name, "reactjs")
			}
		}
		for range detail.ActionResult.Artifactorys {
			frontendPackage := db.FrontendPackage{
				ProjectId:        workflowDetail.ProjectId,
				WorkflowId:       workflowDetail.WorkflowId,
				WorkflowDetailId: workflowDetail.Id,
				Name:             projectName,
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
		if project.Type == uint(consts.FRONTEND) {
			if project.FrameType == 1 {
				image = "https://g.alpha.hamsternet.io/ipfs/QmdhtgKNuQn2aqkdTn4DxRQidSLmmtggtpaBgh8vuyVjxd"
			} else if project.FrameType == 2 {
				image = "https://g.alpha.hamsternet.io/ipfs/QmUcXKszQxwT21dnqf67vFxzBgFwT8iWEsxExeGA1oFf6N"
			} else if project.FrameType == 3 {
				image = "https://g.alpha.hamsternet.io/ipfs/QmW8DNyCUrvDHaG4a4aKjkDNTbYDy9kwFxhFno2nKmgTKt"
			} else {
				image = "https://g.alpha.hamsternet.io/ipfs/QmPsa61VtwQH3ixzZys7EF9VG1zV7LQHDYjEYBfZpnmPDy"
			}
		} else if project.Type == uint(consts.BLOCKCHAIN) {
			image = "https://g.alpha.hamsternet.io/ipfs/QmPbUjgPNW1eBVxh1zVgF9F7porBWijYrAeMth9QDPwEXk"
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
				packageDeploy.Name = data.Name
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

	projectService := application.GetBean[IProjectService]("projectService")
	project, err := projectService.GetProjectById(projectIdStr)

	if err != nil {
		logger.Errorf("get project fail : %s", projectIdStr)
		return
	}

	// 同步构建物
	if len(jobDetail.Artifactorys) > 0 {
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
		case consts.InternetComputer:
			err = w.syncInternetComputerBuild(projectId, workflowId, workflowDetail, jobDetail)
			return
		case consts.Solana:
			err = w.syncContractSolana(projectId, workflowId, workflowDetail, jobDetail.Artifactorys)
			return
		default:
			for _, arti := range jobDetail.Artifactorys {
				err = w.syncContractEvm(projectId, workflowId, workflowDetail, arti)
			}
		}
	}

	if len(jobDetail.Deploys) > 0 {
		if project.DeployType == int(consts.INTERNET_COMPUTER) || project.FrameType == consts.InternetComputer {
			err = w.syncInternetComputerDeploy(projectId, workflowId, workflowDetail, jobDetail)
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
		if len(jobDetail.MetaScanData) > 0 {
			logger.Info("start sysnc meta scan report--------")
			if len(jobDetail.MetaScanData) > 0 {
				for _, datum := range jobDetail.MetaScanData {
					report := db.Report{
						ProjectId:        projectId,
						WorkflowId:       workflowId,
						WorkflowDetailId: workflowDetail.Id,
						Name:             consts.MetaScanReportTypeMap[consts.CheckToolTypeMap[datum.Tool]],
						Type:             uint(consts.Check),
						CheckTool:        datum.Tool,
						Result:           "success",
						CheckTime:        time.Now(),
						ReportFile:       datum.CheckResult,
						CreateTime:       time.Now(),
						Issues:           int(datum.Total),
						ToolType:         consts.CheckToolTypeMap[datum.Tool],
						MetaScanOverview: datum.ResultOverview,
					}
					w.db.Create(&report)
				}
				logger.Info("end sysnc meta scan report--------")
			}
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
						Issues:     contractCheckResult.Total,
						ToolType:   consts.CheckToolTypeMap[contractCheckResult.Tool],
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
					CheckTool:        "AI",
					Result:           "success",
					CheckTime:        time.Now(),
					ReportFile:       string(report.Content),
					CreateTime:       time.Now(),
					ToolType:         5,
				}
				reportList = append(reportList, report)
			}
		}
		if len(reportList) > 0 {
			logger.Tracef("len(reportList): %d ", len(reportList))
			err = begin.Save(&reportList).Error
			if err != nil {
				logger.Errorf("Save report fail, err is %s", err.Error())
				// return
			}
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
	logger.Info(mv)
	logger.Info(len(mv))
	if len(mv) > 0 {
		for _, s := range mv {
			contract := db.Contract{
				ProjectId:        projectId,
				WorkflowId:       workflowId,
				WorkflowDetailId: workflowDetail.Id,
				Name:             strings.TrimSuffix(artis[s.Index].Name, path.Ext(artis[s.Index].Name)),
				Version:          fmt.Sprintf("%d", workflowDetail.ExecNumber),
				BuildTime:        workflowDetail.CreateTime,
				AbiInfo:          "",
				ByteCode:         byteCode,
				AptosMv:          s.Mv,
				CreateTime:       time.Now(),
				Type:             uint(consts.Aptos),
				Status:           consts.STATUS_SUCCESS,
			}
			err = w.saveContractToDatabase(&contract)
			if err != nil {
				logger.Errorf("save contract to database failed: %s", err.Error())
			}
		}
	}
	// logger.Tracef("aptos contract: %+v", contract)
	return nil
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

func (w *WorkflowService) syncContractSolana(projectId uuid.UUID, workflowId uint, workflowDetail db.WorkflowDetail, artis []model.Artifactory) error {
	var bytecode string
	var idlData string

	for _, arti := range artis {
		if arti.Name == "solana_nft_anchor.so" {
			data, err := os.ReadFile(arti.Url)
			if err != nil {
				logger.Errorf("read bytecode error : %v", err)
			}
			bytecode = base64.StdEncoding.EncodeToString(data)
		}

		if arti.Name == "solana_nft_anchor.json" {
			data, err := os.ReadFile(arti.Url)
			if err != nil {
				logger.Errorf("read bytecode error : %v", err)
			}
			idlData = string(data)
		}
	}

	contract := db.Contract{
		ProjectId:        projectId,
		WorkflowId:       workflowId,
		WorkflowDetailId: workflowDetail.Id,
		Name:             "solana_nft_anchor",
		Version:          fmt.Sprintf("%d", workflowDetail.ExecNumber),
		BuildTime:        workflowDetail.CreateTime,
		AbiInfo:          idlData,
		ByteCode:         bytecode,
		CreateTime:       time.Now(),
		Type:             uint(consts.Solana),
		Status:           consts.STATUS_SUCCESS,
	}

	return w.saveContractToDatabase(&contract)

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

type AptosBuildInfo struct {
	Mv    string
	Index int
}

func (w *WorkflowService) getAptosMvAndByteCode(artis []model.Artifactory) (arr []AptosBuildInfo, byteCode string, err error) {
	var mvs []AptosBuildInfo
	for i, arti := range artis {
		// 以 .bcs 结尾，认为是 byteCode
		if strings.HasSuffix(arti.Url, ".bcs") {
			byteCode, err = utils.FileToHexString(arti.Url)
			if err != nil {
				logger.Errorf("hex string failed: %s", err.Error())
				return mvs, "", err
			}
			continue
		}
		var data AptosBuildInfo
		if strings.HasSuffix(arti.Url, ".mv") {
			mv, err := utils.FileToHexString(arti.Url)
			if err != nil {
				logger.Errorf("hex string failed: %s", err.Error())
				return mvs, "", err
			}
			data.Mv = mv
			data.Index = i
			mvs = append(mvs, data)
			continue
		}
		logger.Warnf("aptos contract file name is not end with .bcs or .mv: %s", arti.Url)
	}
	return mvs, byteCode, nil
}

func (w *WorkflowService) syncInternetComputerBuild(projectId uuid.UUID, workflowId uint, workflowDetail db.WorkflowDetail, jobDetail *model.JobDetail) error {

	var abiInfo string
	for _, arti := range jobDetail.Artifactorys {
		if strings.HasSuffix(arti.Name, "did") {
			// analysis did
			didContent, err := readDid(arti.Url)
			if err != nil {
				return err
			}

			discription, err := candid.ParseDID([]byte(didContent))
			if err != nil {
				return err
			}

			bytes, err := json.Marshal(discription)
			if err != nil {
				return err
			}
			abiInfo = string(bytes)
		}
	}

	for _, arti := range jobDetail.Artifactorys {
		if strings.HasSuffix(arti.Name, "zip") {
			backendPackage := db.BackendPackage{
				ProjectId:        projectId,
				WorkflowId:       workflowId,
				WorkflowDetailId: workflowDetail.Id,
				Name:             arti.Name,
				Version:          fmt.Sprintf("%d", workflowDetail.ExecNumber),
				BuildTime:        workflowDetail.CreateTime,
				AbiInfo:          abiInfo,
				CreateTime:       time.Now(),
				Type:             consts.InternetComputer,
				Status:           consts.DEPLOY_STATUS_SUCCESS,
				Branch:           jobDetail.CodeInfo,
			}
			err := w.db.Save(&backendPackage).Error
			if err != nil {
				return err
			}
		}

	}

	return nil

}

func (w *WorkflowService) syncInternetComputerDeploy(projectId uuid.UUID, workflowId uint, workflowDetail db.WorkflowDetail, jobDetail *model.JobDetail) error {

	for _, deploy := range jobDetail.Deploys {
		var deployInfo db.BackendDeploy
		var deployPackage db.BackendPackage
		buildWorkflowDetailId := jobDetail.Parameter["buildWorkflowDetailId"]
		err := w.db.Model(&db.BackendPackage{}).Where("project_id = ? and workflow_detail_id = ?", projectId.String(), buildWorkflowDetailId).First(&deployPackage).Error
		if err != nil {
			continue
		}

		var buildWorkflowDetail db.WorkflowDetail
		err = w.db.Model(&db.WorkflowDetail{}).Where("id = ?", buildWorkflowDetailId).First(&buildWorkflowDetail).Error
		if err != nil {
			continue
		}

		version := buildWorkflowDetail.ExecNumber

		err = w.db.Model(&db.ContractDeploy{}).Where("contract_id = ? and project_id = ? and version = ?", deployPackage.Id, projectId.String(), version).First(&deployInfo).Error
		if err != nil {
			deployInfo = db.BackendDeploy{
				ProjectId: projectId,
				PackageId: deployPackage.Id,
				Version:   strconv.Itoa(int(version)),
			}
		}
		deployInfo.Type = consts.InternetComputer
		deployInfo.Status = consts.DEPLOY_STATUS_SUCCESS // deployed
		deployInfo.CreateTime = time.Now()

		deployInfo.AbiInfo = deployPackage.AbiInfo
		deployInfo.DeployTime = deployInfo.CreateTime
		icNetwork := os.Getenv("IC_NETWORK")
		if icNetwork == "" {
			icNetwork = "local"
		}
		deployInfo.WorkflowId = workflowId
		deployInfo.WorkflowDetailId = workflowDetail.Id
		deployInfo.Network = icNetwork
		//deploy.Name
		deployInfo.Name = deploy.Name
		canisterId := deploy.Cid
		deployInfo.Address = canisterId

		err = w.db.Save(&deployInfo).Error
		if err != nil {
			return err
		}

		var icpCanister db.IcpCanister

		// 使用First查询满足条件的第一条数据
		if err := w.db.Model(db.IcpCanister{}).Where("project_id = ? and canister_id = ?", projectId.String(), canisterId).First(&icpCanister).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Println("数据不存在")
				icpCanister.CanisterId = canisterId
				icpCanister.CreateTime = sql.NullTime{Time: time.Now(), Valid: true}
				icpCanister.ProjectId = projectId.String()
			} else {
				fmt.Println("查询数据时发生错误:", err)
				continue
			}
		}

		icpCanister.CanisterName = deploy.Name
		icpCanister.Status = db.Running
		icpCanister.Contract = strings.Join([]string{deployPackage.Name, deployPackage.Version}, "_#")
		icpCanister.Cycles = sql.NullString{Valid: false}
		icpCanister.UpdateTime = sql.NullTime{Time: time.Now(), Valid: true}
		if err := w.db.Save(&icpCanister).Error; err != nil {
			fmt.Println("保存数据时发生错误:", err)
			continue
		}

		deployPackage.Network = utils.RemoveDuplicatesAndJoin(deployPackage.Network+","+icNetwork, ",")
	}
	return nil
}

func readDid(filePath string) (string, error) {
	// 打开输入文件进行读取
	inputFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return "", err
	}
	defer inputFile.Close()

	// 打开输出文件进行写入
	var convertedContent string

	scanner := bufio.NewScanner(inputFile)
	var currentLine string

	for scanner.Scan() {
		line := scanner.Text()
		currentLine += line
		// 写入当前行
		convertedContent += currentLine
		currentLine = ""
	}

	// 写入最后一行
	if currentLine != "" {
		convertedContent += currentLine
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading input file:", scanner.Err())
		return "", err
	}

	fmt.Println("File formatting completed.")
	fmt.Println(convertedContent)

	return convertedContent, err
}
