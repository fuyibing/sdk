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

// GetMerchantByOrganizationId 获取连锁信息.
func GetMerchantByOrganizationId(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/partner/detail").SetBody(body).Run(ctx)
}

// GetInsurerByOrganizationId 获取保司信息.
func GetInsurerByOrganizationId(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/insurer/detail").SetBody(body).Run(ctx)
}

// GetCompanyByOrganizationId 获取药联公司信息.
func GetCompanyByOrganizationId(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/company/info").SetBody(body).Run(ctx)
}

// GetRelationshipByOrganizationId 获取商业公司与商户关系.
func GetRelationshipByOrganizationId(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/relationship/listing").SetBody(body).Run(ctx)
}

// GetAutoCreateHxMerchant 批量获取连锁自动生成换新结算单配置.
func GetAutoCreateHxMerchant(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/partner/setting/auto/create/hx/merchant").SetBody(body).Run(ctx)
}

// GetWorker 获取连锁员工信息
func GetPartnerWorkerPaging(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/partner/worker/paging").SetBody(body).Run(ctx)
}
