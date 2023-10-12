package db

import (
	uuid "github.com/iris-contrib/go.uuid"
	"time"
)

type ContractArrange struct {
	Id              uint      `gorm:"primaryKey" json:"id"`
	ProjectId       uuid.UUID `json:"projectId"`
	Version         string    `json:"version"`
	OriginalArrange string    `json:"originalArrange"`
	CreateTime      time.Time `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
	UpdateTime      time.Time `gorm:"column:update_time;default:current_timestamp" json:"updateTime"`
}
