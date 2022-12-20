package db

import (
	"database/sql"
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
	WorkflowId  uint
	ExecNumber  uint
	StageInfo   string
	TriggerUser string
	TriggerMode uint
	CodeInfo    string
	Status      uint
	StartTime   sql.NullTime
	EndTime     sql.NullTime
	CreateTime  time.Time    `json:"create_time"`
	UpdateTime  time.Time    `json:"update_time"`
	DeleteTime  sql.NullTime `json:"delete_time"`
}
