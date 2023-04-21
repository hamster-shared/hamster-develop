package parameter

type TemplateCheck struct {
	Name          string `json:"name"`
	RepositoryUrl string `json:"repositoryUrl"`
}

type MetaScanCheck struct {
	Name string `json:"name"`
	Tool string `json:"tool"`
}
