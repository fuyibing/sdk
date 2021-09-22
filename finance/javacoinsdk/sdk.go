package javacoinsdk
// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

import "github.com/fuyibing/sdk"

const (
	Name = "java.coin"
)

// GetPaymentPageByBusinessNo 获取连锁开票配置.
func GetPaymentPageByBusinessNo(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/payment/page").SetBody(body).Run(ctx)
}
