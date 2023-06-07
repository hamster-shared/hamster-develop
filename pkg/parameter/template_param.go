package parameter

type TemplateCheck struct {
	Name          string `json:"name"`
	RepositoryUrl string `json:"repositoryUrl"`
}

type MetaScanCheck struct {
	Name          string   `json:"name"`
	CheckType     []string `yaml:"checkType"`
	Tool          []string `yaml:"tool"`
	ToolTitle     []string `yaml:"toolTitle"`
	OutNeed       string   `yaml:"outNeed"`
	RepositoryUrl string   `json:"repositoryUrl"`
	InstallTool   []string `json:"installTool"yaml:"installTool"`
}
