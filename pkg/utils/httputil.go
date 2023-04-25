package utils

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type HttpUtil struct {
	client *resty.Client
}

func NewHttp() *HttpUtil {
	return &HttpUtil{client: resty.New()}
}

func (h *HttpUtil) NewRequest() *resty.Request {
	res := h.client.R().
		SetHeader("Accept", "application/json, text/plain, */*").
		SetHeader("Content-Type", "application/json;charset=UTF-8").
		SetContentLength(true)
	return res
}

func MetaScanHttpRequestToken() string {
	url := "https://account.metatrust.io/realms/mt/protocol/openid-connect/token"
	token := struct {
		AccessToken      string `json:"access_token"`
		ExpiresIn        int64  `json:"expires_in"`
		RefreshExpiresIn int64  `json:"refresh_expires_in"`
		RefreshToken     string `json:"refresh_token"`
		TokenType        string `json:"token_type"`
		NotBeforePolicy  int    `json:"not-before-policy"`
		SessionState     string `json:"session_state"`
		Scope            string `json:"scope"`
	}{}
	res, err := NewHttp().NewRequest().SetFormData(map[string]string{
		"grant_type": "password",
		"username":   "tom@hamsternet.io",
		"password":   "pysded-hismoh-3Dagcy",
		"client_id":  "webapp",
	}).SetResult(&token).SetHeader("Content-Type", "application/x-www-form-urlencoded").Post(url)
	if res.StatusCode() != 200 {
		return ""
	}
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
	//return token.AccessToken
}
