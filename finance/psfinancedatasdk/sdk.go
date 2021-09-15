// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

// SDK for bill of finance.
package psfinancedatasdk

import "github.com/fuyibing/sdk"

const (
	Name = "ps-finance-data"
)

// 获取开票配置.
func GetBillInfoByOrganizationId(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/partner/bill/info/unit/detail").SetBody(body).Run(ctx)
}
