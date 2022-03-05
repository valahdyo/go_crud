package main

import (
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	dsn := "root@tcp(127.0.0.1:3306)/go-crud?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to Connect DB")

	}
	return db
}

func Migrate() {
	db := InitDb()
	db.AutoMigrate(&User{})
}
