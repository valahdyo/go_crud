package migration

import (
	"go-crud/db"
	"go-crud/model/entity"
)

func Migrate() {
	db := db.InitDb()
	db.AutoMigrate(&entity.User{})
}
