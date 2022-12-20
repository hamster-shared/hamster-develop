package db

import (
	"database/sql"
	"time"
)

type Project struct {
	Id            uint `gorm:"primaryKey" json:"id"`
	Name          string
	UserId        uint
	Type          uint
	RepositoryUrl string
	FrameType     string
	Creator       uint
	DeleteUser    uint
	UpdateUser    uint
	CreateTime    time.Time    `json:"create_time"`
	UpdateTime    time.Time    `json:"update_time"`
	DeleteTime    sql.NullTime `json:"delete_time"`
}
