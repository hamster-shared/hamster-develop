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
