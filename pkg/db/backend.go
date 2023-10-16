package db

import (
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	uuid "github.com/iris-contrib/go.uuid"
	"time"
)

type BackendPackage struct {
	Id               uint                    `gorm:"primaryKey" json:"id"`
	ProjectId        uuid.UUID               `json:"projectId"`
	WorkflowId       uint                    `json:"workflowId"`
	WorkflowDetailId uint                    `json:"workflowDetailId"`
	Name             string                  `json:"name"`
	Version          string                  `json:"version"`
	BuildTime        time.Time               `json:"buildTime"`
	Network          string                  `json:"network"`
	AbiInfo          string                  `json:"abiInfo"`
	CreateTime       time.Time               `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
	Type             consts.ProjectFrameType `json:"type"`   // see #consts.ProjectFrameType
	Status           consts.DeployStatus     `json:"status"` // see #consts.
	Branch           string                  `json:"branch"`
	CodeInfo         string                  `json:"codeInfo"`
}

type BackendDeploy struct {
	Id               uint                    `gorm:"primaryKey" json:"id"`
	PackageId        uint                    `json:"packageId"`
	ProjectId        uuid.UUID               `json:"projectId"`
	WorkflowId       uint                    `json:"workflowId"`
	WorkflowDetailId uint                    `json:"workflowDetailId"`
	Version          string                  `json:"version"`
	DeployTime       time.Time               `gorm:"column:deploy_time;default:current_timestamp" json:"deployTime"`
	Network          string                  `json:"network"`
	Address          string                  `json:"address"`
	CreateTime       time.Time               `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
	Type             consts.ProjectFrameType `json:"type"` // see #consts.ProjectFrameType
	DeployTxHash     string                  `json:"deployTxHash"`
	Status           consts.DeployStatus     `json:"status"` // 1: deploying, 2: success , 3: fail
	AbiInfo          string                  `json:"abiInfo"`
	Name             string                  `json:"name"`
}
