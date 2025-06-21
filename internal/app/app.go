package app

import (
	"log"

	"github.com/aswindevs/kong_interview-assignment_1/config"
	"github.com/aswindevs/kong_interview-assignment_1/internal/controller"
	"github.com/aswindevs/kong_interview-assignment_1/internal/dependencies"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	svc, err := dependencies.NewDependencies(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}

	logger := svc.Logger
	controller.AddRouter(svc, cfg)

	logger.Info(nil, "Server starting", "port", cfg.HTTP.Port)

	if err := svc.Server.Run(":" + cfg.HTTP.Port); err != nil {
		logger.Fatal(nil, "Failed to start server", "error", err)
	}
}
