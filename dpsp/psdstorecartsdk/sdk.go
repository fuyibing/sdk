/**
* @Author: kang.xiaoqiang
* @Date: 2022/7/21 17:33
 */

package psdstorecartsdk

import "github.com/fuyibing/sdk"

const (
	Name = "ps-dstore-cart"
)

// 预订单详情
func TrolleyDetail(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/trolley/detail").SetBody(body).Run(ctx)
}
