package vo

import (
	uuid "github.com/iris-contrib/go.uuid"
	"time"
)

type WorkflowPage struct {
	Data     []WorkflowVo `json:"data"`
	Total    int          `json:"total"`
	Page     int          `json:"page"`
	PageSize int          `json:"pageSize"`
}

type WorkflowVo struct {
	Id          uint      `json:"id"`
	ProjectId   uuid.UUID `json:"projectId"`
	DetailId    uint      `json:"detailId"`
	Type        uint      `json:"type"`
	LastExecId  uint      `json:"lastExecId"`
	ExecNumber  uint      `json:"execNumber"`
	StageInfo   string    `json:"stageInfo"`
	CodeInfo    string    `json:"codeInfo"`
	TriggerUser string    `json:"triggerUser"`
	TriggerMode uint      `json:"triggerMode"`
	Status      uint      `json:"status"`
	StartTime   time.Time `json:"startTime"`
	Duration    int64     `json:"duration"`
}

type WorkflowDetailVo struct {
	Id          uint      `json:"id"`
	WorkflowId  uint      `json:"workflowId"`
	StageInfo   string    `json:"stageInfo"`
	Status      uint      `json:"status"`
	ExecNumber  uint      `json:"execNumber"`
	StartTime   time.Time `json:"startTime"`
	Duration    int64     `json:"duration"`
	TriggerUser string    `json:"triggerUser"`
}

type DeployResultVo struct {
	WorkflowId uint `json:"workflowId"`
	DetailId   uint `json:"detailId"`
}

const WORKFLOW_STATUS_NOT_RUN uint = 0
const WORKFLOW_STATUS_RUNNING uint = 1
const WORKFLOW_STATUS_SUCCESS uint = 2
const WORKFLOW_STATUS_FAIL uint = 3
const WORKFLOW_STATUS_CANCEL uint = 4
