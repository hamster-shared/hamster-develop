package model

type ContractCheckResult[T ResultDetailType] struct {
	Name    string
	Result  string
	Tool    string
	Context []ContractCheckResultDetails[T]
}

func NewContractCheckResult[T ResultDetailType](name string, result string, tool string, context []ContractCheckResultDetails[T]) ContractCheckResult[T] {
	return ContractCheckResult[T]{
		Name:    name,
		Result:  result,
		Tool:    tool,
		Context: context,
	}
}

type ResultDetailType interface {
	string | []ContractStyleGuideValidationsReportDetails
}

type ContractCheckResultDetails[T ResultDetailType] struct {
	Name    string
	Issue   int
	Message T
}

func NewContractCheckResultDetails[T ResultDetailType](name string, issue int, message T) ContractCheckResultDetails[T] {
	return ContractCheckResultDetails[T]{
		Name:    name,
		Issue:   issue,
		Message: message,
	}
}

type ContractStyleGuideValidationsReportDetails struct {
	Line         string
	Column       string
	Level        string
	OriginalText string
	Note         string
	Tool         string
}

func NewContractStyleGuideValidationsReportDetails(line, column, level, originalText, note, tool string) ContractStyleGuideValidationsReportDetails {
	return ContractStyleGuideValidationsReportDetails{
		Line:         line,
		Column:       column,
		Level:        level,
		OriginalText: originalText,
		Note:         note,
		Tool:         tool,
	}
}
