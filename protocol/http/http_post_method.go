package http

import (
	"github.com/qunv/qn/protocol"
	"sync"
)

var POST = newHttpPostMethod()

type httpPostMethod struct{}

var postIns httpPostMethod
var postOnce sync.Once

func newHttpPostMethod() protocol.Method {
	postOnce.Do(func() {
		postIns = httpPostMethod{}
	})
	return postIns
}

func (h httpPostMethod) GetProtocolType() protocol.ProtocolType {
	return protocol.Http
}

func (h httpPostMethod) GetAction() string {
	return "POST"
}
