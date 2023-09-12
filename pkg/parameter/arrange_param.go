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
