package db

import (
	uuid "github.com/iris-contrib/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Project struct {
	Id            uuid.UUID      `gorm:"primaryKey" json:"id"`
	Name          string         `json:"name"`
	UserId        int64          `json:"UserId"`
	Type          uint           `json:"type"`
	RepositoryUrl string         `json:"repositoryUrl"`
	FrameType     int            `json:"frameType"`
	Creator       int64          `json:"creator"`
	DeleteUser    uint           `json:"deleteUser"`
	UpdateUser    uint           `json:"updateUser"`
	Branch        string         `json:"branch"`
	CreateTime    time.Time      `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
	UpdateTime    time.Time      `json:"updateTime"`
	DeleteTime    gorm.DeletedAt `json:"deleteTime"`
}

func (p *Project) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := uuid.NewV4()
	p.Id = id
	return
}
