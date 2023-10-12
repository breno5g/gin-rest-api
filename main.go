package main

import (
	"github.com/breno5g/gin-rest-api/config"
	"github.com/breno5g/gin-rest-api/router"
)

func main() {
	logger := config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Init()
}
