package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/qunv/qn"
	"github.com/qunv/qn/example/gin/responses"
	"github.com/qunv/qn/protocol"
	"go.uber.org/fx"
)

type RegisterRouterIn struct {
	fx.In
	Engine *gin.Engine
	Apis   []qn.Api `group:"nn_api"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func RegisterGinRouters(p RegisterRouterIn) {
	group := p.Engine.Group("/")

	for _, api := range p.Apis {
		httpApi := api.Lookup(protocol.Http)
		if httpApi != nil {
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

		wsApi := api.Lookup(protocol.Websocket)
		if wsApi != nil {
			group.GET(wsApi.GetEndpoint(), func(ctx *gin.Context) {
				//upgrade get request to websocket protocol
				ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer ws.Close()
				for {
					//Read Message from client
					mt, message, err := ws.ReadMessage()
					if err != nil {
						fmt.Println(err)
						break
					}
					//If client message is ping will return pong
					if string(message) == "ping" {
						message = []byte("pong")
					}
					//Response message to client
					err = ws.WriteMessage(mt, message)
					if err != nil {
						fmt.Println(err)
						break
					}
				}
			})
		}
	}
}
