// @Author markwang <wangyu@uniondrug.cn>
// @Date   2022/4/1

package javambssdk

import "github.com/fuyibing/sdk"

const (
	Name = "java.mbs"
)

// 发布
func TopicPublish(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/topic/publish").SetBody(body).Run(ctx)
}
