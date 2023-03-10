package service

import (
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
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
		containerDeploy.UpdateTime = time.Now()
		err = c.db.Create(&containerDeploy).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ContainerDeployService) QueryDeployParam(projectId string) (db.ContainerDeployParam, error) {
	var containerDeploy db.ContainerDeployParam
	err := c.db.Where("project_id = ?", projectId).First(&containerDeploy).Error
	return containerDeploy, err
}

func (c *ContainerDeployService) CheckDeployParam(projectId string) bool {
	var containerDeploy db.ContainerDeployParam
	err := c.db.Where("project_id = ?", projectId).First(&containerDeploy).Error
	if err != nil {
		return false
	}
	return true
}

func (c *ContainerDeployService) UpdateContainerDeploy(projectId uuid.UUID, data parameter.K8sDeployParam) error {
	var containerDeploy db.ContainerDeployParam
	err := c.db.Where("project_id = ? ", projectId.String()).First(&containerDeploy).Error
	if err != nil {
		containerDeploy.ProjectId = projectId
	}
	containerDeploy.ContainerPort = data.ContainerPort
	containerDeploy.ServicePort = data.ServicePort
	containerDeploy.ServiceProtocol = data.ServiceProtocol
	containerDeploy.ServiceTargetPort = data.ServiceTargetPort
	containerDeploy.UpdateTime = time.Now()
	return c.db.Save(&containerDeploy).Error
}

func (c *ContainerDeployService) GetContainerDeploy(projectId string) db.ContainerDeployParamVo {
	var containerDeploy db.ContainerDeployParam
	var data db.ContainerDeployParamVo
	err := c.db.Where("project_id = ? ", projectId).First(&containerDeploy).Error
	if err != nil {
		return data
	}
	copier.Copy(&data, &containerDeploy)
	return data
}
