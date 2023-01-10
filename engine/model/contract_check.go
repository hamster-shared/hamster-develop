package model

type ContractCheckResult struct {
	Name    string
	Result  string
	Tool    string
	Context []ContractCheckResultDetails
}

func NewContractCheckResult(name string, result string, tool string, context []ContractCheckResultDetails) ContractCheckResult {
	return ContractCheckResult{
		Name:    name,
		Result:  result,
		Tool:    tool,
		Context: context,
	}
}

type ContractCheckResultDetails struct {
	Name    string
	Message string
}

func NewContractCheckResultDetails(name string, message string) ContractCheckResultDetails {
	return ContractCheckResultDetails{
		Name:    name,
		Message: message,
	}
}
