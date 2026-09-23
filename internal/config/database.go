package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open(Cfg.DBPath), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = database
	log.Println("Database connected successfully")
}
