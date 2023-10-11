package router

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.Run(":3001")
}
