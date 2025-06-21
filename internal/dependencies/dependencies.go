package dependencies

import (
	"context"
	"fmt"

	"github.com/aswindevs/kong_interview-assignment_1/config"
	"github.com/aswindevs/kong_interview-assignment_1/internal/usecase"
	"github.com/aswindevs/kong_interview-assignment_1/internal/usecase/repo"
	"github.com/aswindevs/kong_interview-assignment_1/pkg/logger"
	"github.com/aswindevs/kong_interview-assignment_1/pkg/postgres"
	"github.com/aswindevs/kong_interview-assignment_1/pkg/server"

	"github.com/gin-gonic/gin"
)

// Service contains all dependencies for your application
type Dependencies struct {
	// Repositories
	UserRepo repo.UserRepo

	// Use cases
	UserUseCase           *usecase.UserUsecase
	AuthUseCase           *usecase.AuthUsecase
	ServiceCatalogUseCase *usecase.ServiceCatalogUsecase

	// Core dependencies
	DB     *postgres.Client
	Config *config.Config
	Logger logger.Interface
	Server *gin.Engine
}

// NewService initializes and returns a new Service instance
func NewDependencies(cfg *config.Config) (*Dependencies, error) {
	// Initialize logger with proper log level
	log, err := logger.New(cfg.Logger.Format, cfg.Logger.Level)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	// Initialize postgres client
	pgClient, err := postgres.NewClient(cfg, log)
	if err != nil {
		log.Error(context.Background(), "Failed to initialize postgres client", "error", err)
		return nil, err
	}

	// Initialize repositories
	userRepo := repo.NewUserRepoImpl(pgClient.DB)
	serviceRepo := repo.NewServiceRepoImpl(pgClient.DB)

	// Initialize use cases
	userUseCase := usecase.NewUserUsecase(userRepo)
	authUseCase := usecase.NewAuthUsecase(userRepo, &cfg.Auth)
	serviceCatalogUseCase := usecase.NewServiceCatalogUsecase(serviceRepo)

	return &Dependencies{
		// Assign all dependencies
		UserRepo:              userRepo,
		UserUseCase:           userUseCase,
		AuthUseCase:           authUseCase,
		ServiceCatalogUseCase: serviceCatalogUseCase,
		DB:                    pgClient,
		Config:                cfg,
		Logger:                log,
		Server:                server.New(log),
	}, nil
}
