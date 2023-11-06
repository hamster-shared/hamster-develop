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
	ProjectId    string `json:"projectId" form:"projectId"`
	ContractId   uint   `json:"contractId" form:"contractId" `
	ContractName string `json:"contractName" form:"contractName"`
	Version      string `json:"version" form:"version"`
}

type ContractNameArrangeParam struct {
	ProjectId     string   `json:"projectId"  binding:"required"`
	Version       string   `json:"version"  binding:"required"`
	UseContract   []string `json:"useContract"`
	NoUseContract []string `json:"noUseContract"`
}

type ContractNameArrange struct {
	UseContract   []string `json:"useContract"`
	NoUseContract []string `json:"noUseContract"`
}

type ContractInfoQuery struct {
	Id   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name" `
}
