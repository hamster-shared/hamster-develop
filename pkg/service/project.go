package service

import (
	"errors"
	"fmt"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
)

type IProjectService interface {
	GetProjects(userId int, keyword string, page, size int) (*vo.ProjectPage, error)
	CreateProject(createData vo.CreateProjectParam) (uuid.UUID, error)
	GetProject(id string) (*vo.ProjectDetailVo, error)
	UpdateProject(id string, updateData vo.UpdateProjectParam) error
	DeleteProject(id string) error
}

type ProjectService struct {
	db *gorm.DB
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (p *ProjectService) Init(db *gorm.DB) {
	p.db = db
}

func (p *ProjectService) GetProjects(userId int, keyword string, page, size int) (*vo.ProjectPage, error) {
	var total int64
	var projectPage vo.ProjectPage
	var projects []db2.Project
	var projectList []vo.ProjectListVo
	tx := p.db.Model(db2.Project{}).Where("user_id = ?", userId)
	if keyword != "" {
		tx = tx.Where("name like ? ", "%"+keyword+"%")
	}

	result := tx.Order("create_time DESC").Offset((page - 1) * size).Limit(size).Find(&projects)
	if result.Error != nil {
		return &projectPage, result.Error
	}
	if len(projects) > 0 {
		for _, project := range projects {
			var data vo.ProjectListVo
			var recentBuild vo.RecentBuildVo
			var recentCheck vo.RecentCheckVo
			var recentDeploy vo.RecentDeployVo
			var workflowBuildData db2.WorkflowDetail
			var workflowCheckData db2.WorkflowDetail
			copier.Copy(&data, &project)
			err := p.db.Model(db2.WorkflowDetail{}).Where("project_id = ? and type = ?", project.Id, consts.Check).Order("start_time DESC").Limit(1).Find(&workflowCheckData).Error
			if err == nil {
				copier.Copy(&recentCheck, workflowCheckData)
			}
			err = p.db.Model(db2.WorkflowDetail{}).Where("project_id = ? and type = ?", project.Id, consts.Build).Order("start_time DESC").Limit(1).Find(&workflowBuildData).Error
			if err == nil {
				copier.Copy(&recentBuild, &workflowBuildData)
				var contractData db2.Contract
				err = p.db.Model(db2.Contract{}).Where("project_id = ?", project.Id).Order("build_time DESC").Limit(1).Find(&contractData).Error
				if err == nil {
					recentBuild.Version = contractData.Version
				}
			}
			var deployData db2.ContractDeploy
			err = p.db.Model(db2.ContractDeploy{}).Where("project_id = ?", project.Id).Order("deploy_time DESC").Limit(1).Find(&deployData).Error
			if err == nil {
				copier.Copy(&recentDeploy, &deployData)
			}
			data.RecentBuild = recentBuild
			data.RecentCheck = recentCheck
			data.RecentDeploy = recentDeploy
			projectList = append(projectList, data)
		}
		tx.Count(&total)
	}
	projectPage.Data = projectList
	projectPage.Total = int(total)
	projectPage.Page = page
	projectPage.PageSize = size
	return &projectPage, nil
}

func (p *ProjectService) CreateProject(createData vo.CreateProjectParam) (uuid.UUID, error) {
	var project db2.Project
	err := p.db.Where("name=? and user_id=?", createData.Name, createData.UserId).First(&project).Error
	if err == gorm.ErrRecordNotFound {
		project.Name = createData.Name
		project.UserId = createData.UserId
		project.Creator = createData.UserId
		project.CreateTime = time.Now()
		project.UpdateTime = time.Now()
		project.FrameType = createData.FrameType
		project.Type = uint(createData.Type)
		project.RepositoryUrl = createData.TemplateUrl
		project.Branch = "main"
		p.db.Create(&project)
		return project.Id, nil
	}
	return project.Id, errors.New(fmt.Sprintf("application:%s already exists", createData.Name))
}

func (p *ProjectService) GetProject(id string) (*vo.ProjectDetailVo, error) {
	var data db2.Project
	var detail vo.ProjectDetailVo
	result := p.db.Where("id = ? ", id).First(&data)
	if result.Error != nil {
		return &detail, result.Error
	}
	var recentBuild vo.RecentBuildVo
	var recentCheck vo.RecentCheckVo
	var recentDeploy vo.RecentDeployVo
	var workflowBuildData db2.WorkflowDetail
	var workflowCheckData db2.WorkflowDetail
	copier.Copy(&detail, &data)
	err := p.db.Model(db2.WorkflowDetail{}).Where("project_id = ? and type = ?", data.Id, consts.Check).Order("start_time DESC").Limit(1).Find(&workflowCheckData).Error
	if err == nil {
		copier.Copy(&recentCheck, workflowCheckData)
	}
	err = p.db.Model(db2.WorkflowDetail{}).Where("project_id = ? and type = ?", data.Id, consts.Build).Order("start_time DESC").Limit(1).Find(&workflowBuildData).Error
	if err == nil {
		copier.Copy(&recentBuild, &workflowBuildData)
		var contractData db2.Contract
		err = p.db.Model(db2.Contract{}).Where("project_id = ?", data.Id).Order("build_time DESC").Limit(1).Find(&contractData).Error
		if err == nil {
			recentBuild.Version = contractData.Version
		}
	}
	var deployData db2.ContractDeploy
	err = p.db.Model(db2.ContractDeploy{}).Where("project_id = ?", data.Id).Order("deploy_time DESC").Limit(1).Find(&deployData).Error
	if err == nil {
		copier.Copy(&recentDeploy, &deployData)
	}
	detail.RecentBuild = recentBuild
	detail.RecentCheck = recentCheck
	detail.RecentDeploy = recentDeploy
	return &detail, nil
}

func (p *ProjectService) UpdateProject(id string, updateData vo.UpdateProjectParam) error {
	var data db2.Project
	err := p.db.Where("name=? and user_id = ?", updateData.Name, updateData.UserId).First(&data).Error
	if err == gorm.ErrRecordNotFound {
		result := p.db.Model(data).Where("id = ?", id).Updates(db2.Project{Name: updateData.Name, RepositoryUrl: updateData.RepositoryUrl, UpdateTime: time.Now(), UpdateUser: uint(updateData.UserId)})
		if result.Error != nil {
			return result.Error
		}
		return nil
	}
	return errors.New(fmt.Sprintf("application:%s already exists", updateData.Name))
}

func (p *ProjectService) DeleteProject(id string) error {
	//result := p.db.Debug().Where("id = ?", id).Delete(&db2.Project{})
	deleteSql := "delete t,tw,twd,tc,tcd,tr from t_project t left join t_workflow tw on t.id = tw.project_id left join t_workflow_detail twd on t.id = twd.project_id  left join t_contract tc on t.id = tc.project_id left join t_contract_deploy tcd on t.id = tcd.project_id left join t_report tr on t.id = tr.project_id where t.id = ?"
	result := p.db.Exec(deleteSql, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
