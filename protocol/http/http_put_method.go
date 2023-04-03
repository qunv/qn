package http

import (
	"github.com/qunv/qn/protocol"
	"sync"
)

var PUT = newHttpPutMethod()

type httpPutMethod struct{}

var putIns httpPutMethod
var putOnce sync.Once

func newHttpPutMethod() protocol.Method {
	putOnce.Do(func() {
		putIns = httpPutMethod{}
	})
	return putIns
}

func (h httpPutMethod) GetProtocolType() protocol.ProtocolType {
	return protocol.Http
}

func (h httpPutMethod) GetAction() string {
	return "PUT"
}
