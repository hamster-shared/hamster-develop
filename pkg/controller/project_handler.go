package controller

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/service"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
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
	projectTypeStr := gin.DefaultQuery("type", "1")
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
	projectType, err := strconv.Atoi(projectTypeStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	userAny, _ := gin.Get("user")
	user, _ := userAny.(db2.User)
	data, err := h.projectService.GetProjects(int(user.Id), query, page, size, projectType)
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
	tokenAny, _ := g.Get("token")
	token, _ := tokenAny.(string)
	userAny, _ := g.Get("user")
	user, _ := userAny.(db2.User)
	githubService := application.GetBean[*service.GithubService]("githubService")
	repo, res, err := githubService.GetRepo(token, user.Username, createData.Name)
	if err != nil {
		if res != nil {
			if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
				Failed(http.StatusUnauthorized, "access not authorized", g)
				return
			}
		}
		repo, res, err = githubService.CreateRepository(token, createData.Name)
		if err != nil {
			if res != nil {
				if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
					Failed(http.StatusUnauthorized, "access not authorized", g)
					return
				}
			}
			Fail(err.Error(), g)
			return
		}
	}
	//email, err := githubService.GetUserEmail(token)
	//if err != nil {
	//	Fail(err.Error(), g)
	//	return
	//}
	err = githubService.CommitAndPush(token, *repo.CloneURL, user.Username, user.UserEmail, createData.TemplateUrl, createData.TemplateRepo)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	data := vo.CreateProjectParam{
		Name:        createData.Name,
		Type:        createData.Type,
		TemplateUrl: *repo.CloneURL,
		FrameType:   createData.FrameType,
		DeployType:  createData.DeployType,
		UserId:      int64(user.Id),
	}
	id, err := h.projectService.CreateProject(data)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	project, err := h.projectService.GetProject(id.String())
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
	file, err := workflowService.TemplateParse(checkKey, project, consts.Check)
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
	file1, err := workflowService.TemplateParse(buildKey, project, consts.Build)
	if err == nil {
		workflowBuildRes.ExecFile = file1
		workflowService.UpdateWorkflow(workflowBuildRes)
	}

	if project.Type == uint(consts.FRONTEND) {
		workflowDeployData := parameter.SaveWorkflowParam{
			ProjectId:  id,
			Type:       consts.Deploy,
			ExecFile:   "",
			LastExecId: 0,
		}
		workflowDeployRes, err := workflowService.SaveWorkflow(workflowDeployData)
		if err != nil {
			Success(id, g)
			return
		}
		deployKey := workflowService.GetWorkflowKey(id.String(), workflowDeployRes.Id)
		file1, err := workflowService.TemplateParse(deployKey, project, consts.Deploy)
		if err == nil {
			workflowDeployRes.ExecFile = file1
			workflowService.UpdateWorkflow(workflowDeployRes)
		}
	}

	Success(id, g)
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
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	userAny, _ := g.Get("user")
	user, _ := userAny.(db2.User)
	var userVo vo.UserAuth
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
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	var userVo vo.UserAuth
	userAny, _ := g.Get("user")
	user, _ := userAny.(db2.User)
	copier.Copy(&userVo, &user)
	err = workflowService.ExecProjectBuildWorkflow(projectId, userVo)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	Success("", g)
}

func (h *HandlerServer) projectWorkflowDeploy(g *gin.Context) {
	projectIdStr := g.Param("id")
	projectId, err := uuid.FromString(projectIdStr)
	if err != nil {
		Fail("projectId is empty or invalid", g)
		return
	}
	workflowIdStr := g.Param("workflowId")
	detailIdStr := g.Param("detailId")
	workflowId, err := strconv.Atoi(workflowIdStr)
	if err != nil {
		Fail("workflow id is empty or invalid", g)
		return
	}
	detailId, err := strconv.Atoi(detailIdStr)
	if err != nil {
		Fail("detail id is empty or invalid", g)
		return
	}
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	userAny, _ := g.Get("user")
	user, _ := userAny.(db2.User)
	var userVo vo.UserAuth
	copier.Copy(&userVo, &user)
	data, err := workflowService.ExecProjectDeployWorkflow(projectId, workflowId, detailId, userVo)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	Success(data, g)
}
func (h *HandlerServer) configContainerDeploy(g *gin.Context) {
	projectIdStr := g.Param("id")
	if projectIdStr == "" {
		Fail("projectId is empty or invalid", g)
		return
	}
	workflowIdStr := g.Param("workflowId")
	workflowId, err := strconv.Atoi(workflowIdStr)
	if err != nil {
		Fail("workflow id is empty or invalid", g)
		return
	}
	containerDeployService := application.GetBean[*service.ContainerDeployService]("containerDeployService")
	data := containerDeployService.CheckDeployParam(projectIdStr, workflowId)
	Success(data, g)
}

func (h *HandlerServer) containerDeploy(g *gin.Context) {
	projectIdStr := g.Param("id")
	projectId, err := uuid.FromString(projectIdStr)
	if err != nil {
		Fail("projectId is empty or invalid", g)
		return
	}
	workflowIdStr := g.Param("workflowId")
	detailIdStr := g.Param("detailId")
	workflowId, err := strconv.Atoi(workflowIdStr)
	if err != nil {
		Fail("workflow id is empty or invalid", g)
		return
	}
	detailId, err := strconv.Atoi(detailIdStr)
	if err != nil {
		Fail("detail id is empty or invalid", g)
		return
	}
	containerDeployService := application.GetBean[*service.ContainerDeployService]("containerDeployService")
	deployParam := parameter.K8sDeployParam{}
	err = g.BindJSON(&deployParam)
	if err != nil {
		deployData, err := containerDeployService.QueryDeployParam(projectIdStr, workflowId)
		if err != nil {
			Fail("deploy param is empty", g)
			return
		}
		copier.Copy(&deployParam, &deployData)
	} else {
		err = containerDeployService.SaveDeployParam(projectId, workflowId, deployParam)
		if err != nil {
			Fail("save deploy param failed", g)
			return
		}
	}
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	userAny, _ := g.Get("user")
	user, _ := userAny.(db2.User)
	var userVo vo.UserAuth
	copier.Copy(&userVo, &user)
	data, err := workflowService.ExecContainerDeploy(projectId, workflowId, detailId, userVo, deployParam)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	Success(data, g)
}

func (h *HandlerServer) projectContract(g *gin.Context) {

	projectId := g.Param("id")
	if projectId == "" {
		Fail("projectId is empty or invalid", g)
		return
	}

	query := g.Query("query")
	version := g.Query("version")
	network := g.Query("network")
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(g.DefaultQuery("size", "10"))

	contractService := application.GetBean[*service.ContractService]("contractService")

	result, err := contractService.QueryContracts(projectId, query, version, network, page, size)

	if err != nil {
		Fail(err.Error(), g)
		return
	}

	Success(result, g)

}

func (h *HandlerServer) projectReport(g *gin.Context) {
	projectId := g.Param("id")
	if projectId == "" {
		Fail("projectId is empty or invalid", g)
		return
	}

	Type := g.DefaultQuery("type", "")
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(g.DefaultQuery("size", "10"))

	reportService := application.GetBean[*service.ReportService]("reportService")

	result, err := reportService.QueryReports(projectId, Type, page, size)

	if err != nil {
		Fail(err.Error(), g)
		return
	}

	Success(result, g)

}
func (h *HandlerServer) projectFrontendReports(g *gin.Context) {
	projectId := g.Param("id")
	if projectId == "" {
		Fail("projectId is empty or invalid", g)
		return
	}
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(g.DefaultQuery("size", "10"))
	reportService := application.GetBean[*service.ReportService]("reportService")

	result, err := reportService.QueryFrontendReports(projectId, page, size)

	if err != nil {
		Fail(err.Error(), g)
		return
	}

	Success(result, g)
}
func (h *HandlerServer) projectPackages(g *gin.Context) {
	projectId := g.Param("id")
	if projectId == "" {
		Fail("projectId is empty or invalid", g)
		return
	}
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(g.DefaultQuery("size", "10"))
	frontendPackageService := application.GetBean[*service.FrontendPackageService]("frontendPackageService")
	result, err := frontendPackageService.QueryFrontendPackages(projectId, page, size)
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
	tokenAny, _ := gin.Get("token")
	token, _ := tokenAny.(string)
	userAny, _ := gin.Get("user")
	user, _ := userAny.(db2.User)
	updateData.UserId = int(user.Id)
	project, err := h.projectService.GetProject(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	githubService := application.GetBean[*service.GithubService]("githubService")
	repo, res, err := githubService.UpdateRepo(token, user.Username, project.Name, updateData.Name)
	if err != nil {
		if res != nil {
			if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
				Failed(http.StatusUnauthorized, "access not authorized", gin)
				return
			}
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
	tokenAny, _ := gin.Get("token")
	token, _ := tokenAny.(string)
	githubService := application.GetBean[*service.GithubService]("githubService")
	data := githubService.CheckName(token, checkData.Owner, checkData.Name)
	Success(data, gin)
}

func (h *HandlerServer) createProjectByCode(gin *gin.Context) {
	createData := parameter.CreateByCodeParam{}
	err := gin.BindJSON(&createData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	tokenAny, _ := gin.Get("token")
	token, _ := tokenAny.(string)
	userAny, _ := gin.Get("user")
	user, _ := userAny.(db2.User)
	githubService := application.GetBean[*service.GithubService]("githubService")

	repo, res, err := githubService.GetRepo(token, user.Username, createData.Name)
	if err != nil {
		if res != nil {
			if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
				Failed(http.StatusUnauthorized, "access not authorized", gin)
				return
			}
		}
		repo, res, err = githubService.CreateRepository(token, createData.Name)
		if err != nil {
			if res != nil {
				if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
					Failed(http.StatusUnauthorized, "access not authorized", gin)
					return
				}
			}
			Fail(err.Error(), gin)
			return
		}
	}
	err = githubService.CommitAndPush(token, *repo.CloneURL, user.Username, user.UserEmail, consts.TemplateUrl, consts.TemplateRepoName)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	// add file
	_, res, err = githubService.AddFile(token, user.Username, createData.Name, createData.Content, createData.FileName)
	if err != nil {
		if res != nil {
			if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
				Failed(http.StatusUnauthorized, "access not authorized", gin)
				return
			}
		}
		Fail(err.Error(), gin)
		return
	}
	//create project
	data := vo.CreateProjectParam{
		Name:        createData.Name,
		Type:        createData.Type,
		TemplateUrl: *repo.CloneURL,
		FrameType:   createData.FrameType,
		UserId:      int64(user.Id),
	}
	id, err := h.projectService.CreateProject(data)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	project, err := h.projectService.GetProject(id.String())
	if err != nil {
		Fail(err.Error(), gin)
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
		Success(id, gin)
		return
	}
	checkKey := workflowService.GetWorkflowKey(id.String(), workflowCheckRes.Id)
	file, err := workflowService.TemplateParse(checkKey, project, consts.Check)
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
		Success(id, gin)
		return
	}
	buildKey := workflowService.GetWorkflowKey(id.String(), workflowBuildRes.Id)
	file1, err := workflowService.TemplateParse(buildKey, project, consts.Build)
	if err == nil {
		workflowBuildRes.ExecFile = file1
		workflowService.UpdateWorkflow(workflowBuildRes)
	}

	if project.Type == uint(consts.FRONTEND) {
		workflowDeployData := parameter.SaveWorkflowParam{
			ProjectId:  id,
			Type:       consts.Deploy,
			ExecFile:   "",
			LastExecId: 0,
		}
		workflowDeployRes, err := workflowService.SaveWorkflow(workflowDeployData)
		if err != nil {
			Success(id, gin)
			return
		}
		deployKey := workflowService.GetWorkflowKey(id.String(), workflowDeployRes.Id)
		file1, err := workflowService.TemplateParse(deployKey, project, consts.Deploy)
		if err == nil {
			workflowDeployRes.ExecFile = file1
			workflowService.UpdateWorkflow(workflowDeployRes)
		}
	}

	Success(id, gin)
}
