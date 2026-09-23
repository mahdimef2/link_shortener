package router

import (
	"github.com/gin-gonic/gin"
	"url_shotener/internal/handler"
)

func SetupRouter(UrlHandler *handler.URLHandler) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.POST("/shorten", UrlHandler.Shorten)

	}
	r.GET("/:code", UrlHandler.Redirect)
	return r

}
