package parameter

type TemplateCheck struct {
	Name          string `json:"name"`
	RepositoryUrl string `json:"repositoryUrl"`
}

type MetaScanCheck struct {
	Name          string   `json:"name"`
	CheckType     []string `yaml:"checkType"`
	Tool          []string `yaml:"tool"`
	RepositoryUrl string   `json:"repositoryUrl"`
}
