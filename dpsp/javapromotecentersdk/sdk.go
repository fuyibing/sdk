package javapromotecentersdk

import "github.com/fuyibing/sdk"

const (
	Name = "java.promotecenter.service"
)

// 发放卡
func SendCard(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/card/sendCard").SetBody(body).Run(ctx)
}