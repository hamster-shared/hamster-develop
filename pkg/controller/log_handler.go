package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	engine "github.com/hamster-shared/aline-engine"
	"github.com/hamster-shared/aline-engine/model"
	"github.com/hamster-shared/aline-engine/utils"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/service"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 或者编写一个函数过滤好多请求源。
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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

func (h *HandlerServer) getDeployFrontendLog(gin *gin.Context) {
	projectIdStr := gin.Param("id")
	if projectIdStr == "" {
		Fail("projectId is empty or invalid", gin)
		return
	}
	username := gin.Param("username")
	if username == "" {
		Fail("username is empty or invalid", gin)
		return
	}
	//userAny, _ := gin.Get("user")
	//user, _ := userAny.(db2.User)

	project, err := h.projectService.GetProject(projectIdStr)
	if err != nil {
		log.Println("get project failed", err.Error())
		Fail(err.Error(), gin)
		return
	}
	conn, err := upgrader.Upgrade(gin.Writer, gin.Request, nil)
	if err != nil {
		log.Println("websocket connection failed", err.Error())
		Fail(err.Error(), gin)
		return
	}
	name := fmt.Sprintf("%s-%s", username, project.Name)
	req, err := utils.GetPodLogs(name, name, username)
	if err != nil {
		log.Println("get pod logs failed", err.Error())
		Fail(err.Error(), gin)
		return
	}
	stream, err := req.Stream(context.TODO())
	if err != nil {
		log.Println("get pod logs failed", err.Error())
		Fail(err.Error(), gin)
		return
	}
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		numBytes, err := stream.Read(buf)
		if err != nil {
			log.Printf("Error reading from log stream for pod: %s\n", err.Error())
			Fail(err.Error(), gin)
			return
		}
		msgBytes := buf[:numBytes]
		err = conn.WriteMessage(websocket.TextMessage, msgBytes)
		if err != nil {
			Fail(err.Error(), gin)
			log.Println(err)
			return
		}
	}
}
