package main

import "github.com/gin-gonic/gin"

const INVENTORY_URI = "/api/inventory"

type InventoryController struct {
	inventoryService InventoryService
}

func NewInventoryController(inventoryService InventoryService) *InventoryController {
	return &InventoryController{
		inventoryService: inventoryService,
	}
}

func (controller *InventoryController) GetStock(ctx *gin.Context) {
	skuCode := ctx.Query("sku_code")

	response := controller.inventoryService.GetStock(skuCode)
	ctx.JSON(response.Code, response)
}

func (controller *InventoryController) AddStock(ctx *gin.Context) {
	request := InventoryRequest{}
	err := ctx.ShouldBindJSON(&request)
	ErrorPanic(err)

	response := controller.inventoryService.AddStock(request)
	ctx.JSON(response.Code, response)
}
