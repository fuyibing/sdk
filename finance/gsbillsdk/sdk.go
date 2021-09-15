// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

// SDK for bill of finance.
package gsbillsdk

import "github.com/fuyibing/sdk"

const (
	Name = "gs.bill"
)

// Create bill.
func BillPayoutBusinessApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/bill/payout/business/apply").SetBody(body).Run(ctx)
}
