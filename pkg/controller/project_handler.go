package controller

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/a-line/pkg/application"
	"github.com/hamster-shared/a-line/pkg/controller/parameters"
	"github.com/hamster-shared/a-line/pkg/db"
	"github.com/young2j/gocopy"
	"gorm.io/gorm"
)

//go:embed templates
var temp embed.FS

func (h *HandlerServer) projectList(g *gin.Context) {

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

func (h *HandlerServer) projectDetail(g *gin.Context) {

}

func (h *HandlerServer) projectWorkflowCheck(g *gin.Context) {

}

func (h *HandlerServer) projectWorkflowBuild(g *gin.Context) {

}

func (h *HandlerServer) projectWorkflowDeploy(g *gin.Context) {

}

func (h *HandlerServer) projectContract(g *gin.Context) {

}

func (h *HandlerServer) projectReport(g *gin.Context) {

}

func (h *HandlerServer) updateProject(g *gin.Context) {

}

func (h *HandlerServer) deleteProject(g *gin.Context) {

}
