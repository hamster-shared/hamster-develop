package parameter

type LoginParam struct {
	Code         string `json:"code"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
