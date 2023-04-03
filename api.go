package qn

import "github.com/qunv/qn/protocol"

type Registry interface {
	GetEndpoint() string
	GetMethod() protocol.Method
	GetTags() []string
}

type Foundation interface {
	Lookup(p protocol.ProtocolType) Registry
}

type Api interface {
	Foundation
	Handle(request Request) Response
}
