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

type CreateByCodeParam struct {
	Name      string `json:"name"`
	Type      int    `json:"type"`
	FrameType int    `json:"frameType"`
	FileName  string `json:"fileName"`
	Content   string `json:"content"`
}

type CheckNameParam struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
}
