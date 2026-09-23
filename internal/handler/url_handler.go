package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url_shotener/internal/service"
)

type URLHandler struct {
	service service.URLService
}

func NewUrlHandler(service service.URLService) *URLHandler {
	return &URLHandler{service: service}
}

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

type ShortenResponse struct {
	Code     string `json:"code"`
	ShortUrl string `json:"short_url"`
}

func (h *URLHandler) Shorten(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.service.ShortenURL(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ShortenResponse{
		Code:     result.Code,
		ShortUrl: "http://localhost:8080/" + result.Code,
	})
}
func (h *URLHandler) Redirect(c *gin.Context) {
	code := c.Param("code")
	result, err := h.service.GetOriginalURL(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, result.OriginalURL)
}
