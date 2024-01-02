package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"services/shared/eureka"
)

const serviceName = "product-service"

func main() {
	fx.New(
		fx.Provide(
			InitDatabase,
			NewGinEngine,
			NewProductRepositoryImpl,
			NewProductServiceImpl,
			NewProductController,
			eureka.NewEurekaServiceImpl,
		),
		fx.Invoke(RouteProvider),
	).Run()
}

func NewGinEngine() *gin.Engine {
	engine := gin.Default()
	return engine
}

func RouteProvider(engine *gin.Engine, controller *ProductController, eurekaService eureka.EurekaService) {
	port := eureka.RandomPort()

	engine.GET(PRODUCT_URI, controller.FindAll)
	engine.POST(PRODUCT_URI, controller.Create)
	engine.PUT(PRODUCT_URI+"/:productId", controller.Update)
	engine.GET(PRODUCT_URI+"/:slug", controller.FindBySlug)

	eurekaService.Register(serviceName, "localhost", port)

	engine.Run(fmt.Sprintf(":%d", port))
}
