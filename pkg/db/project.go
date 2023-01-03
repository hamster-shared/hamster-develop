package db

import (
	"gorm.io/gorm"
	"time"
)

type Project struct {
	Id            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `json:"name"`
	UserId        int64          `json:"UserId"`
	Type          uint           `json:"type"`
	RepositoryUrl string         `json:"RepositoryUrl"`
	FrameType     int            `json:"frameType"`
	Creator       int64          `json:"creator"`
	DeleteUser    uint           `json:"deleteUser"`
	UpdateUser    uint           `json:"updateUser"`
	CreateTime    time.Time      `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
	UpdateTime    time.Time      `json:"updateTime"`
	DeleteTime    gorm.DeletedAt `json:"deleteTime"`
}
