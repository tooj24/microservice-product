package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const PRODUCT_URI = "/api/product"

type ProductController struct {
	productService ProductService
}

func (controller *ProductController) FindAll(ctx *gin.Context) {
	products := controller.productService.FindAll()

	response := Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   products,
	}
	ctx.JSON(http.StatusOK, response)
}

func (controller *ProductController) Create(ctx *gin.Context) {
	request := ProductRequest{}
	err := ctx.ShouldBindJSON(&request)
	ErrorPanic(err)

	response := controller.productService.Create(request)

	ctx.JSON(response.Code, response)
}

func (controller *ProductController) Update(ctx *gin.Context) {
	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	ErrorPanic(err)

	request := ProductRequest{}
	err = ctx.ShouldBindJSON(&request)
	ErrorPanic(err)

	response := controller.productService.Update(id, request)

	ctx.JSON(response.Code, response)
}

func (controller *ProductController) FindBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	response := controller.productService.FindBySlug(slug)

	ctx.JSON(response.Code, response)
}

func NewProductController(productService ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}
