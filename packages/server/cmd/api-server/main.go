package main

import (
	"log"
	"my-codelens-app/internal/common"
	"my-codelens-app/internal/common/logger"
	"my-codelens-app/internal/config"
)

func main() {
	//Application starts
	//Create root context
	// ctx  → used by different parts of application
	// cancel → function that cancels/stops that context
	ctx, cancel := common.GlobalContext()
	defer cancel()

	// This load return .env config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	//Initialize() takes your application's config, configures how Zap should behave, builds the actual logger, and stores it so the rest of the application can use it.
	if err := logger.Initialize(*cfg); err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	logger.Info("Configuration loaded Successfully")
}
