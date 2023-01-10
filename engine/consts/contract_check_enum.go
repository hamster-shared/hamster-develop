package consts

type ContractCheckEnum struct {
	Name   string
	Result string
	Tool   string
}

func contractCheckResult(name string, tool string) ContractCheckEnum {
	return ContractCheckEnum{
		Name: name,
		Tool: tool,
	}
}

var (
	SolHint     = contractCheckResult("Analysis of contract specifications", "Solhint")
	SolProfiler = contractCheckResult("List all properties of contract methods", "sol-profiler")
	MythRil     = contractCheckResult("Security analysis tool for EVM bytecodes", "mythril")
)

type ContractCheckResultDetails struct {
	Result  string
	message string
}

func contractCheckResultDetails(result string, message string) ContractCheckResultDetails {
	return ContractCheckResultDetails{
		Result:  result,
		message: message,
	}
}

var (
	CheckSuccess = contractCheckResultDetails("Success", "检查成功")
	CheckFail    = contractCheckResultDetails("Fail", "检查失败")
)
