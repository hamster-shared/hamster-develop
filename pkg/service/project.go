package service

import (
	"errors"
	"fmt"
	"github.com/hamster-shared/a-line/pkg/consts"
	db2 "github.com/hamster-shared/a-line/pkg/db"
	"github.com/hamster-shared/a-line/pkg/vo"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
)

type IProjectService interface {
	GetProjects(userId int, keyword string, page, size int) (*vo.ProjectPage, error)
	CreateProject(createData vo.CreateProjectParam) (uint, error)
	GetProject(id int) (*db2.Project, error)
	UpdateProject(id int, updateData vo.UpdateProjectParam) error
	DeleteProject(id int) error
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
	result := tx.Offset((page - 1) * size).Limit(size).Find(&projects).Count(&total)
	if result.Error != nil {
		return &projectPage, result.Error
	}
	if len(projects) > 0 {
		for _, project := range projects {
			var data vo.ProjectListVo
			var recentBuild vo.RecentBuildVo
			var recentCheck vo.RecentCheckVo
			var recentDeploy vo.RecentDeployVo
			var workflowData db2.WorkflowDetail
			copier.Copy(&data, &project)
			err := p.db.Model(db2.WorkflowDetail{}).Where("project_id = ? and type = ?", project.Id, consts.Check).Order("start_time DESC").Limit(1).Find(&workflowData).Error
			if err == nil {
				copier.Copy(&recentCheck, workflowData)
			}
			err = p.db.Model(db2.WorkflowDetail{}).Where("project_id = ? and type = ?", project.Id, consts.Build).Order("start_time DESC").Limit(1).Find(&workflowData).Error
			if err == nil {
				copier.Copy(&recentBuild, &workflowData)
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
	}
	projectPage.Data = projectList
	projectPage.Total = int(total)
	projectPage.Page = page
	projectPage.PageSize = size
	return &projectPage, nil
}

func (p *ProjectService) CreateProject(createData vo.CreateProjectParam) (uint, error) {
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
		p.db.Create(&project)
		return project.Id, nil
	}
	return project.Id, errors.New(fmt.Sprintf("application:%s already exists", createData.Name))
}

func (p *ProjectService) GetProject(id int) (*db2.Project, error) {
	var data db2.Project
	result := p.db.Where("id = ? ", id).First(&data)
	if result.Error != nil {
		return &data, result.Error
	}
	return &data, nil
}

func (p *ProjectService) UpdateProject(id int, updateData vo.UpdateProjectParam) error {
	var data db2.Project
	err := p.db.Where("name=? and user_id = ?", updateData.Name, updateData.UserId).First(&data).Error
	if err == gorm.ErrRecordNotFound {
		result := p.db.Model(data).Where("id = ?", id).Updates(db2.Project{Name: updateData.Name, UpdateTime: time.Now(), UpdateUser: uint(updateData.UserId)})
		if result.Error != nil {
			return result.Error
		}
		return nil
	}
	return errors.New(fmt.Sprintf("application:%s already exists", updateData.Name))
}

func (p *ProjectService) DeleteProject(id int) error {
	result := p.db.Debug().Where("id = ?", id).Delete(&db2.Project{})
	if result.Error != nil {
		return result.Error
	}
	//todo delete workflow,workflow detail,contract,report
	return nil
}
