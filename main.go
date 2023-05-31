package main

import (
	"github.com/ReeVicente/gopportunities/config"
	"github.com/ReeVicente/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()

	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
