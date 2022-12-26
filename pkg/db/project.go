package db

import (
	"gorm.io/gorm"
	"time"
)

type Project struct {
	Id            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `json:"name"`
	UserId        uint           `json:"UserId"`
	Type          uint           `json:"type"`
	RepositoryUrl string         `json:"RepositoryUrl"`
	FrameType     string         `json:"frameType"`
	Creator       uint           `json:"creator"`
	DeleteUser    uint           `json:"deleteUser"`
	UpdateUser    uint           `json:"updateUser"`
	CreateTime    time.Time      `json:"createTime"`
	UpdateTime    time.Time      `json:"updateTime"`
	DeleteTime    gorm.DeletedAt `json:"deleteTime"`
}
