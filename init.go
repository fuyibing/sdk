// author: wsfuyibing <websearch@163.com>
// date: 2021-03-03

// Package service sdk.
//
// Any application publish sdk functions to specified directory of this application,
// then all function can be called in any application.
//
//   ctx := log.NewContext()
//   res := bill.GetByNo(ctx, `{"billNo":"example-no"}`)
//   if res.HasError(){
//       return res.GetError()
//   }
//
package sdk

import (
	"net/http"
	"regexp"
	"sync"
)

const (
	ErrorDefaultCode = 1
)

const (
	DefaultCacheLifetime  float64 = 3.0
	DefaultConsulAddress          = "127.0.0.1:8500"
	DefaultConsulScheme           = "http"
	DefaultConsulTimeout          = 2
	DefaultContentType            = "application/json"
	DefaultRequestMethod          = http.MethodPost
	DefaultRequestTimeout         = 25
	DefaultServiceScheme          = "http"
	Name                          = "SDK"
	Version                       = "1.0.0"
)

var (
	Config               *configuration
	Consul               *consul
	regexpServiceAddress = regexp.MustCompile(`^https?://`)
)

// Initialize Sdk package.
func init() {
	new(sync.Once).Do(func() {
		// initialize config.
		Config = &configuration{}
		Config.initialize()
		// initialize consul.
		Consul = &consul{}
		Consul.initialize()
		// initialize pool.
		initPoolHttp()
		initPoolRpc()
	})
}
