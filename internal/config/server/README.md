# Server Configuration

Configuration for the Weather Subscription Service HTTP server.

## Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `APP_ENV` | Environment (`development` or `production`) | `production` | No |
| `ALLOW_ORIGIN` | Server URL for CORS origins | - | Yes (in production) |
| `PORT` | Server listening port | `3000` | No |

## Features

- **Development mode**: Allows CORS requests from all origins (`*`)
- **Production mode**: Restricts CORS origins to value of `ALLOW_ORIGIN`

## Public API

| Function | Description | Return Type |
|----------|-------------|------------|
| `NewConfig()` | Creates configuration from environment variables | `(*Config, error)` |
| `config.IsDevelopment()` | Checks if environment is development | `bool` |
| `config.IsProduction()` | Checks if environment is production | `bool` |
| `config.AllowOrigin()` | Returns the allowed origin | `string` |
| `config.Port()` | Returns the configured port | `string` |

## Usage

```go
import "github.com/tropicaltux/weather-subscription-service/internal/config/server"

func main() {
    config, err := server.NewConfig()
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Server running on port %s\n", config.Port())
}
``` 