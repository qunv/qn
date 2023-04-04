# qn

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)
[![Build Status](https://img.shields.io/travis/tamnd/httpclient/master.svg?style=flat-square)](https://travis-ci.org/qunv/qn)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/qunv/qn)

Simplified wrap router handler for [Go](http://www.golang.org).

[qn](https://github.com/qunv/qn) is designed to be the simplest way possible to make http, ws, event-bus requests in the
same gin handler
See example with [gin](https://github.com/gin-gonic/gin) and [uber fx](https://github.com/uber-go/fx)

Simple Handler:

```go
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
    return qn.SuccessResponse{
        Payload: struct {
        Message string
    }{
        Message: "GET Success",
    },
}
```

Provide it to fx

```go

func ProvideApis() fx.Option {
    return fx.Options(
        ProvideIApi(apis.NewGetSynonymApi),
        ProvideIApi(apis.NewPutSynonymApi),
		// more and more api here
    )
}

func ProvideIApi(constructor interface{}) fx.Option {
    return fx.Provide(fx.Annotated{
        Group:  "nn_api",
        Target: constructor,
    })
}
```

Then router must just do in just one logic

```go

type RegisterRouterIn struct {
    fx.In
    Engine         *gin.Engine
    Apis           []qn.Api   `group:"nn_api"`
    RouterStrategy []Strategy `group:"nn_router_strategy"`
}

func RegisterGinRouters(p RegisterRouterIn) {
    group := p.Engine.Group("/")
    for _, api := range p.Apis {
	for _, strategy := range p.RouterStrategy {
            strategy.Handle(group, api)
        }
    }
}
```

## Why I made it?

Because I'm tired to wire more handler with the same feature in other router handler, and a controller struct has many
functions that is used in gin router, and a file is so long.
And how can I reuse just one handler for own HTTP, WS and EventBus protocol?

HTTP

user_controller.go

```go
func (a *User) GetUser(c *gin.Context) {
	fmt.Println("get function")
}

func (a *User) DeleteUser(c *gin.Context) {
    fmt.Println("Delete function")
}

func (a *User) UpdateUser(c *gin.Context) {
    fmt.Println("Delete function")
}

//... more functions in here
```

router.go

```go
userController := &UserController{}
engine := gin.New()
group := engine.Group("/")
group.GET("/get/:id", userController.GetUser)
group.DELETE("/delete/:id", userController.DeleteUser)
group.PUT("/put/:id", userController.PutUser)

// ... more and more GET, PUT, POST here @@
```

## When to use it?

I have no idea

## Is it any good?

[May be](https://news.ycombinator.com/item?id=3067434).

## Install

```shell
go get github.com/qunv/qn
```

## Usage

```go
import "github.com/qunv/qn"
```

## Roadmap
- [X] HTTP, WS
- [ ] graphql, grpc, eventbus
- [ ] Handle tags
- [ ] Check duplicate protocol method 
- [ ] New http method
- [ ] Refactor 
- [ ] Testing
- 
## Contribute
- Fork repository
- Create a feature branch
- Open a new pull request
- Create an issue for bug report or feature request

## License
The MIT License (MIT). Please see [LICENSE](LICENSE) for more information.
