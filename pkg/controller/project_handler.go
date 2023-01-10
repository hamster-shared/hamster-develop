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
	uuid "github.com/iris-contrib/go.uuid"
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
	accessToken := gin.Request.Header.Get("Access-Token")
	if accessToken == "" {
		Failed(http.StatusUnauthorized, "No access", gin)
		return
	}
	token := utils.AesDecrypt(accessToken, consts.SecretKey)
	userService := application.GetBean[*service.UserService]("userService")
	user, err := userService.GetUserByToken(token)
	data, err := h.projectService.GetProjects(int(user.Id), query, page, size)
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
	accessToken := g.Request.Header.Get("Access-Token")
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
	userService := application.GetBean[*service.UserService]("userService")
	user, err := userService.GetUserByToken(token)
	if err != nil {
		Fail("get user info failed", g)
		return
	}
	data := vo.CreateProjectParam{
		Name:        createData.Name,
		Type:        createData.Type,
		TemplateUrl: *repo.CloneURL,
		FrameType:   createData.FrameType,
		UserId:      int64(user.Id),
	}
	id, err := h.projectService.CreateProject(data)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	workflowCheckData := parameter.SaveWorkflowParam{
		ProjectId:  id,
		Type:       consts.Check,
		ExecFile:   "",
		LastExecId: 0,
	}
	workflowCheckRes, err := workflowService.SaveWorkflow(workflowCheckData)
	if err != nil {
		Success(id, g)
		return
	}
	checkKey := workflowService.GetWorkflowKey(id.String(), workflowCheckRes.Id)
	file, err := workflowService.TemplateParse(checkKey, *repo.CloneURL, consts.Check)
	if err == nil {
		workflowCheckRes.ExecFile = file
		workflowService.UpdateWorkflow(workflowCheckRes)
	}
	workflowBuildData := parameter.SaveWorkflowParam{
		ProjectId:  id,
		Type:       consts.Build,
		ExecFile:   "",
		LastExecId: 0,
	}
	workflowBuildRes, err := workflowService.SaveWorkflow(workflowBuildData)
	if err != nil {
		Success(id, g)
		return
	}
	buildKey := workflowService.GetWorkflowKey(id.String(), workflowBuildRes.Id)
	file1, err := workflowService.TemplateParse(buildKey, *repo.CloneURL, consts.Build)
	if err == nil {
		workflowBuildRes.ExecFile = file1
		workflowService.UpdateWorkflow(workflowBuildRes)
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
	id := gin.Param("id")
	data, err := h.projectService.GetProject(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) projectWorkflowCheck(g *gin.Context) {

	projectIdStr := g.Param("id")
	projectId, err := uuid.FromString(projectIdStr)
	if err != nil {
		Fail("projectId is empty or invalid", g)
		return
	}
	accessToken := g.Request.Header.Get("Access-Token")
	if accessToken == "" {
		Failed(http.StatusUnauthorized, "No access", g)
		return
	}
	token := utils.AesDecrypt(accessToken, consts.SecretKey)
	userService := application.GetBean[*service.UserService]("userService")
	var userVo vo.UserAuth
	user, err := userService.GetUserByToken(token)
	if err != nil {
		Fail("get user info failed", g)
		return
	}
	copier.Copy(&userVo, &user)
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	_ = workflowService.ExecProjectCheckWorkflow(projectId, userVo)
	Success("", g)
}

func (h *HandlerServer) projectWorkflowBuild(g *gin.Context) {

	projectIdStr := g.Param("id")
	projectId, err := uuid.FromString(projectIdStr)
	if err != nil {
		Fail("projectId is empty or invalid", g)
		return
	}
	accessToken := g.Request.Header.Get("Access-Token")
	if accessToken == "" {
		Failed(http.StatusUnauthorized, "No access", g)
		return
	}
	token := utils.AesDecrypt(accessToken, consts.SecretKey)
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	userService := application.GetBean[*service.UserService]("userService")
	var userVo vo.UserAuth
	user, err := userService.GetUserByToken(token)
	if err != nil {
		Fail("get user info failed", g)
		return
	}
	copier.Copy(&userVo, &user)
	err = workflowService.ExecProjectBuildWorkflow(projectId, userVo)
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
	network := g.Query("network")
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(g.DefaultQuery("size", "10"))

	contractService := application.GetBean[*service.ContractService]("contractService")

	result, err := contractService.QueryContracts(uint(projectId), query, version, network, page, size)

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

	Type := g.DefaultQuery("type", "")
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(g.DefaultQuery("size", "10"))

	reportService := application.GetBean[*service.ReportService]("reportService")

	result, err := reportService.QueryReports(uint(projectId), Type, page, size)

	if err != nil {
		Fail(err.Error(), g)
		return
	}

	Success(result, g)

}

func (h *HandlerServer) updateProject(gin *gin.Context) {
	id := gin.Param("id")
	var updateData vo.UpdateProjectParam
	err := gin.BindJSON(&updateData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	accessToken := gin.Request.Header.Get("Access-Token")
	if accessToken == "" {
		Failed(http.StatusUnauthorized, "No access", gin)
		return
	}
	token := utils.AesDecrypt(accessToken, consts.SecretKey)
	userService := application.GetBean[*service.UserService]("userService")
	user, err := userService.GetUserByToken(token)
	updateData.UserId = int(user.Id)
	project, err := h.projectService.GetProject(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	githubService := application.GetBean[*service.GithubService]("githubService")
	repo, res, err := githubService.UpdateRepo(token, user.Username, project.Name, updateData.Name)
	if err != nil {
		if res.StatusCode == http.StatusUnauthorized {
			Failed(http.StatusUnauthorized, "access not authorized", gin)
			return
		}
		Fail(err.Error(), gin)
		return
	}
	updateData.RepositoryUrl = *repo.CloneURL
	err = h.projectService.UpdateProject(id, updateData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

func (h *HandlerServer) deleteProject(gin *gin.Context) {
	id := gin.Param("id")
	err := h.projectService.DeleteProject(id)
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
	accessToken := gin.Request.Header.Get("Access-Token")
	token := utils.AesDecrypt(accessToken, consts.SecretKey)
	githubService := application.GetBean[*service.GithubService]("githubService")
	data := githubService.CheckName(token, checkData.Owner, checkData.Name)
	Success(data, gin)
}
