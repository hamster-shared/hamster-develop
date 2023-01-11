package parameter

import (
	"github.com/hamster-shared/a-line/pkg/consts"
	uuid "github.com/iris-contrib/go.uuid"
)

type SaveWorkflowParam struct {
	ProjectId  uuid.UUID           `json:"projectId"`
	Type       consts.WorkflowType `json:"type"`
	ExecFile   string              `json:"execFile"`
	LastExecId uint                `json:"lastExecId"`
}
