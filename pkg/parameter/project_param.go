package parameter

type CreateProjectParam struct {
	Name          string `json:"name"`
	Type          int    `json:"type"`
	TemplateOwner string `json:"templateOwner"`
	TemplateRepo  string `json:"templateRepo"`
	FrameType     int    `json:"frameType"`
	RepoOwner     string `json:"repoOwner"`
	UserId        int64  `json:"userId"`
	Branch        string `json:"branch"`
}

type CheckNameParam struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
}
