package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTask(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "Get Task",
	})
}
