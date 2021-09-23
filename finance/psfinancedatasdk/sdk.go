// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

// SDK for bill of finance.
package psfinancedatasdk

import "github.com/fuyibing/sdk"

const (
	Name = "ps-finance-data"
)

// GetBusinessByOrganizationId 获取组织详情.
func GetBusinessByOrganizationId(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/organize/detail").SetBody(body).Run(ctx)
}

// GetPartnerBillInfoByOrganizationId 获取连锁开票配置.
func GetPartnerBillInfoByOrganizationId(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/partner/bill/info/unit/detail").SetBody(body).Run(ctx)
}

// GetInsurerBillInfoByOrganizationId 获取保司开票配置.
func GetInsurerBillInfoByOrganizationId(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/insurer/bill/info/detail").SetBody(body).Run(ctx)
}

// GetUnionDrugBillInfoByOrganizationId 获取药联子公司开票配置.
func GetUnionDrugBillInfoByOrganizationId(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/company/bill/info/detail").SetBody(body).Run(ctx)
}

// GetPartnerTaxServiceUnitDetail 财税服务费配置核算单位详情
func GetPartnerTaxServiceUnitDetail(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/partner/tax/service/unit/detail").SetBody(body).Run(ctx)
}

// GetBankAccountDetailByType 查询银行帐号信息
func GetBankAccountDetailByType(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bank/account/detail/by/type").SetBody(body).Run(ctx)
}
