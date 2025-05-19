package http

import (
	"context"
	"strings"

	api "github.com/tropicaltux/weather-subscription-service/internal/api/http"
	"github.com/tropicaltux/weather-subscription-service/internal/services"
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

	// TODO: Implement city validation to check if city exists and is correctly formatted

	weatherData, err := h.weatherService.GetCurrentWeather(ctx, request.Params.City)
	if err != nil {
		if err == services.ErrCityEmpty {
			return api.GetWeather400JSONResponse{
				Message: "city parameter is required",
			}, nil
		}

		return api.GetWeather404JSONResponse{
			Message: "Failed to retrieve weather data: " + err.Error(),
		}, nil
	}

	// Convert weather data to response
	temp := weatherData.Temperature
	humidity := weatherData.Humidity
	description := weatherData.Description

	return api.GetWeather200JSONResponse{
		Temperature: &temp,
		Humidity:    &humidity,
		Description: &description,
	}, nil
}
