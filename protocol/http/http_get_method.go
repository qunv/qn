package http

import (
	"github.com/qunv/qn/protocol"
	"sync"
)

var GET = newHttpGetMethod()

type httpGetMethod struct{}

var getIns httpGetMethod
var gOnce sync.Once

func newHttpGetMethod() protocol.Method {
	gOnce.Do(func() {
		getIns = httpGetMethod{}
	})
	return getIns
}

func (h httpGetMethod) GetProtocolType() protocol.ProtocolType {
	return protocol.Http
}

func (h httpGetMethod) GetAction() string {
	return "GET"
}
