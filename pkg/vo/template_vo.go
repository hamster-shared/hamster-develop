package vo

type TemplateTypeVo struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TemplateVo struct {
	Id             uint   `json:"id"`
	TemplateTypeId uint   `json:"templateTypeId"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Audited        bool   `json:"audited"`
	LastVersion    string `json:"lastVersion"`
	Logo           string `json:"logo"`
	Image          string `json:"image"`
}

type TemplateDetailVo struct {
	Id               uint   `json:"id"`
	TemplateId       string `json:"templateId"`
	Name             string `json:"name"`
	Audited          bool   `json:"audited"`
	Extensions       string `json:"extensions"`
	Description      string `json:"description"`
	Examples         string `json:"examples"`
	Resources        string `json:"resources"`
	AbiInfo          string `json:"abiInfo"`
	Author           string `json:"author"`
	RepositoryName   string `json:"repositoryName"`
	RepositoryUrl    string `json:"repositoryUrl"`
	Version          string `json:"version"`
	Branch           string `json:"branch"`
	CodeSources      string `json:"codeSources"`
	Title            string `json:"title"`
	TitleDescription string `json:"titleDescription"`
}
