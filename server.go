package main

import (
	"go-vs-js-json-bench/generate"
	"gopkg.in/gin-gonic/gin.v1"
	"os"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5001"
	}
	rooms := generate.Response(20)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, rooms)
	})

	r.Run(":" + port)
}
