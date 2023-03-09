package db

import (
	uuid "github.com/iris-contrib/go.uuid"
	"gorm.io/gorm"
	"time"
)

type ContainerDeployParam struct {
	Id                int            `gorm:"primaryKey" json:"id"`
	ProjectId         uuid.UUID      `json:"projectId"`
	WorkflowId        int            `json:"workflowId"`
	ContainerPort     int32          `json:"containerPort"`
	ServiceProtocol   string         `json:"serviceProtocol"`
	ServicePort       int32          `json:"servicePort"`
	ServiceTargetPort int32          `json:"serviceTargetPort"`
	CreateTime        time.Time      `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
	UpdateTime        time.Time      `json:"updateTime"`
	DeleteTime        gorm.DeletedAt `json:"deleteTime"`
}
