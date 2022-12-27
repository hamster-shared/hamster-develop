package parameter

import "github.com/hamster-shared/a-line/pkg/consts"

type SaveWorkflowParam struct {
	ProjectId  uint                `json:"projectId"`
	Type       consts.WorkflowType `json:"type"`
	ExecFile   string              `json:"execFile"`
	LastExecId uint                `json:"lastExecId"`
}
