// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

// SDK for bill of finance.
package jmbsl

import (
	"github.com/fuyibing/sdk"
	"github.com/google/uuid"
	"strings"
)

const (
	Name = "java.mbsl"
)

type MsgBody struct {
	TopicName string
	TopicTag  string
	MsgKey    string
	ReqNo     string
	Message   []interface{}
}

func (o *MsgBody) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"topic":   o.TopicName,
		"tag":     o.TopicTag,
		"reqNo":   o.getUUID(),
		"msgKey":  o.MsgKey,
		"message": o.Message,
	}
}

func (o *MsgBody) getUUID() string {
	u, _ := uuid.NewUUID()
	return strings.ReplaceAll(u.String(), "-", "")
}

func Publish(ctx interface{}, body *MsgBody) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/topic/publish").SetBody(body.ToMap()).Run(ctx)
}

func Batch(ctx interface{}, body *MsgBody) *sdk.ClientResponse {
	return sdk.NewHttp(Name).SetRoute("/topic/batch").SetBody(body.ToMap()).Run(ctx)
}
