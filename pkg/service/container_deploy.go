package service

import (
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	uuid "github.com/iris-contrib/go.uuid"
	"gorm.io/gorm"
)

type ContainerDeployService struct {
	db *gorm.DB
}

func NewContainerDeployService() *ContainerDeployService {
	return &ContainerDeployService{
		db: application.GetBean[*gorm.DB]("db"),
	}
}

func (c *ContainerDeployService) SaveDeployParam(projectId uuid.UUID, workflowId int, data parameter.K8sDeployParam) error {
	var containerDeploy db.ContainerDeployParam
	err := c.db.Where("project_id = ? and workflow_id = ? ", projectId, workflowId).First(&containerDeploy).Error
	if err == gorm.ErrRecordNotFound {
		containerDeploy.ContainerPort = data.ContainerPort
		containerDeploy.ServicePort = data.ServicePort
		containerDeploy.ServiceProtocol = data.ServiceProtocol
		containerDeploy.ServiceTargetPort = data.ServiceTargetPort
		containerDeploy.ProjectId = projectId
		containerDeploy.WorkflowId = workflowId
		err = c.db.Create(&containerDeploy).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ContainerDeployService) QueryDeployParam(projectId string, workflowId int) (db.ContainerDeployParam, error) {
	var containerDeploy db.ContainerDeployParam
	err := c.db.Where("project_id = ? and workflow_id = ? ", projectId, workflowId).First(&containerDeploy).Error
	return containerDeploy, err
}

func (c *ContainerDeployService) CheckDeployParam(projectId string, workflowId int) bool {
	var containerDeploy db.ContainerDeployParam
	err := c.db.Where("project_id = ? and workflow_id = ? ", projectId, workflowId).First(&containerDeploy).Error
	if err != nil {
		return false
	}
	return true
}
