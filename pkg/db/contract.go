package db

import "time"

type Contract struct {
	Id               uint `gorm:"primaryKey" json:"id"`
	ProjectId        uint
	WorkflowId       uint
	WorkflowDetailId uint
	Name             string
	Version          string
	Network          string
	BuildTime        time.Time
	AbiInfo          string
	ByteCode         string
	CreateTime       time.Time
}

type ContractDeploy struct {
	Id         uint `gorm:"primaryKey" json:"id"`
	ContractId uint
	ProjectId  uint
	Version    string
	DeployTime time.Time
	Network    string
	Address    string
	CreateTime time.Time
}
