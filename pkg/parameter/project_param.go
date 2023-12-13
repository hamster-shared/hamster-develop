package parameter

type CreateProjectParam struct {
	Name            string `json:"name"`
	Type            int    `json:"type"`
	DeployType      int    `json:"deployType"`
	LabelDisplay    string `json:"labelDisplay"`
	TemplateOwner   string `json:"templateOwner"`
	TemplateUrl     string `json:"templateUrl"`
	TemplateRepo    string `json:"templateRepo"`
	FrameType       uint   `json:"frameType"`
	RepoOwner       string `json:"repoOwner"`
	GistId          string `json:"gistId"`
	DefaultFile     string `json:"defaultFile"`
	Branch          string `json:"branch"`
	EvmTemplateType uint   `json:"evmTemplateType"`
}

type CreateProjectParamV2 struct {
	Name            string `json:"name"`
	Type            int    `json:"type"`
	DeployType      int    `json:"deployType"`
	LabelDisplay    string `json:"labelDisplay"`
	TemplateOwner   string `json:"templateOwner"`
	TemplateUrl     string `json:"templateUrl"`
	TemplateRepo    string `json:"templateRepo"`
	FrameType       uint   `json:"frameType"`
	RepoOwner       string `json:"repoOwner"`
	GistId          string `json:"gistId"`
	DefaultFile     string `json:"defaultFile"`
	Branch          string `json:"branch"`
	EvmTemplateType uint   `json:"evmTemplateType"`
	GithubName      string `json:"githubName"`
}

type ImportProjectParam struct {
	Name      string `json:"name"`
	Ecosystem int    `json:"ecosystem"`
	CloneURL  string `json:"cloneUrl"`
	Type      int    `json:"type"`
	InstallId int64  `json:"installId"`
	// only used for frontend
	DeployType int `json:"deployType"`
}

type CreateByCodeParam struct {
	Name      string `json:"name"`
	Type      int    `json:"type"`
	FrameType uint   `json:"frameType"`
	FileName  string `json:"fileName"`
	Content   string `json:"content"`
}

type CheckNameParam struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type K8sDeployParam struct {
	ContainerPort     int32  `json:"containerPort"`
	ServiceProtocol   string `json:"serviceProtocol"`
	ServicePort       int32  `json:"servicePort"`
	ServiceTargetPort int32  `json:"serviceTargetPort"`
}
