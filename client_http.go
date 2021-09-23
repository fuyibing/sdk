// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

package sdk

import (
	"sync"
	"sync/atomic"
)

var (
	poolHttp      *sync.Pool
	poolHttpIndex uint64
)

// HTTPSdk结构体.
//
//   req := NewHttp("service")
//   req.SetBody(`{"key":"value"}`)
//   req.SetContentType("application/json")
//   req.SetMethod("POST")
//   req.SetRoute("/site/index")
//   req.SetTimeout(10)
//
//   ctx := log.NewContext()
//   res := req.Run(ctx)
//
//   if res.HasError(){
//       return fmt.Errorf("sdk request error: %v", res.GetError())
//   }
//
type ClientHttp struct {
	index   uint64 // pool index.
	route   string // service route.
	service string // consul service name.
	request *ClientRequest
}

// 执行SDK.
// 本方法必须在Set类方法之后执行, 并返回SDK请求结果.
//
//   headers := []map[string]string{
//      {
//          "content-type":"application/json",
//      },
//      {
//          "user-agent":"app/ver",
////    }
//   }
//
//   ctx := context.Background()
//   sdk.Run(ctx, headers...)
func (o *ClientHttp) Run(ctx interface{}, headers ...map[string]string) (res *ClientResponse) {
	host, err := Consul.Get(ctx, o.service)
	// normal error.
	if err != nil {
		res = &ClientResponse{err: err}
		res.end()
		return res
	}
	o.request.SetUrl(host + "/" + o.route)

	for _, header := range headers {
		for k, v := range header {
			o.request.SetHeader(k, v)
		}
	}

	return o.request.Run(ctx)
}

// 设置请求入参.
// 请求入参支持string(JSON字符串), []byte(Byte).
func (o *ClientHttp) SetBody(body interface{}) *ClientHttp {
	o.request.SetBody(body)
	return o
}

// 设置请求格式.
func (o *ClientHttp) SetContentType(ct string) *ClientHttp {
	o.request.SetContentType(ct)
	return o
}

// 设置请求方法.
// SDK默认以POST方式向服务提供方发起请求, 可以在SDK中设置请求方式, 以覆盖
// 默认请求方式.
func (o *ClientHttp) SetMethod(method string) *ClientHttp {
	o.request.SetMethod(method)
	return o
}

// 设置路由地址.
// 路由用于区分服务, 不同的路由用于识别不同的服务.
func (o *ClientHttp) SetRoute(route string) *ClientHttp {
	if route[0:1] == "/" {
		o.route = route[1:]
	} else {
		o.route = route
	}
	return o
}

// 设置超时时长.
// 本方法可选, 若不指定则使用配置文件预置的标准超时时长(默认: 25秒), SDK发布
// 时, 可以针对个别方法指定个别超时时长.
func (o *ClientHttp) SetTimeout(seconds int) *ClientHttp {
	o.request.SetTimeout(seconds)
	return o
}

// 开始HTTPSdk.
// 重置HTTPSdk字段, 避免池中实例复用时污染.
func (o *ClientHttp) begin(service string) {
	o.request = NewRequest()
	o.service = service
}

// 释放HTTPSdk.
// 本方法在Run()方法执行完成后触发, 将用完的HTTPSdk放回
// 池中.
func (o *ClientHttp) release() {
	o.request = nil
	poolHttp.Put(o)
}

// 获取HTTPSdk.
// 从HTTPSdk池中提取Client实例, 并重置实例.
func NewHttp(service string) *ClientHttp {
	o := poolHttp.Get().(*ClientHttp)
	o.begin(service)
	return o
}

// 初始化HTTPSdk池.
func initPoolHttp() {
	poolHttp = &sync.Pool{New: func() interface{} {
		return &ClientHttp{
			index: atomic.AddUint64(&poolHttpIndex, 1),
		}
	}}
}
