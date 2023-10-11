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

func CreateTask(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "Create task",
	})
}

func GetTask(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "Get Task",
	})
}

func EditTask(ctx *gin.Context) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "Edit task",
	})
}

func DeleteTask(ctx *gin.Context) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "Delete Task",
	})
}
