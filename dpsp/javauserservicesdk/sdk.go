package javauserservicesdk

import "github.com/fuyibing/sdk"

const (
	Name = "java.user.service"
)

// 添加用户
func UserInfoAddOr(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/v2/api/userBasic/addOrQuery").SetBody(body).Run(ctx)
}
