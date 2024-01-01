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

func RouteProvider(engine *gin.Engine, controller *InventoryController) {
	engine.GET(INVENTORY_URI, controller.GetStock)
	engine.POST(INVENTORY_URI, controller.AddStock)
	engine.Run(":8081")
}
