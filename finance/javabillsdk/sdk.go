// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

// SDK for bill of finance.
package javabillsdk

import "github.com/fuyibing/sdk"

const (
	Name = "jm.bill"
)

// Create bill.
func PayoutBillApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/payout/save").SetBody(body).Run(ctx)
}
