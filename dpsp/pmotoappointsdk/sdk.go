package pmotoappointsdk

import "github.com/fuyibing/sdk"

const (
	Name = "pm-oto-appoint"
)

// 小程序会员购买埋点
func GoBuried(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/order/goSensor").SetBody(body).Run(ctx)
}
