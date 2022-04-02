// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

// SDK for bill of finance.
package insuresdk

import "github.com/fuyibing/sdk"

const (
	Name = "jm.insure"
)

// GetByInsurerIdAndPolicyNo get policy by insurerId+policyNo.
func GetByInsurerIdAndPolicyNo(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/policy/getByInsurerIdAndPolicyNo").SetBody(body).Run(ctx)
}

// ChangeOrderPaidTime change order pay time.
func ChangeOrderPaidTime(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/mbs/order/update/orderPaidTime").SetBody(body).Run(ctx)
}

// ListReverseByOrder .
func ListReverseByOrder(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/reverse/listReverseByOrder").SetBody(body).Run(ctx)
}
