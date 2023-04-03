package qn

import (
	"github.com/gin-gonic/gin"
	"github.com/qunv/qn/protocol"
	"github.com/qunv/qn/protocol/eb"
	http2 "github.com/qunv/qn/protocol/http"
	"github.com/qunv/qn/protocol/ws"
)

type regBuilder interface {
	Tags(...string) regBuilder
	MiddleWare(...gin.HandlerFunc) regBuilder
	New() Registry
}

func WithEndpoint(endpoint string) _withRegFunc {
	return func(reg *reg) {
		reg.endpoint = endpoint
	}
}

func WithMethod(method protocol.Method) _withRegFunc {
	return func(reg *reg) {
		reg.method = method
	}
}

func WithTags(tags ...string) _withRegFunc {
	return func(reg *reg) {
		reg.tags = tags
	}
}

func WithMiddleWare(middlewares ...gin.HandlerFunc) _withRegFunc {
	return func(reg *reg) {
		reg.middlewares = middlewares
	}
}

func Register(fns ..._withRegFunc) Registry {
	reg := new(reg)
	for _, with := range fns {
		with(reg)
	}
	return reg
}

// HTTP_GET is a builder function that apply HTTP protocol and method GET
func HTTP_GET(endpoint string) regBuilder {
	return &reg{
		method:   http2.GET,
		endpoint: endpoint,
	}
}

func HTTP_POST(endpoint string) regBuilder {
	return &reg{
		method:   http2.POST,
		endpoint: endpoint,
	}
}

func HTTP_PUT(endpoint string) regBuilder {
	return &reg{
		method:   http2.PUT,
		endpoint: endpoint,
	}
}

func HTTP_DELETE(endpoint string) regBuilder {
	return &reg{
		method:   http2.DELETE,
		endpoint: endpoint,
	}
}

func WS(endpoint string) regBuilder {
	return &reg{
		method:   ws.WS,
		endpoint: endpoint,
	}
}

func EVENT_BUS(endpoint string) regBuilder {
	return &reg{
		method:   eb.EvenBus,
		endpoint: endpoint,
	}
}
