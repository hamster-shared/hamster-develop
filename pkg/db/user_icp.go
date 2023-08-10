package db

import (
	"database/sql"
	"gorm.io/gorm"
)

type UserIcp struct {
	Id           int            `gorm:"primaryKey" json:"id"`
	FkUserId     uint           `json:"fkUserId"`
	IdentityName string         `json:"identityName"`
	AccountId    string         `json:"accountId"`
	PrincipalId  string         `json:"principalId"`
	WalletId     string         `json:"walletId"`
	CreateTime   sql.NullTime   `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
	UpdateTime   sql.NullTime   `json:"updateTime"`
	DeleteTime   gorm.DeletedAt `json:"deleteTime"`
}
