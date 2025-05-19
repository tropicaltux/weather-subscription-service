# Weather Subscription Service

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.24-00ADD8?style=flat-square&logo=go" alt="Go Version" />
  <img src="https://img.shields.io/badge/License-MIT-blue?style=flat-square" alt="License" />
  <img src="https://img.shields.io/badge/API-OpenAPI%203.0-green?style=flat-square" alt="API Specification" />
</p>

A Go-based service that allows users to subscribe to weather updates for their city via email. The service provides weather data using the OpenMeteo API and allows users to manage their subscriptions through a RESTful API.

This project uses an **API-first approach** with OpenAPI specification. The API design is defined in the [api/openapi.yaml](api/openapi.yaml) file, and the server code is generated from this specification.

## 🏗️ Project Structure

```
.
├── api/                # API specification and documentation
├── cmd/                # Application entry points
│   └── server/         # Main HTTP server executable
├── internal/           # Private application code
│   ├── api/            # API implementation
│   ├── app/            # Application components
│   ├── config/         # Configuration management
│   ├── database/       # Database connection and migration
│   ├── handlers/       # HTTP request handlers
│   ├── repository/     # Data access layer
│   └── services/       # Business logic services
├── pkg/                # Public libraries that can be used by external applications
│   └── weather/        # Weather provider implementations
└── compose.yml         # Docker Compose configuration for local development
```

## 📚 Documentation

- [API Documentation](api/README.md) - OpenAPI specification and API details
- [Server Package](cmd/server/README.md) - Main server executable documentation
- [Weather Package](pkg/weather/README.md) - Weather provider implementation
- [Server Configuration](internal/config/server/README.md) - Server configuration docs

## 🚀 Running the Application

### Prerequisites

- Docker and Docker Compose
- Go 1.24 or higher (for local development)
- PostgreSQL (if running without Docker)

### Using Docker Compose

```bash
# Start all services (API server, PostgreSQL database, Swagger UI)
docker-compose up -d

# Check logs
docker-compose logs -f

# Stop all services
docker-compose down
```

### Local Development

```bash
# Set up environment variables
export APP_ENV=development
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=weather_subscription

# Run the server
go run cmd/server/main.go
```

### API Access

- API Endpoint: http://localhost:3000/api
- Swagger UI: http://localhost:8080

## ✅ Completed

- [x] API-first design with OpenAPI specification
- [x] Basic REST API structure
- [x] Weather data provider implementation
- [x] Database connection and migrations
- [x] Subscription data model
- [x] Configuration management
- [x] Docker and Docker Compose setup
- [x] CORS and rate limiting middleware

## 📝 TODO

- [ ] Implement email sending functionality for subscription notifications
- [ ] Add scheduler for sending periodic weather updates
- [ ] Implement production-grade database migrations
- [ ] Create frontend application for subscription management
- [ ] Add comprehensive unit and integration tests
- [ ] Deploy to cloud infrastructure using Terraform

## 📄 License

This project is licensed under the terms found in the [LICENSE](LICENSE) file. 