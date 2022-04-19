//@author: markwang <617514823@qq.com>
//@date: 2022/4/19
package javagoodscentermngsdk

import "github.com/fuyibing/sdk"

const (
	Name = "java.goods.center.mng.module"
)

// 创建/编辑商品
func SkuRegister(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/configsku/register").SetBody(body).Run(ctx)
}

// 商品上架
func SkuChannels(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/configsku/channels").SetBody(body).Run(ctx)
}
