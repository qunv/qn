package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qunv/qn"
	"go.uber.org/fx"
)

type RegisterRouterIn struct {
	fx.In
	Engine         *gin.Engine
	Apis           []qn.Api   `group:"nn_api"`
	RouterStrategy []Strategy `group:"nn_router_strategy"`
}

func RegisterGinRouters(p RegisterRouterIn) {
	group := p.Engine.Group("/")
	for _, api := range p.Apis {
		for _, strategy := range p.RouterStrategy {
			strategy.Handle(group, api)
		}
	}
}
