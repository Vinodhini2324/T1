package models

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("mysql", "user:password@tcp(localhost:3306)/test")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
