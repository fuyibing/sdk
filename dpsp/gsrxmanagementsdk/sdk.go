//@author: markwang <617514823@qq.com>
//@date: 2022/3/24
package gsrxmanagementsdk

import "github.com/fuyibing/sdk"

const (
	Name = "gs-rx-management"
)

// 处方更新
func RxUpdate(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/rx/update").SetBody(body).Run(ctx)
}

// 处方供应商回调写入
func RxSupplierCallbackAdd(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/rx/supplier/callback/add").SetBody(body).Run(ctx)
}
