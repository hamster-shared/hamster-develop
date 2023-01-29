package utils

import "github.com/go-resty/resty/v2"

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
