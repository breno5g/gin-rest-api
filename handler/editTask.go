package handler

import (
	"net/http"

	"github.com/breno5g/gin-rest-api/schemas"
	"github.com/gin-gonic/gin"
)

func EditTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "Param id is required")
		return
	}

	request := UpdateTaskRequest{}
	err := ctx.BindJSON(&request)
	if err != nil {
		logger.Errorf("error binding request: %v", err)
		sendError(ctx, http.StatusBadRequest, "error binding request")
		return
	}

	var task schemas.Task

	err = db.First(&task, id).Error
	if err != nil {
		sendError(ctx, http.StatusNotFound, "Task not found")
		return
	}

	task.Title = request.Title
	task.Description = request.Description
	task.Completed = request.Completed

	err = db.Save(&task).Error
	if err != nil {
		logger.Errorf("error updating task: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error updating task")
		return
	}

	sendSuccess(ctx, "editTask", task)
}
