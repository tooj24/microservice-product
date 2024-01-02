package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", "root", "root", "product_service")

	println(dsn)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	return db
}
