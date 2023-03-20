package controller

import (
	"embed"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/aline-engine/logger"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/service"
	"github.com/hamster-shared/hamster-develop/pkg/utils"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/jinzhu/copier"
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
	logger.Tracef("projectWorkflowCheck")

	projectIdStr := g.Param("id")
	projectId, err := uuid.FromString(projectIdStr)
	if err != nil {
		logger.Errorf("projectWorkflowCheck error: %s", err.Error())
		Fail(err.Error(), g)
		return
	}
	userAny, _ := g.Get("user")
	user, _ := userAny.(db2.User)
	var userVo vo.UserAuth
	copier.Copy(&userVo, &user)
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	err = workflowService.ExecProjectCheckWorkflow(projectId, userVo)
	if err != nil {
		logger.Errorf("projectWorkflowCheck error: %s", err.Error())
		Fail(err.Error(), g)
		return
	}
	Success("", g)
}

func (h *HandlerServer) projectWorkflowBuild(g *gin.Context) {
	logger.Tracef("projectWorkflowBuild")
	projectIdStr := g.Param("id")
	projectId, err := uuid.FromString(projectIdStr)
	if err != nil {
		logger.Errorf("projectWorkflowBuild error: %s", err.Error())
		Fail("projectId is empty or invalid", g)
		return
	}
	workflowService := application.GetBean[*service.WorkflowService]("workflowService")
	var userVo vo.UserAuth
	userAny, _ := g.Get("user")
	user, _ := userAny.(db2.User)
	copier.Copy(&userVo, &user)
	data, err := workflowService.ExecProjectBuildWorkflow(projectId, userVo)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	Success(data, g)
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
	containerDeployService := application.GetBean[*service.ContainerDeployService]("containerDeployService")
	data := containerDeployService.CheckDeployParam(projectIdStr)
	Success(data, g)
}

func (h *HandlerServer) updateContainerDeploy(g *gin.Context) {
	projectIdStr := g.Param("id")
	if projectIdStr == "" {
		Fail("projectId is empty or invalid", g)
		return
	}
	projectId, err := uuid.FromString(projectIdStr)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	deployParam := parameter.K8sDeployParam{}
	err = g.BindJSON(&deployParam)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	containerDeployService := application.GetBean[*service.ContainerDeployService]("containerDeployService")
	err = containerDeployService.UpdateContainerDeploy(projectId, deployParam)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	Success("", g)

}

func (h *HandlerServer) getContainerDeploy(g *gin.Context) {
	projectIdStr := g.Param("id")
	if projectIdStr == "" {
		Fail("projectId is empty or invalid", g)
		return
	}
	containerDeployService := application.GetBean[*service.ContainerDeployService]("containerDeployService")
	data := containerDeployService.GetContainerDeploy(projectIdStr)
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
	deployData, err := containerDeployService.QueryDeployParam(projectIdStr)
	if err != nil {
		Fail("please config deploy param", g)
		return
	}
	copier.Copy(&deployParam, &deployData)
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

func (h *HandlerServer) queryAptosParams(g *gin.Context) {
	projectID := g.Param("id")
	if projectID == "" {
		Fail("projectId is empty or invalid", g)
		return
	}
	// 先去数据库查询，如果有，直接返回
	params, err := h.projectService.GetProjectParams(projectID)
	if err == nil {
		if params != "" {
			params, err := utils.KeyValuesFromString(params)
			if err != nil {
				Fail(err.Error(), g)
				return
			}
			Success(params, g)
			return
		}
	}

	// 先查询到此项目的 github 仓库信息
	data, err := h.projectService.GetProject(projectID)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	rawUrl := getGithubRawUrl(data.RepositoryUrl, data.Branch, "Move.toml")
	// 获取到这个文件，解析它，得到里面的内容
	resp, err := http.Get(rawUrl)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	// if 404 not found
	if resp.StatusCode == 404 {
		Fail("Move.toml not found", g)
		return
	}
	// 解析这个文件
	moveToml, err := utils.ParseMoveTomlWithString(string(body))
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	keyValues := moveToml.GetAddressField()
	if len(keyValues) == 0 {
		Fail("no address field", g)
		return
	}
	updateData := vo.UpdateProjectParams{
		Params: utils.KeyValuesToString(keyValues),
	}
	// 保存到数据库一份
	err = h.projectService.UpdateProjectParams(projectID, updateData)
	if err != nil {
		logger.Errorf("update project params error: %s", err.Error())
	}
	// 返回给前端
	Success(keyValues, g)
}

type AptosParams struct {
	Params []utils.KeyValue `json:"params"`
}

func (h *HandlerServer) saveAptosParams(g *gin.Context) {
	projectID := g.Param("id")
	if projectID == "" {
		Fail("projectId is empty or invalid", g)
		return
	}
	var params AptosParams
	err := g.BindJSON(&params)
	if err != nil {
		Fail(err.Error(), g)
		return
	}
	// 保存到数据库一份
	updateData := vo.UpdateProjectParams{
		Params: utils.KeyValuesToString(params.Params),
	}
	err = h.projectService.UpdateProjectParams(projectID, updateData)
	if err != nil {
		logger.Errorf("update project params error: %s", err.Error())
		Fail(err.Error(), g)
		return
	}
	Success(nil, g)
}

func getGithubRawUrl(reportUrl, branch, path string) string {
	url := strings.Replace(reportUrl, "github.com", "raw.githubusercontent.com", 1)
	// 如果以 .git 结尾，去掉
	url = strings.TrimSuffix(url, ".git")
	url = strings.Replace(url, "www.", "", 1)
	url = strings.Join([]string{url, branch, path}, "/")
	return url
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
