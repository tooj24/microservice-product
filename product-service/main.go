package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			InitDatabase,
			NewGinEngine,
			NewProductRepositoryImpl,
			NewProductServiceImpl,
			NewProductController,
		),
		fx.Invoke(RouteProvider),
	).Run()
}

func NewGinEngine() *gin.Engine {
	engine := gin.Default()
	return engine
}

func RouteProvider(engine *gin.Engine, controller *ProductController) {
	engine.GET(PRODUCT_URI, controller.FindAll)
	engine.POST(PRODUCT_URI, controller.Create)
	engine.PUT(PRODUCT_URI+"/:productId", controller.Update)
	engine.GET(PRODUCT_URI+"/:slug", controller.FindBySlug)
	engine.Run()
}
