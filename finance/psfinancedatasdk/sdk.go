// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

// SDK for bill of finance.
package psfinancedatasdk

import "github.com/fuyibing/sdk"

const (
	Name = "ps-finance-data"
)

// GetBillInfoByOrganizationIdAndUnitId 根据组织ID、核算单位ID获取开票配置(连锁、保司、药联集团公司等)
func GetBillInfoByOrganizationIdAndUnitId(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/info/detail").SetBody(body).Run(ctx)
}
