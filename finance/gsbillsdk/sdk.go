// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

// SDK for bill of finance.
package gsbillsdk

import (
	"github.com/fuyibing/sdk"
)

const (
	Name = "gs.bill"
)

// BillPayoutDirectApply 应付-财务结算-直付结算-开票申请
func BillPayoutDirectApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/payout/direct/apply").SetBody(body).Run(ctx)
}

// BillPayoutIntegralApply 应付-财务结算-积分结算-开票申请
func BillPayoutIntegralApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/payout/integral/apply").SetBody(body).Run(ctx)
}

// BillPayoutCommissionApply 应付-财务结算-佣金结算-开票申请
func BillPayoutCommissionApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/payout/commission/apply").SetBody(body).Run(ctx)
}

// BillPayoutHealthApply 应付-财务结算-健康结算-开票申请
func BillPayoutHealthApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/payout/health/apply").SetBody(body).Run(ctx)
}

// BillPayoutRenewApply 应付-财务结算-换新结算-开票申请
func BillPayoutRenewApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/payout/renew/apply").SetBody(body).Run(ctx)
}

// BillIncomeRenewApply 应收-财务结算-换新结算-开票申请
func BillIncomeRenewApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/income/renew/apply").SetBody(body).Run(ctx)
}

// BillIncomeIndemnityApply 应收-投保理赔-药联理赔单理赔款-开票申请
func BillIncomeIndemnityApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/income/indemnity/apply").SetBody(body).Run(ctx)
}

// BillIncomeCommissionApply 应收-投保理赔-保单手续费-开票申请
func BillIncomeCommissionApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/income/commission/apply").SetBody(body).Run(ctx)
}

// BillIncomePurchaseApply 应收-采购系统-采购福联社-开票申请
func BillIncomePurchaseApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/income/purchase/apply").SetBody(body).Run(ctx)
}

// BillIncomeClientApply 应收-C端-C端微信公众号-开票申请
func BillIncomeClientApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/income/client/apply").SetBody(body).Run(ctx)
}

// BillSaleCreate 创建开票单销售清单
func BillSaleCreate(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/sale/create").SetBody(body).Run(ctx)
}

// BillProtocolCreate 创建开票协议
func BillProtocolCreate(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/protocol/create").SetBody(body).Run(ctx)
}

// BillProtocolDetail 查询开票协议
func BillProtocolDetail(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/protocol/detail").SetBody(body).Run(ctx)
}

// InvoiceCheck 发票检测
func InvoiceCheck(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/invoice/check").SetBody(body).Run(ctx)
}

// InvoiceCreate 发票录入
func InvoiceCreate(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/invoice/create").SetBody(body).Run(ctx)
}
