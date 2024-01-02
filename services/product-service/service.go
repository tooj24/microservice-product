package main

import "net/http"

type ProductService interface {
	FindAll() []Product
	FindBySlug(slug string) Response
	Create(data ProductRequest) Response
	Update(id int, data ProductRequest) Response
}

type ProductServiceImpl struct {
	repository ProductRepository
}

// FindBySlug implements ProductService.
func (service *ProductServiceImpl) FindBySlug(slug string) Response {
	var response Response
	var statusCode int

	product := service.repository.FindBySlug(slug)

	if product.Id != 0 {
		statusCode = http.StatusOK
		response = Response{
			Code:   statusCode,
			Status: "ok",
			Data:   product,
		}
	} else {
		statusCode = http.StatusNotFound
		response = Response{
			Code:   statusCode,
			Status: "product not found",
			Data:   nil,
		}
	}
	return response
}

// Update implements ProductService.
func (service *ProductServiceImpl) Update(id int, data ProductRequest) Response {
	product := Product{
		Id:          id,
		SkuCode:     Slugify(data.Name),
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
	}

	return Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   service.repository.Update(product),
	}
}

// FindAll implements ProductService.
func (service *ProductServiceImpl) FindAll() []Product {
	return service.repository.FindAll()
}

// Create implements ProductService.
func (service *ProductServiceImpl) Create(data ProductRequest) Response {
	product := Product{
		SkuCode:     Slugify(data.Name),
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
	}

	return Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   service.repository.Create(product),
	}
}

func NewProductServiceImpl(repository ProductRepository) ProductService {
	return &ProductServiceImpl{repository: repository}
}
