package db

import "time"

type Report struct {
	Id               uint `gorm:"primaryKey" json:"id"`
	ProjectId        uint
	WorkflowId       uint
	WorkflowDetailId uint
	Name             string
	Type             uint
	CheckTool        string
	Result           string
	CheckTime        time.Time
	ReportFile       string
	CreateTime       time.Time
}
