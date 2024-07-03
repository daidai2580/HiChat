package main

import (
	"HiChat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Li244088167@tcp(120.24.168.49:3306)/hinet?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Part{})
	if err != nil {
		panic(err)
	}
}
