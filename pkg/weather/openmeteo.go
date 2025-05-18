package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	openMeteoBaseURL            = "https://api.open-meteo.com/v1/forecast"
	defaultCacheExpiration      = 5 * time.Minute
	defaultCacheCleanupInterval = 10 * time.Minute
)

// openMeteoProvider implements the Provider interface using Open-Meteo API
type openMeteoProvider struct {
	client *http.Client
	cache  *cache.Cache
}

func newOpenMeteoProvider() *openMeteoProvider {
	return &openMeteoProvider{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		cache: cache.New(defaultCacheExpiration, defaultCacheCleanupInterval),
	}
}

// GetWeather retrieves weather data for the specified city
// If forceFresh is true, it will always fetch fresh data
// forecastType determines whether to get current weather or tomorrow's forecast
func (p *openMeteoProvider) GetWeather(city string, forceFresh bool, forecastType ForecastType) (*WeatherData, error) {
	cacheKey := fmt.Sprintf("%s_%d", city, forecastType)

	// Check cache first if not forcing fresh data
	if !forceFresh {
		if cachedData, found := p.cache.Get(cacheKey); found {
			return cachedData.(*WeatherData), nil
		}
	}

	// Fetch fresh data
	var weatherData *WeatherData
	var err error

	switch forecastType {
	case ForecastCurrent:
		weatherData, err = p.getCurrentWeather(city)
	case ForecastTomorrow:
		weatherData, err = p.getTomorrowWeather(city)
	default:
		return nil, fmt.Errorf("unknown forecast type: %d", forecastType)
	}

	if err != nil {
		return nil, err
	}

	// Store in cache
	p.cache.Set(cacheKey, weatherData, defaultCacheExpiration)

	return weatherData, nil
}

// openMeteoCurrentResponse represents the response from Open-Meteo API for current weather
type openMeteoCurrentResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Current   struct {
		Temperature2m    float64 `json:"temperature_2m"`
		RelativeHumidity float64 `json:"relative_humidity_2m"`
		WeatherCode      int     `json:"weather_code"`
	} `json:"current"`
}

// openMeteoForecastResponse represents the response from Open-Meteo API for forecast weather
type openMeteoForecastResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Daily     struct {
		Time               []string  `json:"time"`
		Temperature2mMax   []float64 `json:"temperature_2m_max"`
		Temperature2mMin   []float64 `json:"temperature_2m_min"`
		RelativeHumidity2m []float64 `json:"relative_humidity_2m_mean"`
		WeatherCode        []int     `json:"weather_code"`
	} `json:"daily"`
}

// makeAPIRequest performs an HTTP request to the API and handles the common response processing
func (p *openMeteoProvider) makeAPIRequest(url string, v interface{}) error {
	// Make the request
	resp, err := p.client.Get(url)
	if err != nil {
		return fmt.Errorf("error making request to API: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	// Parse the response
	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("error decoding API response: %w", err)
	}

	return nil
}

// getCurrentWeather fetches current weather data
func (p *openMeteoProvider) getCurrentWeather(city string) (*WeatherData, error) {
	// Build the request URL
	url := fmt.Sprintf("%s?latitude=0&longitude=0&current=temperature_2m,relative_humidity_2m,weather_code&geocoding_api=true&city=%s", openMeteoBaseURL, city)

	// Make API request
	var apiResp openMeteoCurrentResponse
	if err := p.makeAPIRequest(url, &apiResp); err != nil {
		return nil, err
	}

	// Convert to our weather data format
	weatherData := &WeatherData{
		Temperature: apiResp.Current.Temperature2m,
		Humidity:    apiResp.Current.RelativeHumidity,
		Description: weatherCodeToDescription(apiResp.Current.WeatherCode),
	}

	return weatherData, nil
}

// getTomorrowWeather fetches tomorrow's weather forecast
func (p *openMeteoProvider) getTomorrowWeather(city string) (*WeatherData, error) {
	// Build the request URL
	url := fmt.Sprintf("%s?latitude=0&longitude=0&daily=temperature_2m_max,temperature_2m_min,relative_humidity_2m_mean,weather_code&forecast_days=2&geocoding_api=true&city=%s", openMeteoBaseURL, city)

	// Make API request
	var apiResp openMeteoForecastResponse
	if err := p.makeAPIRequest(url, &apiResp); err != nil {
		return nil, err
	}

	// Make sure we have data for tomorrow
	if len(apiResp.Daily.Time) < 2 || len(apiResp.Daily.Temperature2mMax) < 2 ||
		len(apiResp.Daily.Temperature2mMin) < 2 || len(apiResp.Daily.RelativeHumidity2m) < 2 ||
		len(apiResp.Daily.WeatherCode) < 2 {
		return nil, fmt.Errorf("not enough forecast data returned")
	}

	// Get the average temperature for tomorrow (index 1)
	avgTemp := (apiResp.Daily.Temperature2mMax[1] + apiResp.Daily.Temperature2mMin[1]) / 2

	// Convert to our weather data format
	weatherData := &WeatherData{
		Temperature: avgTemp,
		Humidity:    apiResp.Daily.RelativeHumidity2m[1],
		Description: weatherCodeToDescription(apiResp.Daily.WeatherCode[1]),
	}

	return weatherData, nil
}

// weatherCodeToDescription converts Open-Meteo weather codes to human-readable descriptions
func weatherCodeToDescription(code int) string {
	switch code {
	case 0:
		return "Clear sky"
	case 1:
		return "Mainly clear"
	case 2:
		return "Partly cloudy"
	case 3:
		return "Overcast"
	case 45, 48:
		return "Fog"
	case 51, 53, 55:
		return "Drizzle"
	case 56, 57:
		return "Freezing drizzle"
	case 61, 63, 65:
		return "Rain"
	case 66, 67:
		return "Freezing rain"
	case 71, 73, 75:
		return "Snow fall"
	case 77:
		return "Snow grains"
	case 80, 81, 82:
		return "Rain showers"
	case 85, 86:
		return "Snow showers"
	case 95:
		return "Thunderstorm"
	case 96, 99:
		return "Thunderstorm with hail"
	default:
		return "Unknown"
	}
}
