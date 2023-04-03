package qn

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Request interface {
	//TODO: init some thing here
	GetContext() *gin.Context
	GetBody() ([]byte, error)
}

type DefaultRequest struct {
	ctx *gin.Context
}

func NewDefaultRequest(ctx *gin.Context) Request {
	return &DefaultRequest{
		ctx: ctx,
	}
}

func (d DefaultRequest) GetContext() *gin.Context {
	return d.ctx
}

func (d DefaultRequest) GetBody() ([]byte, error) {
	body := d.ctx.Request.Body
	return ioutil.ReadAll(body)
}
