package main

import (
	"log"
	"my-codelens-app/internal/common/logger"
	"my-codelens-app/internal/config"
	"my-codelens-app/internal/routes"

	commonMiddleware "my-codelens-app/internal/common/middleware"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"go.uber.org/zap"
)

func main() {
	//Application starts
	//Create root context
	// ctx  → used by different parts of application
	// cancel → function that cancels/stops that context
	// ctx, cancel := common.GlobalContext()
	// defer cancel()

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

	e := echo.New()

	allowedOrigins := []string{cfg.FrontendOrigin}
	if cfg.AppEnv == config.Development {
		allowedOrigins = append(allowedOrigins, "http://localhost:3000")
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           86400, // 24 hours
	}))

	// Add Zap middleware
	e.Use(commonMiddleware.ZapLogger())
	e.Use(commonMiddleware.Recovery())

	routes.RegisterRoutes(e)

	logger.Info("Starting HTTP server", zap.String("port", cfg.Port))

	if err := e.Start(":" + cfg.Port); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
