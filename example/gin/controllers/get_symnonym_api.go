package controllers

import (
	"fmt"
	"github.com/qunv/qn"
)

type GetSynonymApi struct {
	qn.Regs
}

func NewGetSynonymApi() qn.Api {
	//http := qn.Register(qn.WithMethod(qn.HttpGet), qn.WithEndpoint("/v1/:id"))
	ws := qn.Register(qn.WithMethod(qn.WS), qn.WithEndpoint("/v1/:id"))
	htp := qn.HTTP_GET("/v1/:id").Tags("private", "public").New()
	return &GetSynonymApi{
		Regs: qn.Registers(htp, ws),
	}
}

func (s *GetSynonymApi) Handle(r qn.Request) qn.Response {
	id := r.GetContext().Param("id")
	fmt.Println("Id=", id)
	if id == "1" {
		return qn.ErrorResponse{
			Code:    1000,
			Message: "error nay",
		}
	}
	return qn.SuccessResponse{
		Payload: struct {
			Message string
		}{
			Message: "success",
		},
	}
}
