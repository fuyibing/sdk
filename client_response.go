// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

package sdk

import (
	"time"
)

// Client response interface.
type IResponse interface {
	GetBody() interface{}
	GetDuration() float64
	GetError() error
	HasError() bool
}

// SDK结果结构体.
type ClientResponse struct {
	body     []byte  // SDK返回原文.
	duration float64 // SDK请求用时.
	err      error   // SDK请求出错.
	retry    bool
	time     time.Time // SDK请求实际开始时间.
	url      string    // SDK请求完整地址.
	// n.
	host string // SDK请求服务地址(注册在Consul中的服务地址).
}

// 创建请求结果.
func NewResponse() *ClientResponse {
	return &ClientResponse{}
}

// 读取请求原文.
func (o *ClientResponse) GetBody() []byte {
	return o.body
}

// 读取请求时长.
func (o *ClientResponse) GetDuration() float64 {
	return o.duration
}

// 读取错误原因.
func (o *ClientResponse) GetError() error {
	return o.err
}

// 请求是否出错.
func (o *ClientResponse) HasError() bool {
	return o.err != nil
}

// 结束请求过程.
func (o *ClientResponse) end() {}
