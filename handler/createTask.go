package handler

import (
	"net/http"

	"github.com/breno5g/gin-rest-api/schemas"
	"github.com/gin-gonic/gin"
)

func CreateTask(ctx *gin.Context) {
	request := CreateTaskRequest{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		logger.Errorf("error binding request: %v", err)
		sendError(ctx, http.StatusBadRequest, "error binding request")
		return
	}

	newTask := schemas.Task{
		Title:       request.Title,
		Description: request.Description,
		Completed:   false,
	}

	err = db.Create(&newTask).Error
	if err != nil {
		logger.Errorf("error creating task: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error creating task")
		return
	}

	sendSuccess(ctx, "createTask", newTask)
}
