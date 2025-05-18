# Weather Provider

Open-Meteo client with caching support.

## Usage

```go
provider := weather.NewOpenMeteoProvider()

// Current weather
data, err := provider.GetWeather("London", false, weather.ForecastCurrent)

// Tomorrow's forecast
tomorrow, err := provider.GetWeather("London", false, weather.ForecastTomorrow)

// Bypass cache
fresh, err := provider.GetWeather("London", true, weather.ForecastCurrent)
```

Returns `WeatherData` with Temperature (°C), Humidity (%), and Description.

## Weather Data
- `Temperature` - in Celsius
- `Humidity` - percentage
- `Description` - human-readable weather condition

## Thread Safety

The provider is thread-safe and can be safely used from multiple goroutines.

## About Open-Meteo

[Open-Meteo](https://open-meteo.com/) is an open-source weather API offering free access for non-commercial use with no API key required. The API is based on open data from national weather services. 