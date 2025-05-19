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
}

// NewHandler creates a new HTTP handler
func NewHandler(weatherProvider weather.Provider, subscriptionRepo repository.SubscriptionRepository) *Handler {
	weatherService := services.NewWeatherService(weatherProvider)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	return &Handler{
		subscriptionService: subscriptionService,
		weatherService:      weatherService,
	}
}
