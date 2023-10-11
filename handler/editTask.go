package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditTask(ctx *gin.Context) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "Edit task",
	})
}
