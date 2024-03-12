package db

import (
	uuid "github.com/iris-contrib/go.uuid"
	"time"
)

type Report struct {
	Id               uint      `gorm:"primaryKey" json:"id"`
	ProjectId        uuid.UUID `json:"projectId"`
	WorkflowId       uint      `json:"workflowId"`
	WorkflowDetailId uint      `json:"workflowDetailId"`
	Name             string    `json:"name"`
	Type             uint      `json:"type"`
	ToolType         int       `json:"toolType"`
	CheckTool        string    `json:"checkTool"`
	CheckVersion     string    `json:"checkVersion"`
	Result           string    `json:"result"`
	CheckTime        time.Time `json:"checkTime"`
	ReportFile       string    `json:"reportFile"`
	Issues           int       `json:"issues"`
	MetaScanOverview string    `json:"metaScanOverview"`
	CreateTime       time.Time `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
	Branch           string    `json:"branch"`
	CommitId         string    `json:"commitId"`
	CommitInfo       string    `json:"commitInfo"`
}
