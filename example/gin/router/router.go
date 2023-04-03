package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qunv/qn"
	"github.com/qunv/qn/example/gin/responses"
	"go.uber.org/fx"
)

type RegisterRouterIn struct {
	fx.In
	Engine *gin.Engine
	Apis   []qn.Api `group:"nn_api"`
}

func RegisterGinRouters(p RegisterRouterIn) {
	group := p.Engine.Group("/")

	for _, api := range p.Apis {
		httpApi := api.Lookup(qn.Http)
		group.Handle(httpApi.GetMethod().GetAction(), httpApi.GetEndpoint(), func(ctx *gin.Context) {
			resp := api.Handle(qn.NewDefaultRequest(ctx))
			success, ok := resp.(qn.SuccessResponse)
			if ok {
				responses.OfSucceed(ctx, 200, success.Payload)
				return
			}
			err, ok := resp.(qn.ErrorResponse)
			if ok {
				responses.OfError(ctx, 400, err.Message)
				return
			}
		})
	}
}
