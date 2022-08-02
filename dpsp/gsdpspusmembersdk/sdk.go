package gsdpspusmembersdk

import "github.com/fuyibing/sdk"

const (
	Name = "gs-dpsp-us-member"
)

// 添加会员
func UserCreate(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/user/create").SetBody(body).Run(ctx)
}
