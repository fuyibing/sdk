// author: wsfuyibing <websearch@163.com>
// date: 2021-03-05

package sdk

import (
	"sync"
	"sync/atomic"
)

var (
	poolRpc      *sync.Pool
	poolRpcIndex uint64
)

type ClientRpc struct {
	index uint64
}

// Acquire rpc sdk client.
func NewRpc() *ClientRpc {
	o := poolRpc.Get().(*ClientRpc)
	return o
}

// 初始化RPC池.
func initPoolRpc() {
	poolRpc = &sync.Pool{New: func() interface{} {
		return &ClientRpc{
			index: atomic.AddUint64(&poolRpcIndex, 1),
		}
	}}
}
