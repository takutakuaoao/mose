package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/health-checker", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "success",
		})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":8080")
}
