package controller

import (
	"github.com/gin-gonic/gin"
	engine "github.com/hamster-shared/aline-engine"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/service"
	uuid "github.com/iris-contrib/go.uuid"
	"strconv"
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
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	data, err := workflowService.GetWorkflowDetail(workflowId, detailId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) workflowContract(gin *gin.Context) {
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
	contractService := application.GetBean[*service.ContractService]("contractService")
	data, err := contractService.QueryContractByWorkflow(workflowId, workflowDetailId)
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
	engineService := application.GetBean[*engine.Engine]("engine")
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
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	err = workflowService.DeleteWorkflow(workflowId, detailId)
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
	frontendPackageService := application.GetBean[*service.FrontendPackageService]("frontendPackageService")
	deployData, err := frontendPackageService.QueryFrontendDeployInfo(workflowId, detailId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	//sh := shell.NewShell(consts.IpfsUrl)
	//err = sh.Unpin(deployData.DeployInfo)
	//if err != nil {
	//	Fail(err.Error(), gin)
	//	return
	//}
	data, err := frontendPackageService.QueryPackageById(int(deployData.PackageId))
	if err == nil {
		data.Domain = ""
		frontendPackageService.UpdateFrontedPackage(data)
	}
	frontendPackageService.DeleteFrontendDeploy(workflowId, detailId)
	Success("", gin)
}
