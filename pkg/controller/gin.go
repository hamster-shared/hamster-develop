package controller

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"gorm.io/gorm"
)

// go:embed dist
// var content embed.FS

type HttpServer struct {
	handlerServer HandlerServer
	port          int
}

func NewHttpService(handlerServer HandlerServer, port int) *HttpServer {
	return &HttpServer{
		handlerServer: handlerServer,
		port:          port,
	}
}

func (h *HttpServer) StartHttpServer() {
	r := gin.Default()
	api := r.Group("/api")
	api.Use(h.handlerServer.RequestLog())

	api.POST("/login", h.handlerServer.loginWithGithub)
	api.POST("/v2/login", h.handlerServer.githubLogin)
	api.POST("/metamask/login", h.handlerServer.metamaskLogin)
	api.POST("/github/install", h.handlerServer.githubInstall)
	api.POST("/repo/authorization", h.handlerServer.githubRepoAuth)
	api.POST("/github/webhook", h.handlerServer.githubWebHook)
	api.POST("/v2/github/webhook", h.handlerServer.githubWebHookV2)
	api.GET("/projects/:id/:username/frontend/logs", h.handlerServer.getDeployFrontendLog)
	api.GET("/user/count", h.handlerServer.getUserCount)
	// project_template
	api.GET("/templates-category", h.handlerServer.templatesCategory)
	api.GET("/templates-category/:id/templates", h.handlerServer.templates)
	api.GET("/templates/:id", h.handlerServer.templateDetail)
	api.GET("/frontend-templates/:id", h.handlerServer.frontendTemplateDetail)
	api.GET("/chain-templates/:id", h.handlerServer.chainTemplateDetail)
	api.GET("/templates/show", h.handlerServer.templateShow)
	api.POST("/templates/:id/download", h.handlerServer.templateDownload)
	api.Use(h.handlerServer.Authorize())
	//api.Use(h.handlerServer.JwtAuthorize())
	api.POST("/user/wallet", h.handlerServer.saveUserWallet)
	// project
	api.GET("/projects", h.handlerServer.projectList)
	api.POST("/projects", h.handlerServer.createProject)
	api.POST("/projects/import", h.handlerServer.importProject)
	api.POST("/projects/code", h.handlerServer.createProjectByCode)
	api.GET("/projects/:id", h.handlerServer.projectDetail)
	api.PUT("/projects/:id", h.handlerServer.updateProject)
	api.DELETE("projects/:id", h.handlerServer.deleteProject)
	api.POST("/projects/check-name", h.handlerServer.checkName)
	api.GET("/user", h.handlerServer.getUseInfo)
	api.GET("/v2/user", h.handlerServer.getUseInfoV2)
	api.PUT("/user/first/state", h.handlerServer.updateFirstStateV2)
	api.PUT("/v2/user/first/state", h.handlerServer.updateFirstStateV2)

	// new add
	api.GET("/github/install", h.handlerServer.githubInstallCheck)
	api.POST("/github/install/auth", h.handlerServer.githubInstallAuth)
	api.POST("/v2/github/install", h.handlerServer.githubInstallV2)
	api.GET("/github/users/installations", h.handlerServer.getUsersInstallations)
	api.GET("/github/installation/:id/repositories", h.handlerServer.getGithubRepos)
	// set check pipeline
	api.POST("/project/:id/workflow/setting", h.handlerServer.workflowSetting)
	api.GET("/project/:id/workflow/setting/check", h.handlerServer.workflowSettingCheck)
	// repository
	api.GET("/repositories", h.handlerServer.getRepositories)
	api.GET("/repository/type", h.handlerServer.repositoryType)
	api.GET("/repository/selection", h.handlerServer.getGitHubRepositorySelection)
	/*
		创建项目返回项目 ID
		缺登录的接口
		模版详情接口缺少返回字段
		缺少仓库校验接口（根据 project 名称）
		删除 deploy 接口
		保存部署信息传参数有问题
		Workflow 详情返回字段缺少流水线类型（check, build）
		查询 workflow 下的合约列表使用 workflowDetailId
		缺少项目日志接口使用 ID
		获取合约列表改为项目合约列表（改名字）
		合约部署详情接口有问题
		根据版本查询合约信息（返回 abi 信息和 byte code）
	*/
	api.POST("/projects/:id/workflows/:workflowId/detail/:detailId/stop", h.handlerServer.stopWorkflow)
	api.POST("/projects/:id/check", h.handlerServer.projectWorkflowCheck)
	api.POST("/projects/:id/build", h.handlerServer.projectWorkflowBuild)
	// frontend deploy
	api.POST("/projects/:id/workflows/:workflowId/detail/:detailId/deploy", h.handlerServer.projectWorkflowDeploy)
	api.POST("/projects/:id/workflows/:workflowId/detail/:detailId/container/deploy", h.handlerServer.containerDeploy)
	api.GET("/projects/:id/workflows/:workflowId/container/check", h.handlerServer.configContainerDeploy)
	api.POST("/projects/:id/container/deploy", h.handlerServer.updateContainerDeploy)
	api.GET("/projects/:id/container/deploy", h.handlerServer.getContainerDeploy)
	api.GET("/projects/:id/contract", h.handlerServer.projectContract)
	api.GET("/projects/:id/reports", h.handlerServer.projectReport)

	// aptos params
	api.GET("/projects/:id/params/aptos", h.handlerServer.queryAptosParams)
	api.GET("/projects/:id/is-needs-params/aptos", h.handlerServer.isAptosNeedsParams)
	api.POST("/projects/:id/params/aptos", h.handlerServer.saveAptosParams)
	api.GET("/projects/:id/aptos-build", h.handlerServer.projectWorkflowAptosBuild)

	api.GET("/projects/:id/frontend/reports", h.handlerServer.projectFrontendReports)
	api.GET("/projects/:id/packages", h.handlerServer.projectPackages)
	api.POST("/projects/:id/contract/deploying", h.handlerServer.saveContractDeployIngInfo)
	api.POST("/projects/:id/contract/deploy", h.handlerServer.saveContractDeployInfo)
	api.GET("/projects/:id/contract/deploy/:contractDeployId", h.handlerServer.contractDeployInfo)

	//workflow
	api.GET("/projects/:id/workflows", h.handlerServer.workflowList)
	api.DELETE("/workflows/:workflowId/detail/:detailId", h.handlerServer.deleteWorkflow)
	api.GET("/workflows/:id/detail/:detailId", h.handlerServer.workflowDetail)
	api.GET("/projects/:id/workflows/:workflowId/detail/:detailId/contract", h.handlerServer.workflowContract)
	//delete frontend deploy
	api.DELETE("/package/:id/deploy-info", h.handlerServer.deleteWorkflowDeploy)
	api.GET("/workflows/:id/detail/:detailId/report", h.handlerServer.workflowReport)
	api.GET("/workflows/:id/detail/:detailId/report/overview", h.handlerServer.workflowReportOverview)
	api.GET("/report/:id", h.handlerServer.reportDetail)
	api.GET("/metascan/file/*key", h.handlerServer.metaScanFile)
	api.GET("/project/:id/contract/content/*name", h.handlerServer.contractFileContent)
	// frontend reports
	api.GET("/workflows/:id/detail/:detailId/frontend/report", h.handlerServer.workflowFrontendReports)
	//workflow frontend packages
	api.GET("/workflows/:id/detail/:detailId/package", h.handlerServer.workflowFrontendPackage)
	api.GET("/workflows/:id/detail/:detailId/deploy-info", h.handlerServer.workflowFrontendDeployInfo)
	//deploy detail
	api.GET("/package/:id/deploy/detail", h.handlerServer.workflowFrontendPackageDetail)

	//contract
	api.GET("/projects/:id/contract/:version", h.handlerServer.contractInfo)
	api.GET("/projects/:id/contract/deploy/detail", h.handlerServer.contractDeployDetailByVersion)
	api.GET("/projects/:id/versions", h.handlerServer.versionList)
	api.GET("/projects/:id/contract/name", h.handlerServer.queryContractNameList)
	api.GET("/projects/:id/contract/network", h.handlerServer.queryNetworkList)
	api.GET("/projects/:id/check-tools", h.handlerServer.queryReportCheckTools)

	//logs
	api.GET("/workflows/:id/detail/:detailId/logs", h.handlerServer.getWorkflowLog)
	api.GET("/workflows/:id/detail/:detailId/logs/:stageName", h.handlerServer.getWorkflowStageLog)

	//ICP-api
	api.GET("/projects/:id/dfx-json", h.handlerServer.getDfxJsonData)
	api.GET("/projects/:id/dfx-json/check", h.handlerServer.isConfigJsonData)
	api.PUT("/projects/:id/dfx-json/:dfxId", h.handlerServer.updateDfxJsonData)
	api.POST("/projects/:id/dfx-json", h.handlerServer.saveDfxJsonData)
	api.GET("/projects/:id/canister", h.handlerServer.getCanisterInfo)
	api.GET("/projects/:id/canister/list", h.handlerServer.getCanisterList)       //罐列表
	api.GET("/projects/:id/icp/account", h.handlerServer.getIcpAccount)           //获取icp的账户信息
	api.POST("/projects/:id/recharge/canister", h.handlerServer.rechargeCanister) //给应用罐罐充值cycles
	api.POST("/projects/:id/recharge/wallet", h.handlerServer.rechargeWallet)     //充值钱包
	api.GET("/projects/:id/wallet/info", h.handlerServer.getWalletInfo)           //查询钱包罐信息
	api.GET("/projects/:id/account/info", h.handlerServer.getAccountInfo)         //查询icp信息accountId和余额
	api.POST("/projects/:id/create/identity", h.handlerServer.createIdentity)     //创建icp身份
	api.POST("/projects/:id/redeem/coupon", h.handlerServer.redeemFaucetCoupon)   //通过优惠卷生成钱包罐

	//arrange-api
	api.GET("/projects/:id/version/info", h.handlerServer.getVersionAndCodeInfoList)                   //获取版本列表和codeInfo
	api.POST("/projects/:id/arrange", h.handlerServer.saveContractArrange)                             //保存编排信息
	api.PUT("/projects/:id/arrange", h.handlerServer.updateContractArrange)                            //更新编排信息
	api.POST("/projects/:id/arrange/execute", h.handlerServer.saveContractArrangeExecute)              //保存编排执行信息
	api.PUT("/projects/:id/arrange/execute", h.handlerServer.updateContractArrangeExecute)             //更新执行信息
	api.GET("/projects/:id/arrange/execute/:executeId", h.handlerServer.getContractArrangeExecuteInfo) //根据执行id获取执行信息
	api.GET("/projects/:id/arrange/contract/:version", h.handlerServer.getDeployArrangeContractList)   //根据项目ID和版本获取部署的合约列表
	//
	//保存部署成的合约信息
	//
	//查询某个那版本下已经部署的合约

	// ======== old api =========//
	// pipeline
	//create pipeline job
	api.POST("/pipeline", h.handlerServer.createPipeline)
	//update pipeline job
	api.PUT("/pipeline/:oldName", h.handlerServer.updatePipeline)
	//get pipeline job
	api.GET("/pipeline/:name", h.handlerServer.getPipeline)
	//delete pipeline job and pipeline job detail
	api.DELETE("/pipeline/:name", h.handlerServer.deletePipeline)
	//get pipeline job list
	api.GET("/pipeline", h.handlerServer.pipelineList)
	//get pipeline job detail info
	api.GET("/pipeline/:name/detail/:id", h.handlerServer.getPipelineDetail)
	//delete pipeline job detail
	api.DELETE("/pipeline/:name/detail/:id", h.handlerServer.deleteJobDetail)
	//get pipeline job detail list
	api.GET("/pipeline/:name/details", h.handlerServer.getPipelineDetailList)
	//exec pipeline job
	api.POST("/pipeline/:name/exec", h.handlerServer.execPipeline)
	//re exec pipeline detail job
	api.POST("/pipeline/:name/:id/re-exec", h.handlerServer.reExecuteJob)
	//stop pipeline job
	api.POST("/pipeline/:name/:id/stop", h.handlerServer.stopJobDetail)
	api.GET("/pipeline/:name/logs/:id", h.handlerServer.getJobLog)
	api.GET("/pipeline/:name/logs/:id/:stagename", h.handlerServer.getJobStageLog)
	// get template list
	//api.GET("/pipeline/templates", h.handlerServer.getTemplates)
	//api.GET("/pipeline/template/:id", h.handlerServer.getTemplateDetail)
	api.GET("/pipeline/:name/detail/:id/artifactory", h.handlerServer.openArtifactoryDir)
	api.GET("/ping", func(c *gin.Context) {

		db := application.GetBean[*gorm.DB]("db")
		fmt.Println(db)
		//输出 json 结果给调用方
		Success("", c)
	})

	api.GET("/workflows/:id/detail/:detailId/logs/:stageName/:stepName", h.handlerServer.getWorkflowStepLog)
	// 下载文件
	api.GET("/download", h.handlerServer.download)

	// fe, _ := fs.Sub(content, "dist")
	// r.NoRoute(gin.WrapH(http.FileServer(http.FS(fe))))
	r.Run(fmt.Sprintf(":%d", h.port)) // listen and serve on
}

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func OpenWeb(port int) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/c", fmt.Sprintf("start http://127.0.0.1:%d", port))
	} else {
		cmd = exec.Command(run, fmt.Sprintf("http://127.0.0.1:%d", port))
	}
	return cmd.Start()
}
