// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

// SDK for bill of finance.
package insuresdk

import "github.com/fuyibing/sdk"

const (
	Name = "jm.insure"
)

// Create bill.
func GetByInsurerIdAndPolicyNo(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/policy/getByInsurerIdAndPolicyNo").SetBody(body).Run(ctx)
}
