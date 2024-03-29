//@author: markwang <617514823@qq.com>
//@date: 2022/3/28
package gmrxgatewaysdk

import "github.com/fuyibing/sdk"

const (
	Name = "gm-rx-gateway"
)

// 供应商申请处方
func SupplierRxApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/supplier/rx/apply").SetBody(body).Run(ctx)
}

func SupplierPayStatus(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/supplier/pay/status").SetBody(body).Run(ctx)
}
