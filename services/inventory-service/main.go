package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"services/shared/eureka"
)

func main() {
	fx.New(
		fx.Provide(
			InitDatabase,
			NewGinEngine,
			eureka.NewEurekaServiceImpl,
			NewInventoryRepositoryImpl,
			NewInventoryServiceImpl,
			NewInventoryController,
		),
		fx.Invoke(RouteProvider),
	).Run()
}

func NewGinEngine() *gin.Engine {
	engine := gin.Default()
	return engine
}

func RouteProvider(engine *gin.Engine, controller *InventoryController, eurekaServer eureka.EurekaService) {
	port := eureka.RandomPort()

	engine.GET(INVENTORY_URI, controller.GetStock)
	engine.POST(INVENTORY_URI, controller.AddStock)

	eurekaServer.Register("inventory-service", "localhost", port)
	engine.Run(fmt.Sprintf(":%d", port))
}
