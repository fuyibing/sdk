// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

// SDK for bill of finance.
package bill

import (
	"github.com/fuyibing/sdk"
)

const (
	Name = "pm-fin-shares"
)

// Create bill.
func Create(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/insure/policy/paging").SetBody(body).Run(ctx)
}
