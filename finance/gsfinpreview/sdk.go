// author: Haven <chenhao@uniondrug.com>
// date: 2022-03-24

package gsfinpreview

import "github.com/fuyibing/sdk"

const (
	Name = "gs-fin-preview"
)

// OrderDetail 订单详情.
func OrderDetail(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/order/detail").SetBody(body).Run(ctx)
}
