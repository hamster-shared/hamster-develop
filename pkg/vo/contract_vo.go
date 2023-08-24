package vo

import uuid "github.com/iris-contrib/go.uuid"

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
}
