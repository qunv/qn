package apis

import (
	"fmt"
	"github.com/qunv/qn"
)

type PutSynonymApi struct {
	qn.Regs
}

func NewPutSynonymApi() qn.Api {
	http := qn.HTTP_PUT("/v1/:id").Tags("private", "public").New()
	return &PutSynonymApi{
		Regs: qn.Registers(http),
	}
}

func (s *PutSynonymApi) Handle(r qn.Request) qn.Response {
	id := r.GetContext().Param("id")
	fmt.Println("Id=", id)
	if id == "1" {
		return qn.ErrorResponse{
			Code:    1000,
			Message: "put error!",
		}
	}
	return qn.SuccessResponse{
		Payload: struct {
			Message string
		}{
			Message: "Put success",
		},
	}
}
