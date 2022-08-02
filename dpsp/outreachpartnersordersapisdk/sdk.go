package outreachpartnersordersapisdk

import "github.com/fuyibing/sdk"

const (
	Name = "outreach-partners-orders-api"
)

// 对接组 会员推送接口
func MemberInsert(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/api/memberInsert").SetBody(body).Run(ctx)
}
