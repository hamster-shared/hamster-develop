package parameter

type ContractDeployParam struct {
	ContractId    int    `json:"contractId"`
	ProjectId     string `json:"projectId"`
	Version       string `json:"version"`
	Network       string `json:"network"`
	Address       string `json:"address"`
	DeclareTxHash string `json:"declareTxHash"`
	DeployTxHash  string `json:"deployTxHash"`
	Status        uint   `json:"status"` // 1: deploying, 2: success , 3: fail
}
