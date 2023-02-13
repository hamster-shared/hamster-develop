package db

import (
	"database/sql"
	uuid "github.com/iris-contrib/go.uuid"
	"gorm.io/gorm"
	"time"
)

type FrontendDeploy struct {
	Id               uint           `gorm:"primaryKey" json:"id"`
	PackageId        uint           `json:"packageId"`
	ProjectId        uuid.UUID      `json:"projectId"`
	WorkflowId       uint           `json:"workflowId"`
	WorkflowDetailId uint           `json:"workflowDetailId"`
	Name             string         `json:"name"`
	Image            string         `json:"image"`
	Version          string         `json:"version"`
	Branch           string         `json:"branch"`
	Domain           string         `json:"domain"`
	DeployInfo       string         `json:"deployInfo"`
	DeployTime       sql.NullTime   `json:"deployTime"`
	CreateTime       time.Time      `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
	DeleteTime       gorm.DeletedAt `gorm:"index;column:delete_time;" json:"deleteTime"`
}
