package qn

type Registry interface {
	GetEndpoint() string
	GetMethod() Method
	GetTags() []string
}

type Foundation interface {
	Lookup(p ProtocolType) Registry
}

type Api interface {
	Foundation
	Handle(request Request) Response
}
