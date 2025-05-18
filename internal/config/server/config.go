package server

import (
	"errors"
	"os"
)

const (
	envDevelopment = "development"
	envProduction  = "production"
)

// Default values
const (
	defaultPort = "3000"
)

// Config holds the configuration for the HTTP server
type Config struct {
	environment string
	port        string
	allowedHost string
}

func NewConfig() (*Config, error) {
	env := getEnvironment()
	allowedHost := getAllowedHost()

	// Validate that ALLOWED_HOST is set in production mode
	if env == envProduction && allowedHost == "" {
		return nil, errors.New("ALLOWED_HOST environment variable is required in production mode")
	}

	config := &Config{
		environment: env,
		port:        getPort(),
		allowedHost: allowedHost,
	}
	return config, nil
}

func (c *Config) IsDevelopment() bool {
	return c.environment == envDevelopment
}

func (c *Config) IsProduction() bool {
	return c.environment == envProduction
}

func (c *Config) AllowedHost() string {
	return c.allowedHost
}

func (c *Config) Port() string {
	return c.port
}

func getEnvironment() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		return envProduction // Default to production if not specified
	}

	// Validate environment value
	if env != envDevelopment && env != envProduction {
		return envProduction // Default to production if invalid value
	}

	return env
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return defaultPort
	}
	return port
}

func getAllowedHost() string {
	return os.Getenv("ALLOWED_HOST")
}
