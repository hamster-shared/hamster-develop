package parameter

type LoginParam struct {
	Code         string `json:"code"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}
