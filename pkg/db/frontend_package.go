package db

import (
	uuid "github.com/iris-contrib/go.uuid"
	"time"
)

type FrontendPackage struct {
	Id               uint      `gorm:"primaryKey" json:"id"`
	ProjectId        uuid.UUID `json:"projectId"`
	WorkflowId       uint      `json:"workflowId"`
	WorkflowDetailId uint      `json:"workflowDetailId"`
	Name             string    `json:"name"`
	Version          string    `json:"version"`
	Branch           string    `json:"branch"`
	Domain           string    `json:"domain"`
	BuildTime        time.Time `json:"buildTime"`
	PackageIdentity  string    `json:"packageIdentity"`
	CreateTime       time.Time `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
}
