package vo

type ContractDeployInfoVo struct {
	Version      string                    `json:"version"`
	ContractInfo map[string]ContractInfoVo `json:"contractInfo"`
}

type ContractInfoVo struct {
	Id               uint          `json:"id"`
	ProjectId        uint          `json:"projectId"`
	WorkflowId       uint          `json:"workflowId"`
	WorkflowDetailId uint          `json:"workflowDetailId"`
	AbiInfo          string        `json:"abiInfo"`
	DeployInfo       []DeployInfVo `json:"deployInfo"`
}

type DeployInfVo struct {
	Network string `json:"network"`
	Address string `json:"address"`
}

type ContractVo struct {
	Id               uint   `json:"id"`
	ProjectId        uint   `json:"projectId"`
	WorkflowId       uint   `json:"workflowId"`
	WorkflowDetailId uint   `json:"workflowDetailId"`
	Name             string `json:"name"`
	Version          string `json:"version"`
	AbiInfo          string `json:"abiInfo"`
	ByteCode         string `json:"byteCode"`
}
