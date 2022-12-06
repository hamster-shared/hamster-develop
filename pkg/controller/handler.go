package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/a-line/pkg/consts"
	"github.com/hamster-shared/a-line/pkg/controller/parameters"
	"github.com/hamster-shared/a-line/pkg/dispatcher"
	"github.com/hamster-shared/a-line/pkg/model"
	"github.com/hamster-shared/a-line/pkg/service"
	"gopkg.in/yaml.v3"
	"log"
	"strconv"
)

type HandlerServer struct {
	jobService      service.IJobService
	dispatch        dispatcher.IDispatcher
	templateService service.ITemplateService
	projectService  service.IProjectService
}

func NewHandlerServer(jobService service.IJobService, dispatch dispatcher.IDispatcher, templateService service.ITemplateService, projectService service.IProjectService) *HandlerServer {
	return &HandlerServer{
		jobService:      jobService,
		dispatch:        dispatch,
		templateService: templateService,
		projectService:  projectService,
	}
}

// createPipeline create pipeline jon
func (h *HandlerServer) createPipeline(gin *gin.Context) {
	projectName := gin.Param("projectName")
	createData := parameters.CreatePipeline{}
	err := gin.BindJSON(&createData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	var jobData model.Job
	err = yaml.Unmarshal([]byte(createData.Yaml), &jobData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	err = h.jobService.SaveJob(projectName, createData.Name, createData.Yaml)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}

	Success("", gin)
}

func (h *HandlerServer) updatePipeline(gin *gin.Context) {
	projectName := gin.Param("projectName")
	oldName := gin.Param("oldName")
	updateData := parameters.UpdatePipeline{}
	err := gin.BindJSON(&updateData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	var jobData model.Job
	err = yaml.Unmarshal([]byte(updateData.Yaml), &jobData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	err = h.jobService.UpdateJob(projectName, oldName, updateData.NewName, updateData.Yaml)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

// getPipeline get pipeline job
func (h *HandlerServer) getPipeline(gin *gin.Context) {
	projectName := gin.Param("projectName")
	name := gin.Param("name")
	pipelineData, err := h.jobService.GetJob(projectName, name)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(pipelineData, gin)
}

// deletePipeline delete pipeline job and pipeline job detail
func (h *HandlerServer) deletePipeline(gin *gin.Context) {
	projectName := gin.Param("projectName")
	name := gin.Param("name")
	err := h.jobService.DeleteJob(projectName, name)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

// pipelineList get pipeline job list
func (h *HandlerServer) pipelineList(gin *gin.Context) {
	projectName := gin.Param("projectName")
	query := gin.Query("query")
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
	jobData := h.jobService.JobList(projectName, query, page, size)
	Success(jobData, gin)
}

// getPipelineDetail get pipeline job detail info
func (h *HandlerServer) getPipelineDetail(gin *gin.Context) {
	projectName := gin.Param("projectName")
	name := gin.Param("name")
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	jobDetailData, err := h.jobService.GetJobDetail(projectName, name, id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(jobDetailData, gin)
}

// deleteJobDetail delete job detail
func (h *HandlerServer) deleteJobDetail(gin *gin.Context) {
	projectName := gin.Param("projectName")
	name := gin.Param("name")
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	err = h.jobService.DeleteJobDetail(projectName, name, id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

// getPipelineDetailList get pipeline job detail list
func (h *HandlerServer) getPipelineDetailList(gin *gin.Context) {
	projectName := gin.Param("projectName")
	name := gin.Param("name")
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
	jobDetailPage := h.jobService.JobDetailList(projectName, name, page, size)
	Success(jobDetailPage, gin)
}

// execPipeline exec pipeline job
func (h *HandlerServer) execPipeline(gin *gin.Context) {
	projectName := gin.Param("projectName")
	name := gin.Param("name")
	job := h.jobService.GetJobObject(projectName, name)
	jobDetail, err := h.jobService.ExecuteJob(projectName, name)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	node := h.dispatch.DispatchNode(job)
	h.dispatch.SendJob(projectName, jobDetail, node)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

// reExecuteJob re exec pipeline job detail
func (h *HandlerServer) reExecuteJob(gin *gin.Context) {
	projectName := gin.Param("projectName")
	name := gin.Param("name")
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	err = h.jobService.ReExecuteJob(projectName, name, id)
	job := h.jobService.GetJobObject(projectName, name)
	jobDetail, err := h.jobService.GetJobDetail(projectName, name, id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	node := h.dispatch.DispatchNode(job)
	h.dispatch.SendJob(projectName, jobDetail, node)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

// stopJobDetail stop pipeline job
func (h *HandlerServer) stopJobDetail(gin *gin.Context) {
	projectName := gin.Param("projectName")
	name := gin.Param("name")
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}

	err = h.jobService.StopJobDetail(projectName, name, id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	job := h.jobService.GetJobObject(projectName, name)
	jobDetail, err := h.jobService.GetJobDetail(projectName, name, id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	node := h.dispatch.DispatchNode(job)
	h.dispatch.CancelJob(projectName, jobDetail, node)
	Success("", gin)
}

// getJobLog get pipeline job detail logs
func (h *HandlerServer) getJobLog(gin *gin.Context) {
	projectName := gin.Param("projectName")
	name := gin.Param("name")
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	jobDetail, err := h.jobService.GetJobDetail(projectName, name, id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	data := h.jobService.GetJobLog(projectName, name, id)
	log.Println(111111111)
	log.Println(data)
	log.Println(111111111)
	gin.Writer.Header().Set("LastLine", strconv.Itoa(data.LastLine))
	gin.Writer.Header().Set("End", strconv.FormatBool(jobDetail.Status != model.STATUS_RUNNING))
	//gin.String(200, data.Content)
	Success(data, gin)
}

// getJobStageLog get job detail stage logs
func (h *HandlerServer) getJobStageLog(gin *gin.Context) {
	projectName := gin.Param("projectName")
	name := gin.Param("name")
	idStr := gin.Param("id")
	stageName := gin.Param("stagename")
	startStr := gin.DefaultQuery("start", "0")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	start, _ := strconv.Atoi(startStr)
	data := h.jobService.GetJobStageLog(projectName, name, id, stageName, start)

	gin.Writer.Header().Set("LastLine", strconv.Itoa(data.LastLine))
	gin.Writer.Header().Set("End", strconv.FormatBool(data.End))
	//gin.String(200, data.Content)
	Success(data, gin)
}

// getTemplates get template list
func (h *HandlerServer) getTemplates(gin *gin.Context) {
	lang := gin.Request.Header.Get("lang")
	if lang == "" {
		lang = consts.LANG_EN
	}
	data := h.templateService.GetTemplates(lang)
	Success(data, gin)
}

// getTemplateDetail get template detail
func (h *HandlerServer) getTemplateDetail(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	data, _ := h.templateService.GetTemplateDetail(id)
	Success(data, gin)
}

// openArtifactoryDir open artifactory folder
func (h *HandlerServer) openArtifactoryDir(gin *gin.Context) {
	projectName := gin.Param("projectName")
	idStr := gin.Param("id")
	name := gin.Param("name")
	err := h.jobService.OpenArtifactoryDir(projectName, name, idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

// createProject create project
func (h *HandlerServer) createProject(gin *gin.Context) {
	var project model.Project
	err := gin.BindJSON(&project)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	err = h.projectService.CreateProject(&project)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

// updateProject update project
func (h *HandlerServer) updateProject(gin *gin.Context) {
	oldName := gin.Param("oldName")
	var project model.Project
	err := gin.BindJSON(&project)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	err = h.projectService.UpdateProject(oldName, &project)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

// getProjects get project list
func (h *HandlerServer) getProjects(gin *gin.Context) {
	query := gin.Query("query")
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
	data := h.projectService.GetProjects(query, page, size)
	Success(data, gin)
}

// deleteProject delete project by name
func (h *HandlerServer) deleteProject(gin *gin.Context) {
	name := gin.Param("name")
	err := h.projectService.DeleteProject(name)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}
