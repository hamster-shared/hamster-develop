package db

import (
	"gorm.io/gorm"
	"time"
)

type RequestLog struct {
	Id     uint   `gorm:"primaryKey" json:"id"`
	Url    string `gorm: "url" json:"url"`
	Method string `gorm: "method" json:"method"`
	Token  string `gorm: "token" json:"token""`

	CreateTime time.Time      `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
	UpdateTime time.Time      `json:"updateTime"`
	DeleteTime gorm.DeletedAt `json:"deleteTime"`
}
