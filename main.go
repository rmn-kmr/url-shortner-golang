package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rmn-kmr/go-url-shortener/handler"
	"github.com/rmn-kmr/go-url-shortener/store"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortner!",
		})
	})

	r.POST("/create-short-url", func(ctx *gin.Context) {
		handler.CreateShortUrl(ctx)
	})

	r.GET("/:shortUrl", func(ctx *gin.Context) {
		handler.HandleShortUrlRedirect(ctx)
	})

	store.IntializeStore()

	err := r.Run(":9988")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
