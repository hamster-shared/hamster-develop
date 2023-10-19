package vo

import (
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
