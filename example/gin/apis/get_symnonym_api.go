package apis

import (
	"fmt"
	"github.com/qunv/qn"
)

type GetSynonymApi struct {
	qn.Regs
}

func NewGetSynonymApi() qn.Api {
	ws := qn.WS("/ws/:id").New()
	http := qn.HTTP_GET("/v1/:id").Tags("private", "public").New()
	return &GetSynonymApi{
		Regs: qn.Registers(http, ws),
	}
}

func (s *GetSynonymApi) Handle(r qn.Request) qn.Response {
	id := r.GetContext().Param("id")
	fmt.Println("Id=", id)
	if id == "1" {
		return qn.ErrorResponse{
			Code:    1000,
			Message: "GET error!!!",
		}
	}
	return qn.SuccessResponse{
		Payload: struct {
			Message string
		}{
			Message: "GET Success",
		},
	}
}
