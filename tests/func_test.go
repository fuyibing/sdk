// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

package tests

import (
	"testing"

	"github.com/fuyibing/sdk"
)

// func TestFunc(t *testing.T) {
// 	ctx := log.NewContext()
// 	wg := new(sync.WaitGroup)
//
// 	url := "http://wxapi.turboradio.cn/v/user/info"
// 	// url := "http://www.facebook.com/index"
//
//
// 	for i := 0; i < 1; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			req := sdk.NewRequest(ctx, url)
// 			req.SetMethod(http.MethodPost)
// 			req.SetBody(`{"key":"value"}`)
// 			req.SetTimeout(3)
// 			res := req.Run()
// 			if res.HasError() {
// 				t.Logf("Duration: [%f] %v.", res.GetDuration(), res.GetError())
// 			} else {
// 				t.Logf("Duration: [%f] Succeed.", res.GetDuration())
// 			}
// 		}()
// 	}
// 	wg.Wait()
//
// 	// sdk.Get(ctx, "http://wxapi.turboradio.cn/v/user/info")
// }
//
// func TestFuncGet(t *testing.T) {
// 	ctx := log.NewContext()
// 	sdk.Get(ctx, "http://wxapi.turboradio.cn/v/user/info")
// }

func TestUrlMap(t *testing.T) {
	sdk.NewRequest().
		SetUrl("http://localhost:18000/").
		SetMethod("POST").
		SetBody(map[string]interface{}{"key": "value", "num": 1}).
		Run(nil)
}
func TestUrlString(t *testing.T) {}
func TestUrlByte(t *testing.T)   {}
func TestUrlStruct(t *testing.T) {

	v := &struct {
		Key   string
		Value int
	}{
		Key:   "key",
		Value: 1001,
	}

	// sdk.NewRequest().
	// 	SetUrl("http://localhost:18000/").
	// 	SetMethod("POST").
	// 	SetBody(v).
	// 	Run(nil)

	sdk.Post(nil, "http://localhost:18000", v)
}
