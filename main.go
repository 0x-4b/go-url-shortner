package main

import (
	"fmt"

	"github.com/0x-4b/go-url-shortner/handler"
	"github.com/0x-4b/go-url-shortner/store"
	"github.com/gin-gonic/gin"
)

func main() {
	store.InitializeStore()

	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("static/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start web server - Error: %v", err))
	}
}
