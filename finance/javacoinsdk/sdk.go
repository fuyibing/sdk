package javacoinsdk

// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

import "github.com/fuyibing/sdk"

const (
	Name = "java.coin"
)

// GetPaymentPageByBusinessNo 查询付款单列表.
func GetPaymentPageByBusinessNo(ctx interface{}, body interface{}) *sdk.ClientResponse {
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