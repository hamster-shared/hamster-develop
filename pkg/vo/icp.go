package vo

import db2 "github.com/hamster-shared/hamster-develop/pkg/db"

type UserIcpInfoVo struct {
	UserId     int    `json:"userId"`
	AccountId  string `json:"accountId"`
	IcpBalance string `json:"icpBalance"`
}

type IcpCanisterPage struct {
	Data     []db2.IcpCanister `json:"data"`
	Total    int               `json:"total"`
	Page     int               `json:"page"`
	PageSize int               `json:"pageSize"`
}
