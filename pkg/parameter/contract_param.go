package parameter

type ContractDeployParam struct {
	ContractId int    `json:"contractId"`
	ProjectId  string `json:"projectId"`
	Version    string `json:"version"`
	Network    string `json:"network"`
	Address    string `json:"address"`
}
