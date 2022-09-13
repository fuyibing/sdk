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

// 永久授权码重置
func TokenPermanentCodeReset(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/token/permanent/code/reset").SetBody(body).Run(ctx)
}

// 授权链接
func TokenAuthUrl(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/token/auth/url").SetBody(body).Run(ctx)
}

// 创建订单
func OrderCreate(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/order/create").SetBody(body).Run(ctx)
}

// 订单支付成功
func OrderSuccessPay(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/order/success/pay").SetBody(body).Run(ctx)
}

// 账号激活
func OrderAccountActive(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/order/account/active").SetBody(body).Run(ctx)
}

// 创建【联系我】二维码
func ExternalContactWayCreate(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/external/contact/way/create").SetBody(body).Run(ctx)
}

// 查询【联系我】二维码
func ExternalContactWayGet(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/external/contact/way/get").SetBody(body).Run(ctx)
}

// 用户详情
func OauthUserDetail(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/oauth/user/detail").SetBody(body).Run(ctx)
}

// 标签回调
func TagCallback(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/tag/callback").SetBody(body).Run(ctx)
}
