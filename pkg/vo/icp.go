package vo

type UserIcpInfoVo struct {
	UserId     int    `json:"userId"`
	AccountId  string `json:"accountId"`
	IcpBalance string `json:"icpBalance"`
}

type UserIcpWalletVo struct {
	UserId        int    `json:"userId"`
	WalletId      string `json:"walletId"`
	CyclesBalance string `json:"cyclesBalance"`
}

type IcpCanisterPage struct {
	Data     []IcpCanisterVo `json:"data"`
	Total    int             `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
}
