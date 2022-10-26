/**
* @Author: kang.xiaoqiang
* @Date: 2022/7/21 17:33
 */

package pmdposssdk

import "github.com/fuyibing/sdk"

const (
	Name = "pm-dposs-operator"
)

func ChangeLogo(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/wechat/changeLogo").SetBody(body).Run(ctx)
}
