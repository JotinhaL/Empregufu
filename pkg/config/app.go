package config

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	dsn := "root:JotaL1213@@tcp(localhost:3306)/empregUFU?charset=utf8mb4&parseTime=True&loc=Local"
  	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})


	if err != nil {
		panic("failed to connect database")
	}

	db = d
	fmt.Println("Database connection established successfully")
	
}

func GetDB() *gorm.DB {
	return db
}

