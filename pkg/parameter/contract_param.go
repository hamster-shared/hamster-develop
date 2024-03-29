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
	AbiInfo       string `json:"abiInfo"`
}

type ContractDeployIngParam struct {
	ContractId   uint   `json:"contractId"`
	ProjectId    string `json:"projectId"`
	Version      string `json:"version"`
	Network      string `json:"network"`
	DeployTxHash string `json:"deployTxHash"`
	RpcUrl       string `json:"rpcUrl"`
}
