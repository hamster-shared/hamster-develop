package parameter

type UpdateDfxData struct {
	JsonData string `json:"jsonData"`
}

type RechargeCanisterParam struct {
	CanisterId string `json:"canisterId" binding:"required"`
	Amount     string `json:"amount" binding:"required"`
}

type RedeemFaucetCouponParam struct {
	Coupon string `json:"coupon" binding:"required"`
}
