package parameters

type ProjectCreateVo struct {
	Name        string `json:"name"`
	Type        uint   `json:"type"`
	TemplateUrl string `json:"templateUrl"`
	FrameType   uint   `json:"frameType"`
}
