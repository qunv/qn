package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qunv/qn"
	"github.com/qunv/qn/protocol"
	"go.uber.org/fx"
)

type Strategy interface {
	Protocol() protocol.ProtocolType
	Handle(group *gin.RouterGroup, api qn.Api)
}

func ProvideRouterStrategy() fx.Option {
	return fx.Options(
		provideRouterStrategy(NewHttpStrategy),
		provideRouterStrategy(NewWsStrategy),
	)
}

func provideRouterStrategy(constructor interface{}) fx.Option {
	return fx.Provide(fx.Annotated{
		Group:  "nn_router_strategy",
		Target: constructor,
	})
}
