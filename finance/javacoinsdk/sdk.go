package javacoinsdk

// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

import "github.com/fuyibing/sdk"

const (
	Name = "java.coin"
)

// PaymentPaging 付款单列表.
func PaymentPaging(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/payment/page").SetBody(body).Run(ctx)
}

// PaymentCreate 创建付款单.
func PaymentCreate(ctx interface{}, body interface{}, headers ...map[string]string) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/payment/create").SetBody(body).Run(ctx, headers...)
}

// PaymentDelete 删除付款单.
func PaymentDelete(ctx interface{}, body interface{}, headers ...map[string]string) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/payment/remove").SetBody(body).Run(ctx, headers...)
}

// MonthPoolStatistics 查询资金池月统计.
func MonthPoolStatistics(ctx interface{}, body interface{}, headers ...map[string]string) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/fundAccount/statisticsRangeDate").SetBody(body).Run(ctx, headers...)
}

// MonthPoolStatisticsV2 查询资金池月统计v2.
func MonthPoolStatisticsV2(ctx interface{}, body interface{}, headers ...map[string]string) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/fundRecordDetail/statisticsV2").SetBody(body).Run(ctx, headers...)
}

// GetFundAccount 查询资金账户详情.
func GetFundAccount(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/fundAccount/get").SetBody(body).Run(ctx)
}

// GetFundAccountBalance 查询资金账户余额.
func GetFundAccountBalance(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/fundAccount/queryBalance").SetBody(body).Run(ctx)
}

// PaymentTaxDeductSummary 查询资金付款单，指定实际付款时间段内，收款方（连锁）直付财税金额汇总.
func PaymentTaxDeductSummary(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/payment/summary/taxDeductAmount").SetBody(body).Run(ctx)
}
