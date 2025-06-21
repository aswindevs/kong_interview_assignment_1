# Kong Interview Assignment

This project is a service catalog API developed for a Kong interview assignment. It is built using a boilerplate based on the [Go Clean Template](https://github.com/evrone/go-clean-template), emphasizing a clean, maintainable, and scalable architecture.

The primary goal of the application is to provide a RESTful API to manage a catalog of services and their versions, complete with authentication, authorization, and advanced data querying features.

## Core Features

*   **Full CRUD Functionality**: Create, Read, Update, and Delete services.
*   **Service Versioning**: Manage and retrieve different versions of a service.
*   **Advanced Listing**: List services with support for:
    *   Filtering by name.
    *   Sorting by name and creation date.
    *   Pagination.
*   **Authentication**: Secure endpoints using JWT-based authentication.
*   **Authorization**: Role-Based Access Control (RBAC) with `admin` and `viewer` roles. Admins have full CRUD access, while viewers have read-only access.
*   **Database Seeding**: Migrations include seed data for roles, permissions, and a default admin user, making the API ready for immediate testing.

## Architectural Decisions

The architecture is based on the principles of **Clean Architecture** to ensure a separation of concerns, making the application modular, testable, and independent of frameworks and external dependencies.

*   **Go Clean Template**: The project structure is heavily inspired by `evrone/go-clean-template`. This provides a solid foundation with a logical separation between business logic (`internal/usecase`), data models (`internal/entity`), and delivery mechanisms (`internal/controller`).
*   **PostgreSQL Database**: PostgreSQL was chosen for its reliability and robust feature set, which is well-suited for the relational data model of a service catalog and RBAC system.
*   **JWT for Authentication**: Stateless authentication is handled via JSON Web Tokens (JWT). This approach is scalable and simplifies the architecture, as no session state is stored on the server.
*   **Database Migrations**: `golang-migrate` is used to manage database schema changes with plain SQL files. This offers direct control over the schema and seeds the database with initial data.

## Database Design

The database schema includes tables for `users`, `roles`, `permissions`, `services`, and `service_versions`. This design supports the core application features, including RBAC and service versioning.

A visual representation of the schema is available on **[dbdiagram.io](https://dbdiagram.io/d/Kong-interview-1-68556831f039ec6d36293354)**.

## Getting Started

### Prerequisites

*   Go (version 1.23 or later)
*   Docker and Docker Compose
*   Make

### Setup and Running

1.  **Start Services**: Launch the PostgreSQL database using Docker.

    ```bash
    make docker-up
    ```

2.  **Run Migrations**: Apply database migrations to set up the schema and seed initial data.

    ```bash
    make migrate
    ```

3.  **Run the Application**: Start the API server.

    ```bash
    make run
    ```

    The application will be running on the port specified in your configuration (default is `8080`).

You can also use the `make init` command to perform all the above steps in one go.

## API Usage

### Postman Collection

A Postman collection is added to this repository to simplify testing of the API endpoints.

**[PostmanCollection](https://speeding-spaceship-875469.postman.co/workspace/HYDRA~a546c872-5d63-40ac-8159-cd1c0aabcf8b/collection/34528588-4f0faac2-7903-4b7a-ab62-96d3e82a98d4?action=share&creator=34528588)**

### Authentication

To access protected endpoints, you must first obtain a JWT token. The database is seeded with a default admin user:

*   **Email**: `admin@kong.org`
*   **Password**: `admin`

Send a `POST` request with these credentials to `/v1/auth/login`.

```json
{
  "email": "admin@kong.org",
  "password": "admin"
}
```

Include the returned token in the `Authorization` header of subsequent requests as a Bearer token.

### API Endpoints

Here are the main endpoints provided by the API:

| Method   | Endpoint                               | Description                                                                                                                                                             |
| :------- | :------------------------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `POST`   | `/v1/auth/login`                       | Authenticate and receive a JWT token.                                                                                                                                   |
| `GET`    | `/v1/services`                         | List all services with filtering, sorting, and pagination. Query params: `name` (filter), `sort` (`name_asc`, `name_desc`, `created_at_asc`, `created_at_desc`), `page`, `size`. |
| `POST`   | `/v1/services`                         | Create a new service (Admin only).                                                                                                                                      |
| `GET`    | `/v1/services/{id}`                    | Get details for a specific service.                                                                                                                                     |
| `PUT`    | `/v1/services/{id}`                    | Update an existing service (Admin only).                                                                                                                                |
| `DELETE` | `/v1/services/{id}`                    | Delete a service (Admin only).                                                                                                                                          |
| `GET`    | `/v1/services/{id}/versions`           | List all versions for a specific service.                                                                                                                               |
| `GET`    | `/v1/services/{id}/versions/{version}` | Get details for a specific service version.                                                                                                                             |

### Available Make Commands

- `make build`: Build the application
- `make run`: Run the application
- `make migrate`: Run database migrations
- `make docker-up`: Start Docker services
- `make docker-down`: Stop Docker services
- `make init`: Initialize development environment (docker-up, migrate, run)
- `make clean`: Clean build files
- `make test`: Run tests
