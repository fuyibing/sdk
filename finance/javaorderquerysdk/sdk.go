package javaorderquerysdk

// author: Haven <chenhao@uniondrug.com>
// date: 2022-03-02

import "github.com/fuyibing/sdk"

const (
	Name = "java.order.query"
)

// OrderQueryMain 查询主订单接口.
func OrderQueryMain(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/order/query/main").SetBody(body).Run(ctx)
}

// OrderQuerySub 查询子订单接口.
func OrderQuerySub(ctx interface{}, body interface{}, headers ...map[string]string) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/order/query/sub").SetBody(body).Run(ctx, headers...)
}
