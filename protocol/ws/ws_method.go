package ws

import (
	"github.com/qunv/qn/protocol"
	"sync"
)

var WS = newWsMethod()

type wsMethod struct {
	protocol.Method
}

var instance wsMethod
var once sync.Once

func newWsMethod() protocol.Method {
	once.Do(func() {
		instance = wsMethod{}
	})
	return instance
}

func (h wsMethod) GetProtocolType() protocol.ProtocolType {
	return protocol.Websocket
}
