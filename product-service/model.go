package main

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type Product struct {
	Id          int `gorm:"primaryKey"`
	SkuCode     string
	Name        string
	Description string
	Price       float32 `gorm:"type:decimal(10,2);"`
}

type ProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type ProductResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
