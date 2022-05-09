//@author: markwang <617514823@qq.com>
//@date: 2022/4/19
package merchantmodulesdk

import "github.com/fuyibing/sdk"

const (
	Name = "merchant.module"
)

// 商户开通/关闭会员服务
func BackEndEditIsMember(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/organizebasebackend/editismember").SetBody(body).Run(ctx)
}

// 商户开通连锁小程序服务
func BackEndEditIsAppletService(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/organizebasebackend/editIsAppletService").SetBody(body).Run(ctx)
}

// 商户详情
func organizebaseInfo(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/organizebase/info").SetBody(body).Run(ctx)
}
