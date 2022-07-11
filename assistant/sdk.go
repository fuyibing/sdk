// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

package assistant

import (
	"github.com/fuyibing/sdk"
)

const (
	Name = "assistant.module"
)

func UserInfo(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/users/info").SetBody(body).Run(ctx)
}
