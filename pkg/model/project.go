package model

type Project struct {
	Name string `yaml:"name,omitempty"json:"name"`
}

type ProjectPage struct {
	Data     []Project `json:"data"`
	Total    int       `json:"total"`
	Page     int       `json:"page"`
	PageSize int       `json:"pageSize"`
}
