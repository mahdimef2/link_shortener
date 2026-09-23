package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	Port    string
	DBPath  string
	GinMode string
}

var Cfg AppConfig

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	Cfg = AppConfig{
		Port:    getEnv("PORT", "8080"),
		DBPath:  getEnv("DB_PATH", "urlshortener.db"),
		GinMode: getEnv("GIN_MODE", "debug"),
	}
}

// getEnv reads an env var, or returns a default if not set
func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
