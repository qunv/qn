package http

import (
	"github.com/qunv/qn/protocol"
	"sync"
)

var DELETE = newHttpDeleteMethod()

type httpDeleteMethod struct{}

var deleteIns httpDeleteMethod
var dOnce sync.Once

func newHttpDeleteMethod() protocol.Method {
	dOnce.Do(func() {
		deleteIns = httpDeleteMethod{}
	})
	return deleteIns
}

func (h httpDeleteMethod) GetProtocolType() protocol.ProtocolType {
	return protocol.Http
}

func (h httpDeleteMethod) GetAction() string {
	return "DELETE"
}
