package handler

import (
	"net/http"

	"github.com/breno5g/gin-rest-api/schemas"
	"github.com/gin-gonic/gin"
)

func ListTask(ctx *gin.Context) {
	tasks := []schemas.Task{}

	err := db.Find(&tasks).Error
	if err != nil {
		logger.Errorf("error getting tasks: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error getting tasks"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
