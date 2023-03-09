package parameter

type CreateProjectParam struct {
	Name          string `json:"name"`
	Type          int    `json:"type"`
	DeployType    int    `json:"deployType"`
	TemplateOwner string `json:"templateOwner"`
	TemplateUrl   string `json:"templateUrl"`
	TemplateRepo  string `json:"templateRepo"`
	FrameType     int    `json:"frameType"`
	RepoOwner     string `json:"repoOwner"`
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

type K8sDeployParam struct {
	ContainerPort     int32  `json:"containerPort"`
	ServiceProtocol   string `json:"serviceProtocol"`
	ServicePort       int32  `json:"servicePort"`
	ServiceTargetPort int32  `json:"serviceTargetPort"`
}
