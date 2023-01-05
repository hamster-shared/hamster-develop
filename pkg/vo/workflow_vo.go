package vo

import (
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
	ProjectId   uint      `json:"projectId"`
	Type        uint      `json:"type"`
	LastExecId  uint      `json:"lastExecId"`
	ExecNumber  uint      `json:"execNumber"`
	StageInfo   string    `json:"stageInfo"`
	CodeInfo    string    `json:"codeInfo"`
	TriggerUser string    `json:"triggerUser"`
	TriggerMode uint      `json:"triggerMode"`
	Status      uint      `json:"status"`
	StartTime   time.Time `json:"startTime"`
}

type WorkflowDetailVo struct {
	Id         uint      `json:"id"`
	WorkflowId uint      `json:"workflowId"`
	StageInfo  string    `json:"stageInfo"`
	Status     uint      `json:"status"`
	StartTime  time.Time `json:"startTime"`
	Duration   int64     `json:"duration"`
}
