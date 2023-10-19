package parameter

type ContractArrangeParam struct {
	ProjectId       string `json:"projectId"  binding:"required"`
	Version         string `json:"version"  binding:"required"`
	OriginalArrange string `json:"originalArrange"  binding:"required"`
}

type ContractArrangeExecuteParam struct {
	ProjectId          string `json:"projectId"  binding:"required"`
	FkArrangeId        string `json:"fkArrangeId"  binding:"required"`
	Version            string `json:"version"  binding:"required"`
	Network            string `json:"network"  binding:"required"`
	ArrangeProcessData string `json:"arrangeProcessData"  binding:"required"`
}

type ContractArrangeExecuteUpdateParam struct {
	Id                 uint   `json:"id"  binding:"required"`
	ArrangeProcessData string `json:"arrangeProcessData"  binding:"required"`
}

type ContractArrangeCacheParam struct {
	ProjectId       string `json:"projectId" binding:"required"`
	ContractId      uint   `json:"contractId" binding:"required"`
	ContractName    string `json:"contractName" binding:"required"`
	Version         string `json:"version" binding:"required"`
	OriginalArrange string `json:"originalArrange" binding:"required"`
}

type ContractArrangeCacheQuery struct {
	ProjectId    string `json:"projectId" from:"projectId"  binding:"required"`
	ContractId   uint   `json:"contractId" from:"contractId"  binding:"required"`
	ContractName string `json:"contractName" from:"contractName"  binding:"required"`
	Version      string `json:"version" from:"version"  binding:"required"`
}

type ContractNameArrangeParam struct {
	ProjectId           string              `json:"projectId"  binding:"required"`
	Version             string              `json:"version"  binding:"required"`
	ContractNameArrange ContractNameArrange `json:"contractNameArrange"  binding:"required"`
}

type ContractNameArrange struct {
	UseContract   []string `json:"useContract"`
	NoUseContract []string `json:"noUseContract"`
}
