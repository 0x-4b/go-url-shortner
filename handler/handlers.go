package handler

import (
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	// Implementation later
}

func HandleShortUrlRedirect(c *gin.Context) {
	// Implementation later
}
