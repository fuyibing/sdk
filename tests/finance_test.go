// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

package tests

import (
	"testing"

	"github.com/fuyibing/log/v2"

	"github.com/fuyibing/sdk/finance/bill"
)

func TestFinanceCreate(t *testing.T) {
	ctx := log.NewContext()
	res := bill.Create(ctx, `{"page":1,"limit":10,"pageData":"2021-03-06T02:39:46.263Z","status":null,"nzSelectedIndex":0,"scrollTopHeight":0,"expandForm":false,"insureLevel":"1"}`)

	if res.HasError() {
		t.Errorf("Error: %v.", res.GetError())
		return
	}
	t.Logf("Response: %v.", res.GetBody())

}
