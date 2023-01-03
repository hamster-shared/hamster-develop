package controller

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/a-line/pkg/application"
	"github.com/hamster-shared/a-line/pkg/consts"
	"github.com/hamster-shared/a-line/pkg/parameter"
	"github.com/hamster-shared/a-line/pkg/service"
	"github.com/hamster-shared/a-line/pkg/utils"
	"github.com/hamster-shared/a-line/pkg/vo"
	"github.com/jinzhu/copier"
	"net/http"
	"strconv"
)

//go:embed templates
var temp embed.FS

func (h *HandlerServer) projectList(gin *gin.Context) {
	query := gin.Query("query")
	pageStr := gin.DefaultQuery("page", "1")
	sizeStr := gin.DefaultQuery("size", "10")
	userStr := gin.Query("user")
	userId, err := strconv.Atoi(userStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
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
	data, err := h.projectService.GetProjects(userId, query, page, size)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) createProject(g *gin.Context) {
	createData := parameter.CreateProjectParam{}
	err := g.BindJSON(&createData)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	accessToken := g.Request.Header.Get("access_token")
	token := utils.AesDecrypt(accessToken, consts.SecretKey)
	githubService := application.GetBean[*service.GithubService]("githubService")

	repo, res, err := githubService.CreateRepo(token, createData.TemplateOwner, createData.TemplateRepo, createData.Name, createData.RepoOwner)
	if err != nil {
		if res.StatusCode == http.StatusUnauthorized {
			Failed(http.StatusUnauthorized, "access not authorized", g)
			return
		}
		Fail(err.Error(), g)
		return
	}
	data := vo.CreateProjectParam{
		Name:        createData.Name,
		Type:        createData.Type,
		TemplateUrl: *repo.URL,
		FrameType:   createData.FrameType,
		UserId:      createData.UserId,
	}
	id, err := h.projectService.CreateProject(data)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	file, err := workflowService.TemplateParse("solidity-check", *repo.URL, consts.Check)
	if err == nil {
		workflowData := parameter.SaveWorkflowParam{
			ProjectId:  id,
			Type:       consts.Check,
			ExecFile:   file,
			LastExecId: 0,
		}
		workflowService.SaveWorkflow(workflowData)
	}
	file1, err := workflowService.TemplateParse("solidity-build", *repo.URL, consts.Build)
	if err == nil {
		workflowData := parameter.SaveWorkflowParam{
			ProjectId:  id,
			Type:       consts.Build,
			ExecFile:   file1,
			LastExecId: 0,
		}
		workflowService.SaveWorkflow(workflowData)
	}
	Success(id, g)

	// TODO ... 检查用户仓库是否被使用

	//TODO ... github 导入仓库

	//TODO ... 保存项目信息

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
	//accessToken := g.Request.Header.Get("access_token")
	//if accessToken == "" {
	//	Failed(http.StatusUnauthorized, "No access",g)
	//	return
	//}
	//token := utils.AesDecrypt(accessToken, consts.SecretKey)
	userService := application.GetBean[*service.UserService]("userService")
	var userVo vo.UserAuth
	user, err := userService.GetUserByToken("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")
	if err != nil {
		Fail("get user info failed", g)
		return
	}
	copier.Copy(&userVo, &user)
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	_ = workflowService.ExecProjectCheckWorkflow(uint(projectId), userVo)
	Success("", g)
}

func (h *HandlerServer) projectWorkflowBuild(g *gin.Context) {

	projectId, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		Fail("projectId is empty or invalid", g)
		return
	}
	//accessToken := g.Request.Header.Get("access_token")
	//if accessToken == "" {
	//	Failed(http.StatusUnauthorized, "No access",g)
	//	return
	//}
	//token := utils.AesDecrypt(accessToken, consts.SecretKey)
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	userService := application.GetBean[*service.UserService]("userService")
	var userVo vo.UserAuth
	user, err := userService.GetUserByToken("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")
	if err != nil {
		Fail("get user info failed", g)
		return
	}
	copier.Copy(&userVo, &user)
	err = workflowService.ExecProjectBuildWorkflow(uint(projectId), userVo)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	Success("", g)
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

	result, err := contractService.QueryContracts(uint(projectId), query, version, "", page, size)

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

func (h *HandlerServer) checkName(gin *gin.Context) {
	var checkData parameter.CheckNameParam
	err := gin.BindJSON(&checkData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	accessToken := gin.Request.Header.Get("access_token")
	token := utils.AesDecrypt(accessToken, consts.SecretKey)
	githubService := application.GetBean[*service.GithubService]("githubService")
	data := githubService.CheckName(token, checkData.Owner, checkData.Name)
	Success(data, gin)
}
