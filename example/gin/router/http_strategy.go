package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qunv/qn"
	"github.com/qunv/qn/example/gin/responses"
	"github.com/qunv/qn/protocol"
)

type HttpStrategy struct {
}

func NewHttpStrategy() Strategy {
	return &HttpStrategy{}
}

func (h HttpStrategy) Protocol() protocol.ProtocolType {
	return protocol.Http
}

func (h HttpStrategy) Handle(group *gin.RouterGroup, api qn.Api) {
	httpApi := api.Lookup(protocol.Http)
	if httpApi == nil {
		return
	}
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
