# Variables
APP_NAME=github.com/aswindevs/kong_interview-assignment_1
MIGRATION_PATH=migrations
ATLAS_CMD=atlas

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean

# Docker commands
DOCKER_COMPOSE=docker-compose

# Build the application
build:
	$(GOBUILD) -o $(APP_NAME) ./cmd/app

# Run the application
run:
	$(GORUN) ./cmd/app/main.go

# Run migrations
migrate:
	$(GORUN) ./cmd/app/main.go -migrate

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(APP_NAME)

# Start Docker services
docker-up:
	$(DOCKER_COMPOSE) up -d

# Stop Docker services
docker-down:
	$(DOCKER_COMPOSE) down

# Initialize development environment
init: docker-up migrate

# Add new migration generation command
migration-new:
	$(ATLAS_CMD) migrate diff --env gorm $(name)

# Run migrations
migrate:
	$(GORUN) ./cmd/app/main.go -migrate

# Run all tests
test:
	$(GOCMD) test ./...



.PHONY: build run migrate clean docker-up docker-down init test