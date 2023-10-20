package vo

import (
	uuid "github.com/iris-contrib/go.uuid"
	"time"
)

type ContractDeployInfoVo struct {
	Version      string                    `json:"version"`
	ContractInfo map[string]ContractInfoVo `json:"contractInfo"`
}

type ContractInfoVo struct {
	Id               uint          `json:"id"`
	ProjectId        uuid.UUID     `json:"projectId"`
	WorkflowId       uint          `json:"workflowId"`
	WorkflowDetailId uint          `json:"workflowDetailId"`
	AbiInfo          string        `json:"abiInfo"`
	DeployInfo       []DeployInfVo `json:"deployInfo"`
}

type DeployInfVo struct {
	Network string `json:"network"`
	Address string `json:"address"`
	Name    string `json:"name"`
}

type ContractVo struct {
	Id               uint      `json:"id"`
	ProjectId        uuid.UUID `json:"projectId"`
	WorkflowId       uint      `json:"workflowId"`
	WorkflowDetailId uint      `json:"workflowDetailId"`
	Name             string    `json:"name"`
	Version          string    `json:"version"`
	AbiInfo          string    `json:"abiInfo"`
	ByteCode         string    `json:"byteCode"`
	Type             uint      `json:"type"` // see #consts.ProjectFrameType
	AptosMv          string    `json:"aptosMv"`
	Branch           string    `json:"branch"`
}

type ContractVersionAndCodeInfoVo struct {
	Version  string `json:"version"`
	Type     int    `json:"type"`
	Branch   string `json:"branch"`
	CodeInfo string `json:"codeInfo"`
	Url      string `json:"url"`
}

type ContractDeployVo struct {
	Id            uint      `json:"id"`
	ContractName  string    `json:"contractName"`
	ContractId    uint      `json:"contractId"`
	ProjectId     uuid.UUID `json:"projectId"`
	Version       string    `json:"version"`
	DeployTime    time.Time `json:"deployTime"`
	Network       string    `json:"network"`
	Address       string    `json:"address"`
	CreateTime    time.Time `json:"createTime"`
	Type          uint      `json:"type"` // see #consts.ProjectFrameType
	DeclareTxHash string    `json:"declareTxHash"`
	DeployTxHash  string    `json:"deployTxHash"`
	Status        uint      `json:"status"` // 1: deploying, 2: success , 3: fail
	AbiInfo       string    `json:"abiInfo"`
	Url           string    `json:"url"`
	CodeInfo      string    `json:"codeInfo"`
}
