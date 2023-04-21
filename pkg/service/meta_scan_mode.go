package service

type SecurityAnalyzerResponse struct {
	Success     bool                     `json:"success"`
	Error       string                   `json:"error"`
	Results     []SecurityAnalyzerResult `json:"results"`
	FileMapping map[string]string        `json:"file_mapping"`
}

type SecurityAnalyzerResult struct {
	Mwe           MalwareWorkflowEngine `json:"mwe"`
	ShowTitle     string                `json:"show_title"`
	AffectedFiles []AffectedFile        `json:"affected_files"`
}

type AffectedFile struct {
	Filepath         string `json:"filepath"`
	FilepathRelative string `json:"filepath_relative"`
	Hightlights      []int  `json:"hightlights"`
	LineStart        int    `json:"line_start"`
	LineEnd          int    `json:"line_end"`
}

type MalwareWorkflowEngine struct {
	Id             string `json:"id"`
	Code           string `json:"code"`
	Severity       string `json:"severity"`
	Title          string `json:"title"`
	Confidence     string `json:"confidence"`
	Description    string `json:"description"`
	Recommendation string `json:"recommendation"`
}
