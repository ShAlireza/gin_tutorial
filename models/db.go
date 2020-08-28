package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("postgres", "username=chitchat password=chitchat dbname=gin_tutorial "+
		"sslmode=disable")

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
