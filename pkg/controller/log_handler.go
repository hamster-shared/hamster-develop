package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/a-line/engine"
	"github.com/hamster-shared/a-line/engine/model"
	"github.com/hamster-shared/a-line/pkg/application"
	"github.com/hamster-shared/a-line/pkg/service"
	"strconv"
)

func (h *HandlerServer) getWorkflowLog(gin *gin.Context) {
	idStr := gin.Param("id")
	detailIdStr := gin.Param("detailId")
	workflowId, err := strconv.Atoi(idStr)
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
	workflow, err := workflowService.QueryWorkflow(workflowId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	workflowDetail, err := workflowService.QueryWorkflowDetail(workflowId, detailId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	logName := workflowService.GetWorkflowKey(workflow.ProjectId.String(), workflow.Id)
	engine := application.GetBean[*engine.Engine]("engine")
	jobDetail := engine.GetJobHistory(logName, int(workflowDetail.ExecNumber))
	data := engine.GetJobHistoryLog(logName, int(workflowDetail.ExecNumber))

	gin.Writer.Header().Set("LastLine", strconv.Itoa(data.LastLine))
	gin.Writer.Header().Set("End", strconv.FormatBool(jobDetail.Status != model.STATUS_RUNNING))
	Success(data, gin)

}

func (h *HandlerServer) getWorkflowStageLog(gin *gin.Context) {
	idStr := gin.Param("id")
	detailIdStr := gin.Param("detailId")
	name := gin.Param("stageName")
	startStr := gin.DefaultQuery("start", "0")
	workflowId, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	detailId, err := strconv.Atoi(detailIdStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	start, err := strconv.Atoi(startStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	engine := application.GetBean[*engine.Engine]("engine")
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	workflow, err := workflowService.QueryWorkflow(workflowId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	workflowDetail, err := workflowService.QueryWorkflowDetail(workflowId, detailId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	logName := workflowService.GetWorkflowKey(workflow.ProjectId.String(), workflow.Id)
	data := engine.GetJobHistoryStageLog(logName, int(workflowDetail.ExecNumber), name, start)
	gin.Writer.Header().Set("LastLine", strconv.Itoa(data.LastLine))
	gin.Writer.Header().Set("End", strconv.FormatBool(data.End))
	Success(data, gin)
}