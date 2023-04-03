package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qunv/qn/example/gin/bootstrap"
	"go.uber.org/fx"
	"log"
)

func main() {
	fx.New(bootstrap.All(), StartOpt()).Run()
}

func StartOpt() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle, engine *gin.Engine) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				port := 8888
				log.Printf("Application will be served at %d", port)
				go func() {
					if err := engine.Run(fmt.Sprintf(":%d", port)); err != nil {
						log.Fatalf("Start app got an error [%v]", err)
					}
				}()
				return nil
			},
		})
	})
}
