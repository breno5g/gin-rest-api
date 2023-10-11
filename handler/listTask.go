package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTask(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "List of tasks",
	})
}
