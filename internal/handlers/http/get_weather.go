package http

import (
	"context"
	"strings"

	api "github.com/tropicaltux/weather-subscription-service/internal/api/http"
)

// GetWeather handles requests for weather information for a city
func (h *Handler) GetWeather(ctx context.Context, request api.GetWeatherRequestObject) (api.GetWeatherResponseObject, error) {
	request.Params.City = strings.TrimSpace(request.Params.City)

	// Validate input
	if request.Params.City == "" {
		return api.GetWeather400JSONResponse{
			Message: "city parameter is required",
		}, nil
	}

	// TODO: Add business logic
	// Return placeholder response
	return api.GetWeather200JSONResponse{
		Temperature: nil,
		Humidity:    nil,
		Description: nil,
	}, nil
}
