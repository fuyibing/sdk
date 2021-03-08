// author: wsfuyibing <websearch@163.com>
// date: 2021-03-03

package sdk

import (
	"net/http"
)

func Get(ctx interface{}, url string) *ClientResponse {
	return NewRequest().
		SetMethod(http.MethodGet).
		SetContentType(Config.ServiceRequestContentType).
		SetUserAgent(Config.ServiceUserAgent).
		SetUrl(url).
		Run(ctx)
}

func Post(ctx interface{}, url string, body interface{}) *ClientResponse {
	return NewRequest().
		SetMethod(http.MethodPost).
		SetContentType(Config.ServiceRequestContentType).
		SetUserAgent(Config.ServiceUserAgent).
		SetBody(body).
		SetUrl(url).
		Run(ctx)
}
