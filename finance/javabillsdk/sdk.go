// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

// SDK for bill of finance.
package javabillsdk

import "github.com/fuyibing/sdk"

const (
	Name = "js-fin-bill"
)

// PayoutBillApply 应付开票
func PayoutBillApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/payout/save").SetBody(body).Run(ctx)
}

// IncomeBillApply 应收开票
func IncomeBillApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/income/save").SetBody(body).Run(ctx)
}

// InvoicePageByBusinessNo 根据 开票单业务单号 查询发票
func InvoicePageByBusinessNo(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/invoice/pageByBusinessNo").SetBody(body).Run(ctx)
}

// InvoiceGetSignAndSendTime 根据 开票单业务单号 查询发票最大的时间
func InvoiceGetSignAndSendTime(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/invoice/get/signAndSendTime").SetBody(body).Run(ctx)
}

// InvoicePaging
func InvoicePaging(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/invoice/page").SetBody(body).Run(ctx)
}

// OpenOnlineBilling 开启在线开票
func OpenOnlineBilling(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/openOnlineBilling").SetBody(body).Run(ctx)
}
