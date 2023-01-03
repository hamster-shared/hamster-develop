package db

import "time"

type Contract struct {
	Id               uint      `gorm:"primaryKey" json:"id"`
	ProjectId        uint      `json:"projectId"`
	WorkflowId       uint      `json:"workflowId"`
	WorkflowDetailId uint      `json:"workflowDetailId"`
	Name             string    `json:"name"`
	Version          string    `json:"version"`
	Network          string    `json:"network"`
	BuildTime        time.Time `json:"buildTime"`
	AbiInfo          string    `json:"abiInfo"`
	ByteCode         string    `json:"byteCode"`
	CreateTime       time.Time `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
}

type ContractDeploy struct {
	Id         uint      `gorm:"primaryKey" json:"id"`
	ContractId uint      `json:"contractId"`
	ProjectId  uint      `json:"projectId"`
	Version    string    `json:"version"`
	DeployTime time.Time `gorm:"column:deploy_time;default:current_timestamp" json:"deployTime"`
	Network    string    `json:"network"`
	Address    string    `json:"address"`
	CreateTime time.Time `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
}
