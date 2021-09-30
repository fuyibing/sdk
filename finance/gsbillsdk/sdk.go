// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

// SDK for bill of finance.
package gsbillsdk

import "github.com/fuyibing/sdk"

const (
	Name = "gs.bill"
)

// BillInternalApply 申请药联票据中心-内部体系开票.
func BillInternalApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/internal/apply").SetBody(body).Run(ctx)
}

// BillExternalApply 申请药联票据中心-外部体系开票.
func BillExternalApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/external/apply").SetBody(body).Run(ctx)
}

// BillSaleCreate 创建销售清单
func BillSaleCreate(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/sale/create").SetBody(body).Run(ctx)
}
