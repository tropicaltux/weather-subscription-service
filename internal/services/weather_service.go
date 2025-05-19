package services

import (
	"context"
	"errors"

	"github.com/tropicaltux/weather-subscription-service/pkg/weather"
)

var (
	ErrCityEmpty          = errors.New("city cannot be empty")
	ErrWeatherUnavailable = errors.New("weather data unavailable")
)

type WeatherData struct {
	City        string
	Temperature float32
	Humidity    float32
	Description string
}

type WeatherService struct {
	weatherProvider weather.Provider
}

func NewWeatherService(weatherProvider weather.Provider) *WeatherService {
	return &WeatherService{
		weatherProvider: weatherProvider,
	}
}

func (s *WeatherService) GetCurrentWeather(ctx context.Context, city string) (*WeatherData, error) {
	if city == "" {
		return nil, ErrCityEmpty
	}

	weatherData, err := s.weatherProvider.GetWeather(city, false, weather.ForecastCurrent)
	if err != nil {
		return nil, err
	}

	return &WeatherData{
		City:        city,
		Temperature: float32(weatherData.Temperature),
		Humidity:    float32(weatherData.Humidity),
		Description: weatherData.Description,
	}, nil
}
