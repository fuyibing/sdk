// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

// Package thirdparty SDK for outreach
package thirdparty

import (
	"github.com/fuyibing/sdk"
)

const (
	Name = "outreach.resource.api"
)

func CallThirdPartyService(ctx interface{}, body interface{}) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/api/adapter").SetBody(body).Run(ctx)
}
