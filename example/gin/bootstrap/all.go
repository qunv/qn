package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/qunv/qn/example/gin/controllers"
	"github.com/qunv/qn/example/gin/router"
	"go.uber.org/fx"
)

func All() fx.Option {
	return fx.Options(

		//init controllers
		Apis(),
		//init gin
		fx.Provide(gin.New),

		//fx.Invoke(models.InitializeDatabase),
		fx.Invoke(router.RegisterGinRouters),
	)
}

func Apis() fx.Option {
	return fx.Options(
		ProvideIApi(controllers.NewGetSynonymApi),
	)
}

func ProvideIApi(constructor interface{}) fx.Option {
	return fx.Provide(fx.Annotated{
		Group:  "nn_api",
		Target: constructor,
	})
}
