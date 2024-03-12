package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	engine "github.com/hamster-shared/aline-engine"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/service"
	uuid "github.com/iris-contrib/go.uuid"
	"strconv"
	"strings"
)

func (h *HandlerServer) workflowList(gin *gin.Context) {
	id := gin.Param("id")
	pageStr := gin.DefaultQuery("page", "1")
	sizeStr := gin.DefaultQuery("size", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	typeStr := gin.DefaultQuery("type", "0")
	workflowType, err := strconv.Atoi(typeStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	data, err := workflowService.GetWorkflowList(id, workflowType, page, size)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) workflowDetail(gin *gin.Context) {
	idStr := gin.Param("id")
	workflowId, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	detailIdStr := gin.Param("detailId")
	detailId, err := strconv.Atoi(detailIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	engineType := gin.DefaultQuery("engine", consts.EngineTypeWorkflow)
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	data, err := workflowService.GetWorkflowDetail(workflowId, detailId, engineType)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) workflowContract(gin *gin.Context) {
	id := gin.Param("id")
	workflowIdStr := gin.Param("workflowId")
	workflowDetailIdStr := gin.Param("detailId")
	workflowId, err := strconv.Atoi(workflowIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	workflowDetailId, err := strconv.Atoi(workflowDetailIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	contractService := application.GetBean[*service.ContractService]("contractService")
	data, err := contractService.QueryContractByWorkflow(id, workflowId, workflowDetailId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) workflowReport(gin *gin.Context) {
	idStr := gin.Param("id")
	workflowDetailIdStr := gin.Param("detailId")
	workflowId, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	workflowDetailId, err := strconv.Atoi(workflowDetailIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	reportService := application.GetBean[*service.ReportService]("reportService")
	data, err := reportService.QueryReportsByWorkflow(workflowId, workflowDetailId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) workflowReportOverview(gin *gin.Context) {
	idStr := gin.Param("id")
	workflowDetailIdStr := gin.Param("detailId")
	workflowId, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	workflowDetailId, err := strconv.Atoi(workflowDetailIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	reportService := application.GetBean[*service.ReportService]("reportService")
	data, err := reportService.ReportOverview(workflowId, workflowDetailId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) reportDetail(gin *gin.Context) {
	idStr := gin.Param("id")
	reportId, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	reportService := application.GetBean[*service.ReportService]("reportService")
	data, err := reportService.ReportDetail(reportId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) metaScanFile(gin *gin.Context) {
	key := gin.Param("key")
	if key == "" {
		Fail("key is empty", gin)
		return
	}
	reportService := application.GetBean[*service.ReportService]("reportService")
	data, err := reportService.GetFile(key)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) contractFileContent(gin *gin.Context) {
	idStr := gin.Param("id")
	if idStr == "" {
		Fail("project id is empty", gin)
		return
	}
	name := gin.Param("name")
	if name == "" {
		Fail("file name is empty", gin)
		return
	}
	userAny, _ := gin.Get("user")
	user, _ := userAny.(db2.User)
	data, err := h.projectService.GetProject(idStr, "")
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	tokenAny, _ := gin.Get("token")
	token, _ := tokenAny.(string)
	githubService := application.GetBean[*service.GithubService]("githubService")
	path := fmt.Sprintf("contracts%s", name)
	if strings.Contains(name, "contracts") {
		path = fmt.Sprintf("%s", name)
	}
	content, err := githubService.GetFileContent(token, user.Username, data.Name, path)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(content, gin)
}

func (h *HandlerServer) workflowFrontendReports(gin *gin.Context) {
	idStr := gin.Param("id")
	workflowDetailIdStr := gin.Param("detailId")
	workflowId, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	workflowDetailId, err := strconv.Atoi(workflowDetailIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	reportService := application.GetBean[*service.ReportService]("reportService")
	data, err := reportService.QueryFrontendReportsByWorkflow(workflowId, workflowDetailId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) workflowFrontendPackage(gin *gin.Context) {
	idStr := gin.Param("id")
	workflowDetailIdStr := gin.Param("detailId")
	workflowId, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	workflowDetailId, err := strconv.Atoi(workflowDetailIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	frontendPackageService := application.GetBean[*service.FrontendPackageService]("frontendPackageService")
	data, err := frontendPackageService.QueryPackages(workflowId, workflowDetailId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) workflowFrontendPackageDetail(gin *gin.Context) {
	packageIdStr := gin.Param("id")
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	frontendPackageService := application.GetBean[*service.FrontendPackageService]("frontendPackageService")
	data, err := frontendPackageService.QueryFrontendDeployById(packageId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) workflowFrontendDeployInfo(gin *gin.Context) {
	idStr := gin.Param("id")
	workflowDetailIdStr := gin.Param("detailId")
	workflowId, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	workflowDetailId, err := strconv.Atoi(workflowDetailIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	frontendPackageService := application.GetBean[*service.FrontendPackageService]("frontendPackageService")
	data, err := frontendPackageService.QueryFrontendDeployInfo(workflowId, workflowDetailId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) stopWorkflow(gin *gin.Context) {
	projectIdStr := gin.Param("id")
	projectId, err := uuid.FromString(projectIdStr)
	if err != nil {
		Fail("projectId is empty or invalid", gin)
		return
	}
	workflowIdStr := gin.Param("workflowId")
	detailIdStr := gin.Param("detailId")
	workflowId, err := strconv.Atoi(workflowIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	detailId, err := strconv.Atoi(detailIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	workflowKey := workflowService.GetWorkflowKey(projectId.String(), uint(workflowId))
	detail, err := workflowService.QueryWorkflowDetail(workflowId, detailId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	engineService := application.GetBean[engine.Engine]("engine")
	err = engineService.TerminalJob(workflowKey, int(detail.ExecNumber))
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

func (h *HandlerServer) deleteWorkflow(gin *gin.Context) {
	workflowIdStr := gin.Param("workflowId")
	detailIdStr := gin.Param("detailId")
	workflowId, err := strconv.Atoi(workflowIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	detailId, err := strconv.Atoi(detailIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	engineType := gin.DefaultQuery("engine", consts.EngineTypeWorkflow)
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	err = workflowService.DeleteWorkflow(workflowId, detailId, engineType)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

func (h *HandlerServer) queryReportCheckTools(gin *gin.Context) {
	id := gin.Param("id")
	reportService := application.GetBean[*service.ReportService]("reportService")
	data, err := reportService.QueryReportCheckTools(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)

}

func (h *HandlerServer) deleteWorkflowDeploy(gin *gin.Context) {
	packageIdStr := gin.Param("id")
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	frontendPackageService := application.GetBean[*service.FrontendPackageService]("frontendPackageService")
	//sh := shell.NewShell(consts.IpfsUrl)
	//err = sh.Unpin(deployData.DeployInfo)
	//if err != nil {
	//	Fail(err.Error(), gin)
	//	return
	//}
	data, err := frontendPackageService.QueryPackageById(packageId)
	if err == nil {
		data.Domain = ""
		_ = frontendPackageService.UpdateFrontedPackage(data)
	}
	_ = frontendPackageService.DeleteFrontendDeploy(packageId)
	Success("", gin)
}
