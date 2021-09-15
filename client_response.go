// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

package sdk

import (
	"encoding/json"
	"fmt"
	"strconv"
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
	body     []byte      // SDK返回原文.
	code     int         // error code.
	data     interface{} // Response data.
	duration float64     // SDK请求用时.
	err      error       // SDK请求出错.
	retry    bool        // allow retry.
	time     time.Time   // SDK请求实际开始时间.
	url      string      // SDK请求完整地址.
}

// 创建请求结果.
func NewResponse() *ClientResponse {
	return &ClientResponse{}
}

// 读取请求原文.
func (o *ClientResponse) GetBody() []byte {
	return o.body
}

// 读取请求原文.
func (o *ClientResponse) GetContents() string {
	return fmt.Sprintf("%s", o.body)
}

// 读取请求原文.
func (o *ClientResponse) GetData() interface{} {
	return o.data
}

// 读取请求时长.
func (o *ClientResponse) GetDuration() float64 {
	return o.duration
}

// 读取错误原因.
func (o *ClientResponse) GetError() (int, error) {
	return o.code, o.err
}

// 返回服务源地址.
func (o *ClientResponse) GetUrl() string {
	return o.url
}

// SDK请求过程中是否有错误.
// 可能的错误主要包括以下几类.
//   1. 运行中Panic抛错.
//   2. Consul调用出错.
//      1) 创建Consul客户端出错.
//      2) 请求Consul服务出错.
//      3) 服务未注册.
//   3. 请求Http服务出错.
//      1)
//      2)
func (o *ClientResponse) HasError() bool {
	return o.err != nil
}

// 结束请求过程.
func (o *ClientResponse) end() {
	// Fill default error code.
	defer func() {
		if o.err != nil && o.code == 0 {
			o.code = ErrorDefaultCode
		}
	}()
	// Outer error.
	if o.err != nil {
		return
	}
	// Null response.
	if o.body == nil {
		o.err = fmt.Errorf("nil response")
		return
	}
	// Parse response body.
	tmp := &struct {
		Errno    interface{} `json:"errno"`
		Error    interface{} `json:"error"`
		Data     interface{} `json:"data"`
		DataType interface{} `json:"dataType"`
	}{}
	// Parse json response failed.
	if err := json.Unmarshal(o.body, tmp); err != nil {
		o.err = fmt.Errorf("invalid json response: %v", err)
		return
	}
	// Remote error response.
	if code := fmt.Sprintf("%v", tmp.Errno); "0" != code {
		o.err = fmt.Errorf("%v", tmp.Error)
		errno, _ := strconv.ParseInt(code, 0, 32)
		o.code = int(errno)
		return
	}
	// Succeed.
	o.data = tmp.Data
}

//
func (o *ClientResponse) Unmarshal(v interface{}) error {
	data, _ := json.Marshal(o.data)
	return json.Unmarshal(data, v)
}
