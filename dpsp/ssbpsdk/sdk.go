// @Author markwang <wangyu@uniondrug.cn>
// @Date   2022/5/5

package ssbpsdk

import "github.com/fuyibing/sdk"

const (
	Name = "ssbp"
)

// 查询会员规则信息
func InfoOutQueryMemberInfo(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/infoOut/queryMemberInfo").SetBody(body).Run(ctx)
}
