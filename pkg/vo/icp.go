package vo

type UserIcpInfoVo struct {
	UserId     int    `json:"userId"`
	AccountId  string `json:"accountId"`
	IcpBalance string `json:"icpBalance"`
}

type IcpCanisterBalanceVo struct {
	UserId        int    `json:"userId"`
	CanisterId    string `json:"canisterId"`
	CyclesBalance string `json:"cyclesBalance"`
}

type IcpCanisterPage struct {
	Data     []IcpCanisterVo `json:"data"`
	Total    int             `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
}

type CanisterStatusRes struct {
	Status  string `json:"status"`
	Balance string `json:"balance"`
}
