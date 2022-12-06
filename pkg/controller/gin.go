package controller

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
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

	{
		project := api.Group("/projects")
		{
			//create project
			project.POST("", h.handlerServer.createProject)
			//update project
			project.PUT("/:oldName", h.handlerServer.updateProject)
			//get project list
			project.GET("/list", h.handlerServer.getProjects)
			//delete project
			project.DELETE("/:name", h.handlerServer.deleteProject)
		}

		projectGroup := api.Group("/project/:projectName")
		{
			//create pipeline job
			projectGroup.POST("/pipeline", h.handlerServer.createPipeline)
			//update pipeline job
			projectGroup.PUT("/pipeline/:oldName", h.handlerServer.updatePipeline)
			//get pipeline job
			projectGroup.GET("/pipeline/:name", h.handlerServer.getPipeline)
			//delete pipeline job and pipeline job detail
			projectGroup.DELETE("/pipeline/:name", h.handlerServer.deletePipeline)
			//get pipeline job list
			projectGroup.GET("/pipeline", h.handlerServer.pipelineList)
			//get pipeline job detail info
			projectGroup.GET("/pipeline/:name/detail/:id", h.handlerServer.getPipelineDetail)
			//delete pipeline job detail
			projectGroup.DELETE("/pipeline/:name/detail/:id", h.handlerServer.deleteJobDetail)
			//get pipeline job detail list
			projectGroup.GET("/pipeline/:name/details", h.handlerServer.getPipelineDetailList)
			//exec pipeline job
			projectGroup.POST("/pipeline/:name/exec", h.handlerServer.execPipeline)
			//re exec pipeline detail job
			projectGroup.POST("/pipeline/:name/:id/re-exec", h.handlerServer.reExecuteJob)
			//stop pipeline job
			projectGroup.POST("/pipeline/:name/:id/stop", h.handlerServer.stopJobDetail)
			projectGroup.GET("/pipeline/:name/logs/:id", h.handlerServer.getJobLog)
			projectGroup.GET("/pipeline/:name/logs/:id/:stagename", h.handlerServer.getJobStageLog)
			projectGroup.GET("/pipeline/:name/detail/:id/artifactory", h.handlerServer.openArtifactoryDir)
		}

		// get template list
		api.GET("/pipeline/templates", h.handlerServer.getTemplates)
		api.GET("/pipeline/template/:id", h.handlerServer.getTemplateDetail)
		api.GET("/ping", func(c *gin.Context) {
			//输出json结果给调用方
			Success("", c)
		})

	}

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
