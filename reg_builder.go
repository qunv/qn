package qn

import "github.com/gin-gonic/gin"

type regBuilder interface {
	Tags(...string) regBuilder
	MiddleWare(...gin.HandlerFunc) regBuilder
	New() Registry
}

func HTTP_GET(endpoint string) regBuilder {
	return &reg{
		method:   HttpGet,
		endpoint: endpoint,
	}
}

func HTTP_POST(endpoint string) regBuilder {
	return &reg{
		method:   HttpPost,
		endpoint: endpoint,
	}
}

func HTTP_PUT(endpoint string) regBuilder {
	return &reg{
		method:   HttpPut,
		endpoint: endpoint,
	}
}

func HTTP_DELETE(endpoint string) regBuilder {
	return &reg{
		method:   HttpDelete,
		endpoint: endpoint,
	}
}
