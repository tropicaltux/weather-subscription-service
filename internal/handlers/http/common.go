package http

import (
	"github.com/tropicaltux/weather-subscription-service/pkg/weather"
)

// Handler implements api.StrictServerInterface
type Handler struct {
	weatherProvider weather.Provider
}

// NewHandler creates a new HTTP handler
func NewHandler(weatherProvider weather.Provider) *Handler {
	return &Handler{
		weatherProvider: weatherProvider,
	}
}
