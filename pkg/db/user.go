package db

import (
	"time"
)

type User struct {
	Id         uint      `gorm:"primaryKey" json:"id"`
	Username   string    `gorm:"username" json:"username"`
	Token      string    `json:"token"`
	AvatarUrl  string    `json:"avatarUrl"`
	HtmlUrl    string    `json:"htmlUrl"`
	FirstState int       `json:"firstState"`
	UserEmail  string    `json:"userEmail"`
	CreateTime time.Time `gorm:"column:create_time;default:current_timestamp" json:"create_time"`
}
