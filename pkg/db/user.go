package db

import "database/sql"

type User struct {
	Id         uint         `gorm:"primaryKey" json:"id"`
	Username   string       `gorm:"username" json:"username"`
	Token      string       `json:"token"`
	CreateTime sql.NullTime `gorm:"create_time" json:"create_time"`
}
