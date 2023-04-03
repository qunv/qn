package qn

import (
	"github.com/gin-gonic/gin"
)

type Request interface {
	//TODO: init some thing here
	GetContext() *gin.Context
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
