package postgres

import (
	"fmt"

	"github.com/aswindevs/kong_interview-assignment_1/config"
	"github.com/aswindevs/kong_interview-assignment_1/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	DB *gorm.DB
}

func NewClient(cfg *config.Config, logger logger.Interface) (*Client, error) {
	gormLogger := NewGormLogger(logger, getGormLogLevel(cfg.PG.Level))

	db, err := gorm.Open(postgres.Open(cfg.PG.URL), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.PG.PoolMax)
	sqlDB.SetMaxOpenConns(cfg.PG.PoolMax)

	return &Client{
		DB: db,
	}, nil
}
