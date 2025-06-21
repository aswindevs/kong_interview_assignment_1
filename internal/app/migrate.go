package app

import (
	"errors"
	"fmt"

	"github.com/aswindevs/kong_interview-assignment_1/config"

	// "github.com/aswindevs/kong_interview-assignment_1/pkg/logger"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrate() {
	config, err := config.NewConfig()
	// logger := logger.New(config.Logger.Level)
	if err != nil {
		fmt.Println("Error reading config:", err)
		panic(err)
	}
	dbUrl := config.PG.URL
	m, err := migrate.New("file://migrations", dbUrl)
	if err != nil {
		fmt.Println("Migrate: postgres error:", err)

	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("Migrate: error:", err)
	}
}
