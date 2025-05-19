# Weather Subscription Service Server

This package contains the main entry point for the Weather Subscription Service HTTP server.

## Usage

To run the server:

```bash
# From the root directory
go run cmd/server/main.go

# Or build and run
go build -o weather-server cmd/server/main.go
./weather-server
```

## Environment Variables

The server uses environment variables for configuration:

```bash
# Development mode
APP_ENV=development go run cmd/server/main.go

# Production mode with specific allowed origin
APP_ENV=production ALLOW_ORIGIN="https://weather.example.com" go run cmd/server/main.go

# Specify port
PORT=8080 go run cmd/server/main.go
```

See the [configuration](../../internal/config/server) package for all available environment variables. 