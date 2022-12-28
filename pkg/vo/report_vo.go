package vo

import "time"

type ReportVo struct {
	Id               uint      `json:"id"`
	ProjectId        uint      `json:"projectId"`
	WorkflowId       uint      `json:"workflowId"`
	WorkflowDetailId uint      `json:"workflowDetailId"`
	Name             string    `json:"name"`
	Type             uint      `json:"type"`
	CheckTool        string    `json:"checkTool"`
	Result           string    `json:"result"`
	CheckTime        time.Time `json:"checkTime"`
	ReportFile       string    `json:"reportFile"`
}
