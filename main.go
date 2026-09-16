package main

import (
	"log"
	"url_shotener/config"
	"url_shotener/handler"
	"url_shotener/models"
	"url_shotener/repository"
	"url_shotener/router"
	"url_shotener/service"
)

const port = ":8080"

func main() {
	config.ConnectDatabase()
	err := config.DB.AutoMigrate(&models.Url{})
	if err != nil {
		return
	}
	urlRepository := repository.NewUrlRepository(config.DB)
	urlService := service.NewUrlService(urlRepository)
	urlHandler := handler.NewUrlHandler(urlService)

	r := router.SetupRouter(urlHandler)
	log.Printf("🚀 Service is up and running on port %s\n", port)
	if err := r.Run(port); err != nil {
		log.Fatal("failed to start server:", err)
	}

}
