package controller

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/a-line/pkg/application"
	"github.com/hamster-shared/a-line/pkg/controller/parameters"
	"github.com/hamster-shared/a-line/pkg/db"
	"github.com/hamster-shared/a-line/pkg/service"
	"github.com/hamster-shared/a-line/pkg/vo"
	"github.com/young2j/gocopy"
	"gorm.io/gorm"
	"strconv"
)

//go:embed templates
var temp embed.FS

func (h *HandlerServer) projectList(gin *gin.Context) {
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
	data, err := h.projectService.GetProjects(query, page, size)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) createProject(g *gin.Context) {

	// TODO ... 检查用户仓库是否被使用

	//TODO ... github 导入仓库

	//TODO ... 保存项目信息
	var createVo parameters.ProjectCreateVo
	err := g.BindJSON(&createVo)
	if err != nil {
		Fail("param error", g)
		return
	}

	database := application.GetBean[*gorm.DB]("db")
	var project db.Project
	gocopy.Copy(&project, &createVo)
	project.RepositoryUrl = createVo.TemplateUrl

	project.UserId = h.getUserInfo(g).Id

	database.Create(project)

	//TODO ... 保存默认流水线
	//fs, err := temp.Open("truffle_check.yml")
	//template.New("check").
	//workflow1 := db.Workflow{
	//	ProjectId:  project.Id,
	//	Type:       1,
	//	ExecFile:   "",
	//	LastExecId: 0,
	//}
}

func (h *HandlerServer) projectDetail(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	data, err := h.projectService.GetProject(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) projectWorkflowCheck(g *gin.Context) {

	projectId, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Fail("projectId is empty or invalid", g)
		return
	}

	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	_ = workflowService.ExecProjectCheckWorkflow(uint(projectId), h.getUserInfo(g))
	Success("", g)
}

func (h *HandlerServer) projectWorkflowBuild(g *gin.Context) {

	Success("", g)
}

func (h *HandlerServer) projectWorkflowDeploy(g *gin.Context) {

}

func (h *HandlerServer) projectContract(g *gin.Context) {

	projectId, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Fail("projectId is empty or invalid", g)
		return
	}

	query := g.Query("query")
	version := g.Query("version")
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(g.DefaultQuery("size", "10"))

	contractService := application.GetBean[*service.ContractService]("contractService")

	result, err := contractService.QueryContracts(uint(projectId), query, version, page, size)

	if err != nil {
		Fail(err.Error(), g)
		return
	}

	Success(result, g)

}

func (h *HandlerServer) projectReport(g *gin.Context) {
	projectId, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Fail("projectId is empty or invalid", g)
		return
	}

	Type, _ := strconv.Atoi(g.DefaultQuery("type", "1"))
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(g.DefaultQuery("size", "10"))

	reportService := application.GetBean[*service.ReportService]("reportService")

	result, err := reportService.QueryReports(uint(projectId), uint(Type), page, size)

	if err != nil {
		Fail(err.Error(), g)
		return
	}

	Success(result, g)

}

func (h *HandlerServer) updateProject(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	var updateData vo.UpdateProjectParam
	err = gin.BindJSON(&updateData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	err = h.projectService.UpdateProject(id, updateData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

func (h *HandlerServer) deleteProject(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	err = h.projectService.DeleteProject(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}
