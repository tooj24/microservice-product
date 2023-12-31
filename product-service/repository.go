package main

import (
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() []Product
	FindBySlug(slug string) Product
	Create(product Product) Product
	Update(product Product) Product
}

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

// FindBySlug implements ProductRepository.
func (repository *ProductRepositoryImpl) FindBySlug(slug string) Product {
	var product Product
	result := repository.DB.Where("sku_code = ?", slug).Find(&product)
	ErrorPanic(result.Error)

	return product
}

// Update implements ProductRepository.
func (repository *ProductRepositoryImpl) Update(product Product) Product {
	result := repository.DB.Model(&product).Where("id = ?", product.Id).Updates(product)
	ErrorPanic(result.Error)
	return product
}

// FindAll implements ProductRepository.
func (repository *ProductRepositoryImpl) FindAll() []Product {
	var results []Product
	repository.DB.Find(&results)
	return results
}

// Create implements ProductRepository.
func (repository *ProductRepositoryImpl) Create(product Product) Product {
	result := repository.DB.Create(&product)
	ErrorPanic(result.Error)
	return product
}

func NewProductRepositoryImpl(db *gorm.DB) ProductRepository {
	db.AutoMigrate(&Product{})

	return &ProductRepositoryImpl{
		DB: db,
	}
}
