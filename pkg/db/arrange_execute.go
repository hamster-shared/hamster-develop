package db

import (
	uuid "github.com/iris-contrib/go.uuid"
	"time"
)

type ContractArrangeExecute struct {
	Id                 uint      `gorm:"primaryKey" json:"id"`
	ProjectId          uuid.UUID `json:"projectId"`
	FkArrangeId        string    `json:"fkArrangeId"`
	Version            string    `json:"version"`
	Network            string    `json:"network"`
	ArrangeProcessData string    `json:"arrangeProcessData"`
	CreateTime         time.Time `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
	UpdateTime         time.Time `gorm:"column:update_time;default:current_timestamp" json:"updateTime"`
}
