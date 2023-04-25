package vo

import (
	uuid "github.com/iris-contrib/go.uuid"
	"time"
)

type ReportVo struct {
	Id               uint      `json:"id"`
	ProjectId        uuid.UUID `json:"projectId"`
	WorkflowId       uint      `json:"workflowId"`
	WorkflowDetailId uint      `json:"workflowDetailId"`
	Name             string    `json:"name"`
	Type             uint      `json:"type"`
	CheckTool        string    `json:"checkTool"`
	Result           string    `json:"result"`
	CheckTime        time.Time `json:"checkTime"`
	Issues           int       `json:"issues"`
	MetaScanOverview string    `json:"metaScanOverview"`
	ToolType         int       `json:"toolType"`
	ReportFile       string    `json:"reportFile"`
}

type ReportOverView struct {
	Title   string     `json:"title"`
	Content []ReportVo `json:"content"`
}
