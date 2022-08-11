// @Author markwang <wangyu@uniondrug.cn>
// @Date   2022/8/9

package gsnrspudchsdk

import "github.com/fuyibing/sdk"

const (
	Name = "gs-nrsp-udch"
)

// 永久授权码创建
func TokenPermanentCodeCreate(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/token/permanent/code/create").SetBody(body).Run(ctx)
}

// 授权链接
func TokenAuthUrl(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/token/auth/url").SetBody(body).Run(ctx)
}

// 创建订单
func OrderCreate(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/order/create").SetBody(body).Run(ctx)
}
