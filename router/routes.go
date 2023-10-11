package router

import (
	"github.com/breno5g/gin-rest-api/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/task", handler.ListTask)
		v1.GET("/task/:id", handler.GetTask)
		v1.POST("/task", handler.CreateTask)
		v1.PUT("/task/:id", handler.EditTask)
		v1.DELETE("/task/:id", handler.DeleteTask)
	}
}
