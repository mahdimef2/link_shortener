package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"url_shotener/internal/config"
	"url_shotener/internal/handler"
	"url_shotener/internal/models"
	"url_shotener/internal/repository"
	"url_shotener/internal/router"
	"url_shotener/internal/service"
)

const port = ":8080"

func main() {
	config.LoadConfig()
	gin.SetMode(config.Cfg.GinMode)
	config.ConnectDatabase()
	err := config.DB.AutoMigrate(&models.Url{})
	if err != nil {
		return
	}
	urlRepository := repository.NewUrlRepository(config.DB)
	urlService := service.NewUrlService(urlRepository)
	urlHandler := handler.NewUrlHandler(urlService)

	r := router.SetupRouter(urlHandler)
	log.Printf("🚀 Service is up and running on port %s\n", config.Cfg.Port)
	if err := r.Run(port); err != nil {
		log.Fatal("failed to start server:", err)
	}

}
