package db

import (
	uuid "github.com/iris-contrib/go.uuid"
	"time"
)

type BackendPackage struct {
	Id               uint      `gorm:"primaryKey" json:"id"`
	ProjectId        uuid.UUID `json:"projectId"`
	WorkflowId       uint      `json:"workflowId"`
	WorkflowDetailId uint      `json:"workflowDetailId"`
	Name             string    `json:"name"`
	Version          string    `json:"version"`
	BuildTime        time.Time `json:"buildTime"`
	AbiInfo          string    `json:"abiInfo"`
	ByteCode         string    `json:"byteCode"`
	CreateTime       time.Time `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
	Type             uint      `json:"type"`   // see #consts.ProjectFrameType
	Status           uint      `json:"status"` // 1: deploying, 2: success , 3: fail
	Branch           string    `json:"branch"`
}

type BackendDeploy struct {
	Id               uint      `gorm:"primaryKey" json:"id"`
	ContractId       uint      `json:"contractId"`
	ProjectId        uuid.UUID `json:"projectId"`
	Version          string    `json:"version"`
	DeployTime       time.Time `gorm:"column:deploy_time;default:current_timestamp" json:"deployTime"`
	Network          string    `json:"network"`
	Address          string    `json:"address"`
	CreateTime       time.Time `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
	Type             uint      `json:"type"` // see #consts.ProjectFrameType
	DeclareTxHash    string    `json:"declareTxHash"`
	DeployTxHash     string    `json:"deployTxHash"`
	Status           uint      `json:"status"` // 1: deploying, 2: success , 3: fail
	WorkflowId       uint      `json:"workflowId"`
	WorkflowDetailId uint      `json:"workflowDetailId"`
}
