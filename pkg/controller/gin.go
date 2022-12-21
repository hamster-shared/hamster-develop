package controller

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/a-line/pkg/application"
	"gorm.io/gorm"
	"io/fs"
	"net/http"
	"os/exec"
	"runtime"
)

//go:embed dist
var content embed.FS

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

	// project_template
	api.GET("/templates-category", h.handlerServer.templatesCategory)
	api.GET("/templates-category/:id/templates", h.handlerServer.templates)
	api.GET("/templates-category/:id/template/:templateId", h.handlerServer.templateDetail)
	// project
	api.GET("/projects", h.handlerServer.projectList)
	api.POST("/projects", h.handlerServer.createProject)
	api.GET("/projects/:id", h.handlerServer.projectDetail)
	api.PUT("/projects/:id", h.handlerServer.updateProject)
	api.DELETE("projects/:id", h.handlerServer.deleteProject)

	api.POST("/projects/:id/check", h.handlerServer.projectWorkflowCheck)
	api.POST("/projects/:id/build", h.handlerServer.projectWorkflowBuild)
	api.POST("/projects/:id/deploy", h.handlerServer.projectWorkflowDeploy)
	api.GET("/projects/:id/contract", h.handlerServer.projectContract)
	api.GET("/projects/:id/reports", h.handlerServer.projectReport)
	api.POST("/projects/:id/contract/deploy", h.handlerServer.saveContractDeployInfo)

	//workflow
	api.GET("/projects/:id/workflows", h.handlerServer.workflowList)
	api.GET("/workflows/:id", h.handlerServer.workflowDetail)
	api.GET("/workflows/:id/contract", h.handlerServer.workflowContract)
	api.GET("/workflows/:id/report", h.handlerServer.workflowReport)

	//contract
	api.GET("/contract/:deployId", h.handlerServer.contractDeployInfo)
	api.GET("/contract/deploy/:deployId/detail", h.handlerServer.contractDeployDetail)
	api.GET("/contract/deploy/detail", h.handlerServer.contractDeployDetailByVersion)

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
	api.GET("/pipeline/templates", h.handlerServer.getTemplates)
	api.GET("/pipeline/template/:id", h.handlerServer.getTemplateDetail)
	api.GET("/pipeline/:name/detail/:id/artifactory", h.handlerServer.openArtifactoryDir)
	api.GET("/ping", func(c *gin.Context) {

		db := application.GetBean[*gorm.DB]("db")
		fmt.Println(db)
		//输出json结果给调用方
		Success("", c)
	})
	fe, _ := fs.Sub(content, "dist")
	r.NoRoute(gin.WrapH(http.FileServer(http.FS(fe))))
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
