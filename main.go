package main

import (
	"ms-gin-gonic-rest-api-go/config"
	"ms-gin-gonic-rest-api-go/router"
)

var (
	logger *config.Logger
)

func main() {
	// init Logger
	logger = config.GetLogger("main")

	// init Config
	err := config.Init()

	if err != nil {
		// panic(err)
		logger.Errorf("config init error: %v", err)
		return
	}

	// init Router
	router.Initialize()
}
