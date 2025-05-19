package http

import (
	"github.com/tropicaltux/weather-subscription-service/internal/repository"
	"github.com/tropicaltux/weather-subscription-service/internal/services"
	"github.com/tropicaltux/weather-subscription-service/pkg/weather"
)

// Handler implements api.StrictServerInterface
type Handler struct {
	subscriptionService *services.SubscriptionService
	weatherService      *services.WeatherService
	// TODO: Add email service dependency to enable sending emails from handlers
	// emailService *services.EmailService
}

// NewHandler creates a new HTTP handler
func NewHandler(weatherProvider weather.Provider, subscriptionRepo repository.SubscriptionRepository) *Handler {
	weatherService := services.NewWeatherService(weatherProvider)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	// TODO: Initialize email service with appropriate configuration
	// emailService := services.NewEmailService(...)

	return &Handler{
		subscriptionService: subscriptionService,
		weatherService:      weatherService,
		// emailService: emailService,
	}
}
