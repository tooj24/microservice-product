package main

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type Inventory struct {
	Id       int `gorm:"primaryKey"`
	SkuCode  string
	Quantity int
}

type InventoryRequest struct {
	SkuCode  string `json:"sku_code"`
	Quantity int    `json:"quantity"`
}
