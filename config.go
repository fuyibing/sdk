// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

package sdk

import (
	"net/http"
)

// Configuration.
type configuration struct {
	// Consul server addr.
	// default: 127.0.0.1:8500.
	Address string

	// Cache lifetime.
	CacheLifetime float64

	RequestMethod string

	// Call service timeout.
	RequestTimeout int

	RequestContentType string

	// Consul server scheme.
	// default: scheme.
	Scheme string

	// Consul server timeout.
	Timeout int

	UserAgent string
}

func (o *configuration) initialize() {
	o.Address = "udsdk.turboradio.cn"
	o.CacheLifetime = CacheLifetime
	o.Scheme = "http"
	o.Timeout = Timeout
	o.RequestMethod = http.MethodPost
	o.RequestTimeout = RequestTime
	o.RequestContentType = "application/json"
	o.UserAgent = "x-sdk/3.2.1"
}
