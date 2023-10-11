package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/task", func(ctx *gin.Context) {
			ctx.Header("Content-type", "application/json")
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
				"data":    "List of tasks",
			})
		})
		v1.GET("/task/:id", func(ctx *gin.Context) {
			ctx.Header("Content-type", "application/json")
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
				"data":    "Get Task",
			})
		})
		v1.POST("/task", func(ctx *gin.Context) {
			ctx.Header("Content-type", "application/json")
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
				"data":    "Create task",
			})
		})
		v1.PUT("/task/:id", func(ctx *gin.Context) {
			ctx.Header("content-type", "application/json")
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
				"data":    "Edit task",
			})
		})
		v1.DELETE("/task/:id", func(ctx *gin.Context) {
			ctx.Header("content-type", "application/json")
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
				"data":    "Delete Task",
			})
		})
	}
}
