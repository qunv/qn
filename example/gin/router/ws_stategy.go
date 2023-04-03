package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/qunv/qn"
	"github.com/qunv/qn/protocol"
	"log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WsStrategy struct {
}

func NewWsStrategy() Strategy {
	return &WsStrategy{}
}

func (w WsStrategy) Protocol() protocol.ProtocolType {
	return protocol.Websocket
}

func (w WsStrategy) Handle(group *gin.RouterGroup, api qn.Api) {
	wsApi := api.Lookup(protocol.Websocket)
	if wsApi == nil {
		return
	}
	group.GET(wsApi.GetEndpoint(), func(ctx *gin.Context) {
		ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer ws.Close()
		for {
			//Read Message from client
			mt, message, err := ws.ReadMessage()
			if err != nil {
				log.Println(err)
				break
			}
			resp := api.Handle(newWsRequest(ctx, message))
			b, _ := json.Marshal(resp)
			//Response message to client
			err = ws.WriteMessage(mt, b)
			if err != nil {
				log.Println(err)
				break
			}
		}
	})
}

type wsRequest struct {
	ctx  *gin.Context
	body []byte
}

func newWsRequest(ctx *gin.Context, body []byte) qn.Request {
	return &wsRequest{
		ctx:  ctx,
		body: body,
	}
}

func (w wsRequest) GetContext() *gin.Context {
	return w.ctx
}

func (w wsRequest) GetBody() ([]byte, error) {
	return w.body, nil
}
