package helloGin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GIN_API() {
	r := gin.Default()
	// r.GET("/docs/*any", ginSwagger)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 0.0.0.0:8080
}
