package main

import (
	"github.com/mattos-smart/Estudos-Go/config"
	"github.com/mattos-smart/Estudos-Go/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// INITIALIZE CONFIGS
	err := config.Init()
	if err != nil {
		// panic(err) // Panic mata a aplicação
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// INITIALIZE ROUTER
	router.Initialize()

}
