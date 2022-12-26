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
	ExecNumber  uint      `json:"execNumber"`
	StageInfo   string    `json:"stageInfo"`
	CodeInfo    string    `json:"codeInfo"`
	TriggerUser string    `json:"triggerUser"`
	TriggerMode uint      `json:"triggerMode"`
	Status      uint      `json:"status"`
	StartTime   time.Time `json:"startTime"`
}

type WorkflowDetailVo struct {
	Id         uint `gorm:"primaryKey" json:"id"`
	WorkflowId uint
	StageInfo  string
	Status     uint
	StartTime  time.Time
	EndTime    time.Time
}
