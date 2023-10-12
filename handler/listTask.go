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
		sendError(ctx, http.StatusInternalServerError, "error getting tasks")
		return
	}

	sendSuccess(ctx, "listTasks", tasks)
}
