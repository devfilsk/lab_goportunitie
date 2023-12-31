package main

import (
	"goportunitie/config"
	"goportunitie/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize configs
	err := config.Init()

	if err != nil {
		logger.Errf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
