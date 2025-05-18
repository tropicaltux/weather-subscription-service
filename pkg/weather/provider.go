package weather

// ForecastType defines the type of forecast to retrieve
type ForecastType int

const (
	ForecastCurrent ForecastType = iota
	ForecastTomorrow
)

type WeatherData struct {
	Temperature float64
	Humidity    float64
	Description string
}

type Provider interface {
	GetWeather(city string, forceFresh bool, forecastType ForecastType) (*WeatherData, error)
}

func NewOpenMeteoProvider() Provider {
	return newOpenMeteoProvider()
}
