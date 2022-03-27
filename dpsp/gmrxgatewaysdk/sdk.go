//@author: markwang <617514823@qq.com>
//@date: 2022/3/28
package gmrxgatewaysdk

import "github.com/fuyibing/sdk"

const (
	Name = "rxapi"
)

// 供应商申请处方
func SupplierRxApply(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/supplier/rx/apply").SetBody(body).Run(ctx)
}
