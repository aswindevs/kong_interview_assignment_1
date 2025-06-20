# Internal Developer Platform

A self-service platform that enables developers to manage DevOps tasks independently. This platform provides automated workflows for infrastructure provisioning, deployments, and other operational activities, reducing dependencies on DevOps teams while ensuring best practices.

## Environment Configuration

The application uses a combination of configuration files and environment variables:

1. **Config File**: Primary configuration is stored in `config/config.yaml`
2. **Environment Variables**: Can override config file values using the following variables:
   - `ENVIRONMENT`: Application environment (dev/prod)
   - `HTTP_PORT`: Server port
   - `LOG_LEVEL`: Logging level
   - `PG_URL`: PostgreSQL connection string
   - `PG_POOL_MAX`: Maximum database connections
   - `AUTH_SECRET_KEY`: JWT secret key
   - `AUTH_TOKEN_EXPIRATION_TIME`: JWT token expiration in seconds

Environment variables take precedence over config file values.

## Getting Started

### Prerequisites

- Go 1.23 or later
- Docker and Docker Compose
- Make

### Setup and Running

1. Start the PostgreSQL database:

```bash
make docker-up
```
2. To create a new migration:

```bash
make migration-new name=your_migration_name
```

This will generate a new migration file in the migrations directory.

3. Initialize the database and run migrations:

```bash
make migrate
```

4. Run the application:

```bash
make run
```

For a complete initialization of the development environment:

```bash
make init
```


### Available Make Commands

- `make build`: Build the application
- `make run`: Run the application
- `make migrate`: Run database migrations
- `make docker-up`: Start Docker services
- `make docker-down`: Stop Docker services
- `make init`: Initialize development environment
- `make clean`: Clean build files
- `make test`: Run tests

### API Endpoints

- Health Check: `GET /health`
- API Base URL: `/v1`
  - Authentication: `/v1/auth`
  - Users: `/v1/users`

## Database Migrations

Migrations are managed using golang-migrate and are located in the `migrations` directory. They are automatically applied when running `make migrate` or `make init`.

## Development

The application follows a clean architecture pattern with the following structure:
- `cmd/`: Application entrypoints
- `config/`: Configuration management
- `internal/`: Internal application code
- `pkg/`: Shared packages
- `migrations/`: Database migrations# kong_interview-assignment_1
