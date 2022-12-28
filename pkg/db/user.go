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
	CreateTime time.Time `gorm:"create_time" json:"create_time"`
}
