package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/qunv/qn/example/gin/apis"
	"github.com/qunv/qn/example/gin/router"
	"go.uber.org/fx"
)

func All() fx.Option {
	return fx.Options(

		//init apis
		Apis(),
		//init gin
		fx.Provide(gin.New),

		//fx.Invoke(models.InitializeDatabase),
		fx.Invoke(router.RegisterGinRouters),
	)
}

func Apis() fx.Option {
	return fx.Options(
		ProvideIApi(apis.NewGetSynonymApi),
		ProvideIApi(apis.NewPutSynonymApi),
	)
}

func ProvideIApi(constructor interface{}) fx.Option {
	return fx.Provide(fx.Annotated{
		Group:  "nn_api",
		Target: constructor,
	})
}
