package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("urlshortener.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = database
	log.Println("Database connected successfully")
}
