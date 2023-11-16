package db

import (
	"time"
)

type ChainNetwork struct {
	Id               uint      `gorm:"primaryKey" json:"id"`
	Logo             string    `json:"logo"`
	Category         string    `json:"category"`
	ChainId          string    `json:"chainId"`
	ChainName        string    `json:"chainName"`
	RpcUrl           string    `json:"rpcUrl"`
	Symbol           string    `json:"symbol"`
	BlockExplorerUrl string    `json:"blockExplorerUrl"`
	Decimals         int       `json:"decimals"`
	CreateTime       time.Time `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
	UpdateTime       time.Time `gorm:"column:update_time;default:current_timestamp" json:"updateTime"`
}
