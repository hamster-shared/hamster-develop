package service

import (
	"errors"
	"github.com/hamster-shared/a-line/pkg/consts"
	"github.com/hamster-shared/a-line/pkg/model"
	"github.com/hamster-shared/a-line/pkg/utils"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type IProjectService interface {
	//CreateProject create project
	CreateProject(project *model.Project) error
	//UpdateProject update project
	UpdateProject(oldName string, project *model.Project) error
	//GetProjects get project list
	GetProjects(keyword string, page, size int) *model.ProjectPage
	//DeleteProject delete project
	DeleteProject(name string) error
}

type ProjectService struct {
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

// CreateProject create project
func (p *ProjectService) CreateProject(project *model.Project) error {
	//project directory path
	projectDirPath := filepath.Join(utils.DefaultConfigDir(), consts.PROJECT_DIR_NAME, project.Name)
	//determine whether the folder exists, and create it if it does not exist
	_, err := os.Stat(projectDirPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(projectDirPath, os.ModePerm)
		if err != nil {
			log.Println("create project directory failed: ", err.Error())
			return err
		}
	} else {
		log.Println("project name already exists")
		return errors.New("project name already exists")
	}
	projectFilePath := filepath.Join(projectDirPath, consts.PROJECT_FILE_NAME+".yml")
	fileContent, err := yaml.Marshal(&project)
	if err != nil {
		log.Println("create project yaml marshal failed: ", err.Error())
		return err
	}
	//write data to yaml file
	err = os.WriteFile(projectFilePath, fileContent, 0777)
	if err != nil {
		log.Println("create project write file failed: ", err.Error())
		return err
	}
	return nil
}

// UpdateProject update project by name
func (p *ProjectService) UpdateProject(oldName string, project *model.Project) error {
	//old project path
	oldProjectPath := filepath.Join(utils.DefaultConfigDir(), consts.PROJECT_DIR_NAME, oldName)
	_, err := os.Stat(oldProjectPath)
	if os.IsNotExist(err) {
		log.Println("update project failed: ", err.Error())
		return err
	}
	//new project path
	newProjectPath := filepath.Join(utils.DefaultConfigDir(), consts.PROJECT_DIR_NAME, project.Name)
	//rename project directory
	err = os.Rename(oldProjectPath, newProjectPath)
	if err != nil {
		log.Println("update project failed: ", err.Error())
		return err
	}
	projectFilePath := filepath.Join(newProjectPath, consts.PROJECT_FILE_NAME+".yml")
	_, err = os.Stat(projectFilePath)
	if os.IsNotExist(err) {
		log.Println("update project file failed: ", err.Error())
		return err
	}
	data, err := yaml.Marshal(project)
	if err != nil {
		log.Println("update project: yaml marshal failed: ", err.Error())
		return err
	}
	err = os.WriteFile(projectFilePath, data, 0777)
	if err != nil {
		log.Println("update project: write file failed: ", err.Error())
		return err
	}
	return nil
}

// GetProjects get project list
func (p *ProjectService) GetProjects(keyword string, page, size int) *model.ProjectPage {
	var projectPage model.ProjectPage
	var projects []model.Project
	//project directory path
	projectDirectory := filepath.Join(utils.DefaultConfigDir(), consts.PROJECT_DIR_NAME)
	_, err := os.Stat(projectDirectory)
	if os.IsNotExist(err) {
		log.Println("get projects: project directory does not exist ", err.Error())
		return &projectPage
	}
	files, err := os.ReadDir(projectDirectory)
	if err != nil {
		log.Println("get projects: read project directory failed ", err.Error())
		return &projectPage
	}
	for _, file := range files {
		var ymlPath string
		if keyword != "" {
			if strings.Contains(file.Name(), keyword) {
				//job yml file path
				ymlPath = filepath.Join(projectDirectory, file.Name(), consts.PROJECT_FILE_NAME+".yml")
			} else {
				continue
			}
		} else {
			ymlPath = filepath.Join(projectDirectory, file.Name(), consts.PROJECT_FILE_NAME+".yml")
		}
		//judge whether the project file exists
		_, err := os.Stat(ymlPath)
		//not exist
		if os.IsNotExist(err) {
			log.Println("project file not exist ", err.Error())
			continue
		}
		fileContent, err := os.ReadFile(ymlPath)
		if err != nil {
			log.Println("get project read file failed ", err.Error())
			continue
		}
		var project model.Project
		//deserialization job yml file
		err = yaml.Unmarshal(fileContent, &project)
		if err != nil {
			log.Println("get projects: deserialization project file failed ", err.Error())
			continue
		}
		projects = append(projects, project)
	}
	pageNum, size, start, end := utils.SlicePage(page, size, len(projects))
	projectPage.Page = pageNum
	projectPage.PageSize = size
	projectPage.Total = len(projects)
	projectPage.Data = projects[start:end]
	return &projectPage
}

// DeleteProject delete project
func (p *ProjectService) DeleteProject(name string) error {
	// project directory path
	src := filepath.Join(utils.DefaultConfigDir(), consts.PROJECT_DIR_NAME, name)
	_, err := os.Stat(src)
	if os.IsNotExist(err) {
		log.Println("delete project failed ", err.Error())
		return err
	}
	err = os.RemoveAll(src)
	if err != nil {
		log.Println("delete project failed ", err.Error())
		return err
	}
	return nil
}
