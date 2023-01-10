package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/a-line/engine"
	"github.com/hamster-shared/a-line/engine/consts"
	"github.com/hamster-shared/a-line/engine/model"
	"github.com/hamster-shared/a-line/engine/utils"
	"github.com/hamster-shared/a-line/engine/utils/platform"
	"github.com/hamster-shared/a-line/pkg/controller/parameters"
	service2 "github.com/hamster-shared/a-line/pkg/service"
	"github.com/hamster-shared/a-line/pkg/vo"
	"gopkg.in/yaml.v3"
	"path/filepath"
	"strconv"
)

type HandlerServer struct {
	Engine          *engine.Engine
	templateService service2.ITemplateService
	projectService  service2.IProjectService
}

func NewHandlerServer(engine *engine.Engine, templateService service2.ITemplateService, projectService service2.IProjectService) *HandlerServer {
	return &HandlerServer{
		Engine:          engine,
		templateService: templateService,
		projectService:  projectService,
	}
}

// createPipeline create pipeline jon
func (h *HandlerServer) createPipeline(gin *gin.Context) {
	createData := parameters.CreatePipeline{}
	err := gin.BindJSON(&createData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	err = h.Engine.CreateJob(createData.Name, createData.Yaml)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}

	Success("", gin)
}

func (h *HandlerServer) updatePipeline(gin *gin.Context) {
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
	err = h.Engine.UpdateJob(oldName, updateData.NewName, updateData.Yaml)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

// getPipeline get pipeline job
func (h *HandlerServer) getPipeline(gin *gin.Context) {
	name := gin.Param("name")
	pipelineData := h.Engine.GetJob(name)
	Success(pipelineData, gin)
}

// deletePipeline delete pipeline job and pipeline job detail
func (h *HandlerServer) deletePipeline(gin *gin.Context) {
	name := gin.Param("name")
	err := h.Engine.DeleteJob(name)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

// pipelineList get pipeline job list
func (h *HandlerServer) pipelineList(gin *gin.Context) {
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
	jobData := h.Engine.GetJobs(query, page, size)
	Success(jobData, gin)
}

// getPipelineDetail get pipeline job detail info
func (h *HandlerServer) getPipelineDetail(gin *gin.Context) {
	name := gin.Param("name")
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	jobDetailData := h.Engine.GetJobHistory(name, id)
	Success(jobDetailData, gin)
}

// deleteJobDetail delete job detail
func (h *HandlerServer) deleteJobDetail(gin *gin.Context) {
	name := gin.Param("name")
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	err = h.Engine.DeleteJobHistory(name, id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

// getPipelineDetailList get pipeline job detail list
func (h *HandlerServer) getPipelineDetailList(gin *gin.Context) {
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
	jobDetailPage := h.Engine.GetJobHistorys(name, page, size)
	Success(jobDetailPage, gin)
}

// execPipeline exec pipeline job
func (h *HandlerServer) execPipeline(gin *gin.Context) {
	name := gin.Param("name")
	_, err := h.Engine.ExecuteJob(name)

	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

// reExecuteJob re exec pipeline job detail
func (h *HandlerServer) reExecuteJob(gin *gin.Context) {
	name := gin.Param("name")
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}

	err = h.Engine.ReExecuteJob(name, id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

// stopJobDetail stop pipeline job
func (h *HandlerServer) stopJobDetail(gin *gin.Context) {
	name := gin.Param("name")
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}

	err = h.Engine.TerminalJob(name, id)
	Success("", gin)
}

// getJobLog get pipeline job detail logs
func (h *HandlerServer) getJobLog(gin *gin.Context) {
	name := gin.Param("name")
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	jobDetail := h.Engine.GetJobHistory(name, id)
	data := h.Engine.GetJobHistoryLog(name, id)

	gin.Writer.Header().Set("LastLine", strconv.Itoa(data.LastLine))
	gin.Writer.Header().Set("End", strconv.FormatBool(jobDetail.Status != model.STATUS_RUNNING))
	//gin.String(200, data.Content)
	Success(data, gin)
}

// getJobStageLog get job detail stage logs
func (h *HandlerServer) getJobStageLog(gin *gin.Context) {
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
	data := h.Engine.GetJobHistoryStageLog(name, id, stageName, start)

	gin.Writer.Header().Set("LastLine", strconv.Itoa(data.LastLine))
	gin.Writer.Header().Set("End", strconv.FormatBool(data.End))
	//gin.String(200, data.Content)
	Success(data, gin)
}

//// getTemplates get template list
//func (h *HandlerServer) getTemplates(gin *gin.Context) {
//	lang := gin.Request.Header.Get("lang")
//	if lang == "" {
//		lang = consts.LANG_EN
//	}
//	data := h.templateService1.GetTemplates(lang)
//	Success(data, gin)
//}
//
//// getTemplateDetail get template detail
//func (h *HandlerServer) getTemplateDetail(gin *gin.Context) {
//	idStr := gin.Param("id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		Fail(err.Error(), gin)
//		return
//	}
//	data, _ := h.templateService1.GetTemplateDetail(id)
//	Success(data, gin)
//}

// openArtifactoryDir open artifactory folder
func (h *HandlerServer) openArtifactoryDir(gin *gin.Context) {
	idStr := gin.Param("id")
	name := gin.Param("name")
	artifactoryDir := filepath.Join(utils.DefaultConfigDir(), consts.JOB_DIR_NAME, name, consts.ArtifactoryDir, idStr)
	err := platform.OpenDir(artifactoryDir)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}

	Success("", gin)
}

func (h *HandlerServer) getUserInfo(gin *gin.Context) vo.UserAuth {

	// token 是什么东西?，方案1：我们自己的jwt token, 方案2: github token
	token := gin.GetHeader("access_token")

	//TODO...
	//token = db_replace(token)

	// TODO ... 根据token 获取用户信息
	return vo.UserAuth{
		Id:       1,
		Username: "admin",
		Token:    token,
	}
}
