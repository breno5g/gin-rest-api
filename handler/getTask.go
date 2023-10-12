package handler

import (
	"net/http"

	"github.com/breno5g/gin-rest-api/schemas"
	"github.com/gin-gonic/gin"
)

func GetTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}

	task := schemas.Task{}
	err := db.Where("id = ?", id).First(&task).Error
	if err != nil {
		logger.Errorf("error getting task: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error getting task")
		return
	}

	sendSuccess(ctx, "getTask", task)
}
