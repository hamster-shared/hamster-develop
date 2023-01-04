package db

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Workflow struct {
	Id         uint `gorm:"primaryKey" json:"id"`
	ProjectId  uint
	Type       uint
	ExecFile   string
	LastExecId uint
	CreateTime time.Time    `json:"create_time"`
	UpdateTime time.Time    `json:"update_time"`
	DeleteTime sql.NullTime `json:"delete_time"`
}

type WorkflowDetail struct {
	Id          uint `gorm:"primaryKey" json:"id"`
	ProjectId   uint
	Type        uint
	WorkflowId  uint
	ExecNumber  uint
	StageInfo   string
	TriggerUser string
	TriggerMode uint
	CodeInfo    string
	Status      uint
	StartTime   time.Time
	EndTime     time.Time
	CreateTime  time.Time      `gorm:"column:create_time;default:current_timestamp" json:"create_time"`
	UpdateTime  time.Time      `json:"update_time"`
	DeleteTime  gorm.DeletedAt `json:"delete_time"`
}
