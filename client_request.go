// author: wsfuyibing <websearch@163.com>
// date: 2021-03-06

package sdk

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/fuyibing/log/v2"
)

// 请求结构体.
type ClientRequest struct {
	url         string
	body        *bytes.Buffer
	headers     map[string]string
	method      string
	contentType string
	timeout     time.Duration
	userAgent   string
}

// 创建请求实例.
func NewRequest() *ClientRequest {
	return &ClientRequest{
		body:        bytes.NewBufferString(``),
		headers:     make(map[string]string),
		method:      Config.RequestMethod,
		timeout:     time.Duration(Config.RequestTimeout) * time.Second,
		contentType: Config.RequestContentType,
		userAgent:   Config.UserAgent,
	}
}

// 执行Request.
func (o *ClientRequest) Run(ctx interface{}) *ClientResponse { return o.RunRetry(ctx, 1) }

// 执行Request.
func (o *ClientRequest) RunRetry(ctx interface{}, max int) (res *ClientResponse) {
	// min attempts reset.
	if max < 1 {
		max = 1
	}
	// loop attempts.
	for try := 1; try <= max; try++ {
		// succeed.
		if res = o.runner(ctx, try); res.err != nil {
			break
		}
		// max attempts.
		if try == max || !res.retry {
			break
		}
	}
	// end request.
	if res != nil {
		res.end()
	}
	return
}

// 设置请求入参.
func (o *ClientRequest) SetBody(body interface{}) *ClientRequest {
	switch reflect.TypeOf(body).Kind() {
	case reflect.Slice:
		o.body = bytes.NewBuffer(body.([]byte))
	case reflect.String:
		o.body = bytes.NewBufferString(body.(string))
	}
	return o
}

// 设置请求类型.
func (o *ClientRequest) SetContentType(contentType string) *ClientRequest {
	o.contentType = contentType
	return o
}

// 设置请求头.
func (o *ClientRequest) SetHeader(key, value string) *ClientRequest {
	o.headers[key] = value
	return o
}

// 设置请求方式.
func (o *ClientRequest) SetMethod(method string) *ClientRequest {
	o.method = method
	return o
}

// 设置请求超时时长.
func (o *ClientRequest) SetTimeout(seconds int) *ClientRequest {
	o.timeout = time.Duration(seconds) * time.Second
	return o
}

// 设置请求地址.
func (o *ClientRequest) SetUrl(url string) *ClientRequest {
	o.url = url
	return o
}

// 设置浏览器名称.
func (o *ClientRequest) SetUserAgent(userAgent string) *ClientRequest {
	o.userAgent = userAgent
	return o
}

// 执行Request过程.
func (o *ClientRequest) runner(ctx interface{}, try int) (res *ClientResponse) {
	// 1. init response.
	res = NewResponse()
	// 2. end request.
	res.time = time.Now()
	defer func() {
		res.duration = time.Now().Sub(res.time).Seconds()
		if r := recover(); r != nil {
			// 2.1 catch panic.
			log.Client.Errorfc(ctx, "[sdk][try=%d][d=%f] send {%s} request on {%s} fatal.", try, res.duration, o.method, o.url)
			// 2.2 fatal trace.
			fatal := strings.TrimSpace(fmt.Sprintf("[sdk][try=%d] %v.", try, r))
			for i := 1; ; i++ {
				if _, f, l, got := runtime.Caller(i); got {
					fatal += fmt.Sprintf("\n%d. %s:%d", i, strings.TrimSpace(f), l)
				} else {
					break
				}
			}
			log.Client.Errorfc(ctx, fatal)
		} else if log.Config.InfoOn() {
			if res.err == nil {
				// 2.3 normal error.
				log.Client.Infofc(ctx, "[sdk][try=%d][d=%f] send {%s} request on {%s} succeed.", try, res.duration, o.method, o.url)
			} else {
				// 2.4 succeed response.
				log.Client.Infofc(ctx, "[sdk][try=%d][d=%f] send {%s} request on {%s} error: %v.", try, res.duration, o.method, o.url, res.err)
			}
		}
	}()
	// 3. prepare.
	var request *http.Request
	var response *http.Response
	// 3.1 create request.
	if request, res.err = http.NewRequest(o.method, o.url, o.body); res.err != nil {
		return
	}
	// 3.2 assign request header.
	//     todo: assign request tracing, open tracing.
	request.Header.Set("content-type", o.contentType)
	request.Header.Set("user-agent", o.userAgent)
	for key, value := range o.headers {
		request.Header.Set(key, value)
	}
	// 3.3 send request.
	if response, res.err = (&http.Client{Timeout: o.timeout}).Do(request); res.err != nil {
		res.retry = true
		return
	}
	// 3.4 call response body.
	defer func() {
		if response.Body != nil {
			_ = response.Body.Close()
		}
	}()
	if res.body, res.err = ioutil.ReadAll(response.Body); res.err != nil {
		return
	}
	return
}
