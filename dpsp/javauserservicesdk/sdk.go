package javauserservicesdk

import "github.com/fuyibing/sdk"

const (
	Name = "java.user.service"
)

// 添加用户
func UserInfoAddOr(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/v2/api/userBasic/addOrQuery").SetBody(body).Run(ctx)
}

// 基于用户unionId批量查看用户详情
func BatchByUnionIds(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/v2/api/userBasic/batchByUnionIds").SetBody(body).Run(ctx)
}

// 获取用户真实证件信息
func UserCardGetReal(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/v2/api/userCard/getReal").SetBody(body).Run(ctx)
}
