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
	"regexp"
	"sync"
)

const (
	CacheLifetime float64 = 3.0
	Timeout               = 2
	RequestTime           = 25
	ServiceScheme         = "http"
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
