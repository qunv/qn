package eb

import (
	"github.com/qunv/qn/protocol"
	"sync"
)

var EvenBus = newEventbusMethod()

type eventbusMethod struct {
	protocol.Method
}

var instance eventbusMethod
var once sync.Once

func newEventbusMethod() protocol.Method {
	once.Do(func() {
		instance = eventbusMethod{}
	})
	return instance
}

func (h eventbusMethod) GetProtocolType() protocol.ProtocolType {
	return protocol.EvenBus
}
