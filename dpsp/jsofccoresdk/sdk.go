package jsofccoresdk

import "github.com/fuyibing/sdk"

const (
	Name = "js-ofc-core"
)

// 申请履约
func OfcAudit(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/ofc/audit").SetBody(body).Run(ctx)
}
