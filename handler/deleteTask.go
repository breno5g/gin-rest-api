package handler

import (
	"net/http"

	"github.com/breno5g/gin-rest-api/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "Param id is required")
		return
	}

	task := db.First(&schemas.Task{}, id).Error

	if task != nil {
		sendError(ctx, http.StatusNotFound, "Task not found")
		return
	}

	err := db.Delete(&schemas.Task{}, id).Error
	if err != nil {
		logger.Errorf("error deleting task: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error deleting task")
		return
	}

	sendSuccess(ctx, "deleteTask", nil)
}
