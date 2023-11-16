package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/google/go-github/v48/github"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/utils"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type IProjectService interface {
	GetProjects(userId int, keyword string, page, size, projectType int) (*vo.ProjectPage, error)
	HandleProjectsByUserId(user db2.User, page, size int, token, filter string) (vo.RepoListPage, error)
	CreateProject(createData vo.CreateProjectParam) (uuid.UUID, error)
	GetProject(id string) (*vo.ProjectDetailVo, error)
	UpdateProject(id string, updateData vo.UpdateProjectParam) error
	DeleteProject(id string) error
	UpdateProjectParams(id string, updateData vo.UpdateProjectParams) error
	GetProjectParams(id string) (string, error)
	GetProjectById(id string) (*db2.Project, error)
	//ParsingFrame(repoContents []*github.RepositoryContent, name, userName, token string) (uint, error)
	ParsingEVMFrame(repoContents []*github.RepositoryContent) (consts.EVMFrameType, error)
	GetChainNetworkList() ([]db2.ChainNetwork, error)
	GetChainNetworkByName(name string) (db2.ChainNetwork, error)
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

func (p *ProjectService) GetProjects(userId int, keyword string, page, size, projectType int) (*vo.ProjectPage, error) {
	var total int64
	var projectPage vo.ProjectPage
	var projects []db2.Project
	var projectList []vo.ProjectListVo
	tx := p.db.Model(db2.Project{}).Where("user_id = ? and type = ?", userId, projectType)
	if keyword != "" {
		tx = tx.Where("name like ? ", "%"+keyword+"%")
	}

	result := tx.Order("create_time DESC").Offset((page - 1) * size).Limit(size).Find(&projects).Offset(-1).Limit(-1).Count(&total)
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
			_ = copier.Copy(&data, &project)
			err := p.db.Model(db2.WorkflowDetail{}).Where("project_id = ? and type = ?", project.Id, consts.Check).Order("start_time DESC").Limit(1).Find(&workflowCheckData).Error
			if err == nil {
				_ = copier.Copy(&recentCheck, workflowCheckData)
			}
			err = p.db.Model(db2.WorkflowDetail{}).Where("project_id = ? and type = ?", project.Id, consts.Build).Order("start_time DESC").Limit(1).Find(&workflowBuildData).Error
			if err == nil {
				_ = copier.Copy(&recentBuild, &workflowBuildData)
				if projectType == int(consts.CONTRACT) {
					var contractData db2.Contract
					err = p.db.Model(db2.Contract{}).Where("project_id = ?", project.Id).Order("build_time DESC").Limit(1).Find(&contractData).Error
					if err == nil {
						recentBuild.Version = contractData.Version
					}
				}
			}
			//recentDeploy
			var workflowDeployData db2.WorkflowDetail
			if projectType == int(consts.CONTRACT) {
				if project.FrameType == consts.InternetComputer {
					err = p.db.Model(db2.WorkflowDetail{}).Where("project_id = ? and type = ?", project.Id, consts.Deploy).Order("create_time DESC").Limit(1).Find(&workflowDeployData).Error
					if err == nil {
						recentDeploy.Id = workflowDeployData.Id
						recentDeploy.Status = workflowDeployData.Status
						recentDeploy.DeployTime = workflowDeployData.StartTime
						var backendDeploy db2.BackendDeploy
						err = p.db.Model(db2.BackendDeploy{}).Where("project_id = ? and workflow_detail_id = ? ", data.Id, workflowDeployData.Id).Order("deploy_time DESC").Limit(1).First(&backendDeploy).Error
						if err == nil {
							_ = copier.Copy(&recentDeploy, &backendDeploy)
						}
					}
				} else {
					var deployData db2.ContractDeploy
					err = p.db.Model(db2.ContractDeploy{}).Where("project_id = ?", project.Id).Order("deploy_time DESC").Limit(1).Find(&deployData).Error
					if err == nil {
						_ = copier.Copy(&recentDeploy, &deployData)
					}
				}
				data.RecentDeploy = recentDeploy
			} else {
				var packageDeploy vo.PackageDeployVo
				var deployData db2.FrontendDeploy
				err = p.db.Model(db2.WorkflowDetail{}).Where("project_id = ? and type = ?", project.Id, consts.Deploy).Order("create_time DESC").Limit(1).Find(&workflowDeployData).Error
				if err == nil {
					copier.Copy(&packageDeploy, workflowDeployData)
					err = p.db.Model(db2.FrontendDeploy{}).Where("project_id = ? and workflow_detail_id = ? ", project.Id, workflowDeployData.Id).Order("deploy_time DESC").Limit(1).Find(&deployData).Error
					if err == nil {
						packageDeploy.PackageId = deployData.PackageId
						packageDeploy.Version = deployData.Version
					}
				}
				data.RecentDeploy = packageDeploy
			}
			data.RecentBuild = recentBuild
			data.RecentCheck = recentCheck
			projectList = append(projectList, data)
		}
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
		project.Branch = createData.Branch
		project.DeployType = createData.DeployType
		project.LabelDisplay = createData.LabelDisplay
		project.GistId = createData.GistId
		project.DefaultFile = createData.DefaultFile
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
	_ = copier.Copy(&detail, &data)
	err := p.db.Model(db2.WorkflowDetail{}).Where("project_id = ? and type = ?", data.Id, consts.Check).Order("start_time DESC").Limit(1).Find(&workflowCheckData).Error
	if err == nil {
		_ = copier.Copy(&recentCheck, workflowCheckData)
	}
	err = p.db.Model(db2.WorkflowDetail{}).Where("project_id = ? and type = ?", data.Id, consts.Build).Order("start_time DESC").Limit(1).Find(&workflowBuildData).Error
	if err == nil {
		_ = copier.Copy(&recentBuild, &workflowBuildData)
		if data.Type == uint(consts.CONTRACT) {
			var contractData db2.Contract
			err = p.db.Model(db2.Contract{}).Where("project_id = ?", data.Id).Order("build_time DESC").Limit(1).Find(&contractData).Error
			if err == nil {
				recentBuild.Version = contractData.Version
			}
		}
	}
	var workflowDeployData db2.WorkflowDetail
	if data.Type == uint(consts.CONTRACT) {
		if data.FrameType == consts.InternetComputer {
			err = p.db.Model(db2.WorkflowDetail{}).Where("project_id = ? and type = ?", data.Id, consts.Deploy).Order("create_time DESC").Limit(1).Find(&workflowDeployData).Error
			if err == nil {
				recentDeploy.Id = workflowDeployData.Id
				recentDeploy.Status = workflowDeployData.Status
				recentDeploy.DeployTime = workflowDeployData.StartTime
				var backendDeploy db2.BackendDeploy
				err = p.db.Model(db2.BackendDeploy{}).Where("project_id = ? and workflow_detail_id = ? ", data.Id, workflowDeployData.Id).Order("deploy_time DESC").Limit(1).First(&backendDeploy).Error
				if err == nil {
					_ = copier.Copy(&recentDeploy, &backendDeploy)
				}
			}
		} else {
			var deployData db2.ContractDeploy
			err = p.db.Model(db2.ContractDeploy{}).Where("project_id = ?", data.Id).Order("deploy_time DESC").Limit(1).Find(&deployData).Error
			if err == nil {
				_ = copier.Copy(&recentDeploy, &deployData)
			}
		}
		detail.RecentDeploy = recentDeploy
	} else {
		var packageDeploy vo.PackageDeployVo
		var deployData db2.FrontendDeploy
		err = p.db.Model(db2.WorkflowDetail{}).Where("project_id = ? and type = ?", data.Id, consts.Deploy).Order("create_time DESC").Limit(1).Find(&workflowDeployData).Error
		if err == nil {
			copier.Copy(&packageDeploy, workflowDeployData)
			err = p.db.Model(db2.FrontendDeploy{}).Where("project_id = ? and workflow_detail_id = ? ", data.Id, workflowDeployData.Id).Order("deploy_time DESC").Limit(1).Find(&deployData).Error
			if err == nil {
				packageDeploy.PackageId = deployData.PackageId
				packageDeploy.Version = deployData.Version
			}
		}
		detail.RecentDeploy = packageDeploy
	}
	detail.RecentBuild = recentBuild
	detail.RecentCheck = recentCheck
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

func (p *ProjectService) GetProjectParams(id string) (string, error) {
	data, err := p.GetProjectById(id)
	if err != nil {
		return "", err
	}
	return data.Params, err
}

func (p *ProjectService) UpdateProjectParams(id string, updateData vo.UpdateProjectParams) error {
	var data db2.Project
	err := p.db.Model(data).Where("id = ?", id).Updates(db2.Project{Params: updateData.Params, UpdateTime: time.Now()}).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ProjectService) DeleteProject(id string) error {
	//result := p.db.Debug().Where("id = ?", id).Delete(&db2.Project{})
	deleteSql := "delete t,tw,twd,tc,tcd,tr from t_project t left join t_workflow tw on t.id = tw.project_id left join t_workflow_detail twd on t.id = twd.project_id  left join t_contract tc on t.id = tc.project_id left join t_contract_deploy tcd on t.id = tcd.project_id left join t_report tr on t.id = tr.project_id where t.id = ?"
	result := p.db.Exec(deleteSql, id)
	if result.Error != nil {
		return result.Error
	}
	filePath := fmt.Sprintf("rm -rf %s*", id)
	go func() {
		//delete workdir
		deleteWorkCmd := exec.Command("/bin/bash", "-c", filePath)
		deleteWorkCmd.Dir = utils.DefaultWorkDir()
		deleteWorkCmd.Run()
		//delete pipelines
		pipelinePath := filepath.Join(utils.DefaultPipelineDir(), consts.JOB_DIR_NAME)
		deletePipeCmd := exec.Command("/bin/bash", "-c", filePath)
		deletePipeCmd.Dir = pipelinePath
		deletePipeCmd.Run()
	}()
	return nil
}

func (p *ProjectService) GetProjectById(id string) (*db2.Project, error) {
	var data db2.Project
	result := p.db.Where("id = ? ", id).First(&data)
	return &data, result.Error
}

func (p *ProjectService) GetChainNetworkList() ([]db2.ChainNetwork, error) {
	var list []db2.ChainNetwork
	err := p.db.Model(db2.ChainNetwork{}).Find(&list).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func (p *ProjectService) GetChainNetworkByName(name string) (db2.ChainNetwork, error) {
	var chainNetwork db2.ChainNetwork
	err := p.db.Model(db2.ChainNetwork{}).Where("chain_name = ?", name).First(&chainNetwork).Error
	if err != nil {
		return chainNetwork, err
	}
	return chainNetwork, nil
}

func (p *ProjectService) HandleProjectsByUserId(user db2.User, page, size int, token, filter string) (vo.RepoListPage, error) {
	var projects []db2.Project
	err := p.db.Model(db2.Project{}).Where("user_id = ?", user.Id).Find(&projects).Error
	if err != nil {
		return vo.RepoListPage{}, err
	}
	githubService := application.GetBean[*GithubService]("githubService")
	if len(projects) > 0 {
		data, err := githubService.GetRepoList(token, user.Username, filter, 1, 1000)
		if err != nil {
			return vo.RepoListPage{}, err
		}
		res := removeElements(data.Data, projects)
		start, end := utils.SlicePage(int64(page), int64(size), int64(len(res)))
		result := res[start:end]
		data.Total = len(res)
		data.Data = result
		data.PageSize = size
		data.Page = page
		return data, nil
	}
	repoListVo, err := githubService.GetRepoList(token, user.Username, filter, page, size)
	if err != nil {
		return vo.RepoListPage{}, err
	}
	return repoListVo, nil

}

func removeElements(arr1 []vo.RepoVo, arr2 []db2.Project) []vo.RepoVo {
	var result []vo.RepoVo
	for _, repoVo := range arr1 {
		found := false
		for _, project := range arr2 {
			if repoVo.Name == project.Name {
				found = true
				break
			}
		}
		if !found {
			result = append(result, repoVo)
		}
	}
	return result
}

// ParsingFrame only parsing EVM frame now
//func (p *ProjectService) ParsingFrame(repoContents []*github.RepositoryContent, name, userName, token string) (uint, error) {
//	for _, v := range repoContents {
//		fileName := v.GetName()
//		if strings.Contains(fileName, "cairo") {
//			return consts.StarkWare, nil
//		} else if strings.Contains(fileName, "Move.toml") {
//			return parsingToml(v, name, userName, token)
//		} else if strings.Contains(fileName, "truffle-config.js") || strings.Contains(fileName, "foundry.toml") || strings.Contains(fileName, "hardhat.config.js") {
//			evmFrameType := getEvmFrameType(fileName)
//			if evmFrameType == 0 {
//				return 0, fmt.Errorf("parsing evm frame type failed")
//			} else {
//				return evmFrameType, nil
//			}
//		}
//	}
//	return 0, fmt.Errorf("parsing frame error")
//}

type FrameType struct {
	isTruffle bool
	isHardhat bool
	isFoundry bool
	isWaffle  bool
}

func (ft *FrameType) GetEvmFrameType() (consts.EVMFrameType, error) {
	if ft.isTruffle {
		return consts.Truffle, nil
	} else if ft.isHardhat {
		return consts.Hardhat, nil
	} else if ft.isFoundry {
		return consts.Foundry, nil
	} else if ft.isWaffle {
		return consts.Waffle, nil
	} else {
		return 0, fmt.Errorf("parsing frame error")
	}
}

func (p *ProjectService) ParsingEVMFrame(repoContents []*github.RepositoryContent) (consts.EVMFrameType, error) {

	ft := FrameType{}

	for _, v := range repoContents {
		fileName := v.GetName()
		if strings.Contains(fileName, "truffle-config.js") || strings.Contains(fileName, "truffle-config.ts") {
			ft.isTruffle = true
		} else if strings.Contains(fileName, "hardhat.config.js") || strings.Contains(fileName, "hardhat.config.ts") {
			ft.isHardhat = true
		} else if strings.Contains(fileName, "foundry.toml") {
			ft.isFoundry = true
		} else if strings.Contains(fileName, ".waffle.json") {
			ft.isWaffle = true
		}
	}

	return ft.GetEvmFrameType()
}

func parsingToml(fileContent *github.RepositoryContent, name, userName, token string) (consts.ProjectFrameType, error) {
	githubService := application.GetBean[*GithubService]("githubService")
	content, err := githubService.GetFileContent(token, userName, name, fileContent.GetPath())
	if err != nil {
		return 0, err
	}
	var tomlData map[string]interface{}
	if err := toml.Unmarshal([]byte(content), &tomlData); err != nil {
		log.Printf("parsing toml failed: %s\n", err.Error())
		return 0, err
	}
	dependenciesData, ok := tomlData["dependencies"].(map[string]interface{})
	if !ok {
		log.Println("get move.toml dependencies failed")
		return 0, fmt.Errorf("get dependencies failed")
	}
	for key, _ := range dependenciesData {
		if strings.Contains(key, "aptos") || strings.Contains(key, "Aptos") {
			return consts.Aptos, nil
		}
		if strings.Contains(key, "Sui") || strings.Contains(key, "sui") {
			return consts.Sui, nil
		}
	}
	return 0, fmt.Errorf("dependencies did not have sui or aptos, it may be not sui or aptos")
}

func parsingPackageJson(fileContent *github.RepositoryContent, name, userName, token string) (consts.ProjectFrameType, error) {
	githubService := application.GetBean[*GithubService]("githubService")
	content, err := githubService.GetFileContent(token, userName, name, fileContent.GetPath())
	if err != nil {
		return 0, err
	}
	var packageData map[string]any
	if err := json.Unmarshal([]byte(content), &packageData); err != nil {
		return 0, err
	}
	if _, ok := packageData["dependencies"].(map[string]interface{})["vue"]; ok {
		return 1, nil
	} else if _, ok := packageData["dependencies"].(map[string]interface{})["react"]; ok {
		return 2, nil
	} else if _, ok := packageData["dependencies"].(map[string]interface{})["nuxt"]; ok {
		return 3, nil
	} else if _, ok := packageData["dependencies"].(map[string]interface{})["next"]; ok {
		return 4, nil
	}
	return 0, fmt.Errorf("canot ensure the frontend frame type")
}
