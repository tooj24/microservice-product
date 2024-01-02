package main

import "net/http"

type InventoryService interface {
	GetStock(slug string) Response
	AddStock(data InventoryRequest) Response
}

type InventoryServiceImpl struct {
	repository InventoryRepository
}

// AddStock implements InventoryService.
func (service *InventoryServiceImpl) AddStock(data InventoryRequest) Response {
	inventory := Inventory{
		SkuCode:  data.SkuCode,
		Quantity: data.Quantity,
	}

	inventory = service.repository.AddStock(inventory)

	return Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   inventory,
	}
}

// GetStock implements InventoryService.
func (service *InventoryServiceImpl) GetStock(skuCode string) Response {
	inventory := service.repository.GetStock(skuCode)

	if inventory.Id == 0 {
		return Response{
			Code:   http.StatusNotFound,
			Status: "inventory not found",
			Data:   nil,
		}
	}

	return Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   inventory,
	}
}

func NewInventoryServiceImpl(repository InventoryRepository) InventoryService {
	return &InventoryServiceImpl{repository: repository}
}
