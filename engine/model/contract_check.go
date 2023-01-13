package model

import "encoding/json"

type ContractCheckResult[T ResultDetailType] struct {
	Name    string                          `json:"name"`
	Result  string                          `json:"result"`
	Tool    string                          `json:"tool"`
	Context []ContractCheckResultDetails[T] `json:"context"`
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
	string | []ContractStyleGuideValidationsReportDetails | []ContractMethodsPropertiesReportDetails | json.RawMessage
}

type ContractCheckResultDetails[T ResultDetailType] struct {
	Name    string `json:"name"`
	Issue   int    `json:"issue"`
	Message T      `json:"message"`
}

func NewContractCheckResultDetails[T ResultDetailType](name string, issue int, message T) ContractCheckResultDetails[T] {
	return ContractCheckResultDetails[T]{
		Name:    name,
		Issue:   issue,
		Message: message,
	}
}

type ContractStyleGuideValidationsReportDetails struct {
	Line         string `json:"line"`
	Column       string `json:"column"`
	Level        string `json:"level"`
	OriginalText string `json:"originalText"`
	Note         string `json:"note"`
	Tool         string `json:"tool"`
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

type ContractMethodsPropertiesReportDetails struct {
	Contract   string `json:"contract"`
	Category   string `json:"category"`
	Function   string `json:"function"`
	Visibility string `json:"visibility"`
	ViewPure   string `json:"view_pure"`
	Returns    string `json:"returns"`
	Modifiers  string `json:"modifiers"`
}

func NewContractMethodsPropertiesReportDetails(contract, category, function, visibility, viewPure, returns, modifiers string) ContractMethodsPropertiesReportDetails {
	return ContractMethodsPropertiesReportDetails{
		Contract:   contract,
		Category:   category,
		Function:   function,
		Visibility: visibility,
		ViewPure:   viewPure,
		Returns:    returns,
		Modifiers:  modifiers,
	}
}
