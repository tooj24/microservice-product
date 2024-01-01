package main

import "gorm.io/gorm"

type InventoryRepository interface {
	GetStock(skuCode string) Inventory
	AddStock(inventory Inventory) Inventory
}

type InventoryRepositoryImpl struct {
	DB *gorm.DB
}

// AddStock implements InventoryRepository.
func (repository *InventoryRepositoryImpl) AddStock(inventory Inventory) Inventory {
	stock := repository.GetStock(inventory.SkuCode)

	if stock.Id == 0 {
		repository.DB.Create(&inventory)
		return inventory
	}

	repository.DB.Model(&stock).Update("quantity", inventory.Quantity)
	return stock
}

// GetStock implements InventoryRepository.
func (repository *InventoryRepositoryImpl) GetStock(skuCode string) Inventory {
	var inventory Inventory
	repository.DB.Where("sku_code = ?", skuCode).Find(&inventory)
	return inventory
}

func NewInventoryRepositoryImpl(db *gorm.DB) InventoryRepository {
	db.AutoMigrate(&Inventory{})

	return &InventoryRepositoryImpl{DB: db}
}
