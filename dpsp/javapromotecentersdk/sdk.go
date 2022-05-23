package javapromotecentersdk

import "github.com/fuyibing/sdk"

const (
	Name = "java.promotecenter.service"
)

// 发放卡
func SendCard(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/card/sendCard").SetBody(body).Run(ctx)
}

// 通过券id查询方案和标签信息
func QuerySchemeInfo(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/coupon/querySchemeInfo").SetBody(body).Run(ctx)
}

// 通过cardId查询卡详情
func QueryCardDetailByCardId(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/card/queryCardDetailByCardId").SetBody(body).Run(ctx)
}
