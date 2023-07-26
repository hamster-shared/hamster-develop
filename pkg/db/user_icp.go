package db

import (
	"gorm.io/gorm"
	"time"
)

type UserIcp struct {
	Id           int            `gorm:"primaryKey" json:"id"`
	FkUserId     int            `json:"fkUserId"`
	IdentityName string         `json:"identityName"`
	AccountId    string         `json:"accountId"`
	PrincipalId  string         `json:"principalId"`
	WalletId     string         `json:"walletId"`
	CreateTime   time.Time      `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
	UpdateTime   time.Time      `json:"updateTime"`
	DeleteTime   gorm.DeletedAt `json:"deleteTime"`
}
