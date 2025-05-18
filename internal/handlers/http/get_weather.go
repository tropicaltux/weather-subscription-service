package http

import (
	"context"
	"strings"

	api "github.com/tropicaltux/weather-subscription-service/internal/api/http"
	"github.com/tropicaltux/weather-subscription-service/pkg/weather"
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

	// Fetch current weather data for the city
	weatherData, err := h.weatherProvider.GetWeather(request.Params.City, false, weather.ForecastCurrent)
	if err != nil {
		return api.GetWeather404JSONResponse{
			Message: "Failed to retrieve weather data: " + err.Error(),
		}, nil
	}

	// Convert weather data to response
	temp := float32(weatherData.Temperature)
	humidity := float32(weatherData.Humidity)
	description := weatherData.Description

	return api.GetWeather200JSONResponse{
		Temperature: &temp,
		Humidity:    &humidity,
		Description: &description,
	}, nil
}
