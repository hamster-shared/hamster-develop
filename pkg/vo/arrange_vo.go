package vo

import (
	"database/sql"
	uuid "github.com/iris-contrib/go.uuid"
	"time"
)

type DeployContractListVo struct {
	Id               uint      `json:"id"`
	ContractId       uint      `json:"contractId"`
	ContractName     string    `json:"contractName"`
	ProjectId        uuid.UUID `json:"projectId"`
	Version          string    `json:"version"`
	DeployTime       time.Time `json:"deployTime"`
	DeployTimeFormat string    `json:"deployTimeFormat"`
	Network          string    `json:"network"`
	Address          string    `json:"address"`
	Type             uint      `json:"type"` // see #consts.ProjectFrameType
	DeclareTxHash    string    `json:"declareTxHash"`
	DeployTxHash     string    `json:"deployTxHash"`
	Status           uint      `json:"status"` // 1: deploying, 2: success , 3: fail
	AbiInfo          string    `json:"abiInfo"`
}

type ContractArrangeCacheVo struct {
	Id              uint      `json:"id"`
	ProjectId       uuid.UUID `json:"projectId"`
	ContractId      uint      `json:"contractId"`
	ContractName    string    `json:"contractName"`
	Version         string    `json:"version"`
	OriginalArrange string    `json:"originalArrange"`
	CreateTime      time.Time `json:"createTime"`
	UpdateTime      time.Time `json:"updateTime"`
}

type ToBeArrangedContractListVo struct {
	UseContract   []ToBeArrangedContractVo `json:"useContract"`
	NoUseContract []ToBeArrangedContractVo `json:"noUseContract"`
}

type ToBeArrangedContractVo struct {
	Id              uint           `gorm:"primaryKey" json:"id"`
	ProjectId       uuid.UUID      `json:"projectId"`
	Name            string         `json:"name"`
	Version         string         `json:"version"`
	Network         sql.NullString `json:"network"`
	BuildTime       time.Time      `json:"buildTime"`
	AbiInfo         string         `json:"abiInfo"`
	ByteCode        string         `json:"byteCode"`
	Branch          string         `json:"branch"`
	CodeInfo        string         `json:"codeInfo"`
	OriginalArrange string         `json:"originalArrange"`
}
