package main

import (
	"fmt"
	"log"
	"os"

	appServer "github.com/tropicaltux/weather-subscription-service/internal/app/server"
	serverConfig "github.com/tropicaltux/weather-subscription-service/internal/config/server"
)

func main() {
	// Create a new configuration
	config, err := serverConfig.NewConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Configuration error: %v\n", err)
		os.Exit(1)
	}

	// Print startup information in development mode
	if config.IsDevelopment() {
		fmt.Println("Starting server in DEVELOPMENT mode")
	} else {
		fmt.Println("Starting server in PRODUCTION mode")
	}

	// Create and start the HTTP server
	srv := appServer.New(config)

	log.Printf("Weather Subscription Service is running on :%s", config.Port())
	if err := srv.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
		os.Exit(1)
	}
}
