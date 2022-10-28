package pmdposssdk

import "github.com/fuyibing/sdk"

const (
	Name = "pm-dposs-operator"
)

func ChangeLogo(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/wechat/changeLogo").SetBody(body).Run(ctx)
}

func UpdateMediaId(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/wechat/updateMediaId").SetBody(body).Run(ctx)
}
