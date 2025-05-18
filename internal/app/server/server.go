package server

import (
	"github.com/gin-gonic/gin"
	api "github.com/tropicaltux/weather-subscription-service/internal/api/http"
	"github.com/tropicaltux/weather-subscription-service/internal/config/server"
	handlers "github.com/tropicaltux/weather-subscription-service/internal/handlers/http"
	"github.com/tropicaltux/weather-subscription-service/pkg/weather"
)

type Server struct {
	router *gin.Engine
	config *server.Config
}

func New(config *server.Config) *Server {
	// Set Gin mode based on environment
	if config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	s := &Server{
		router: gin.New(),
		config: config,
	}

	s.setupMiddleware()
	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	weatherProvider := weather.NewOpenMeteoProvider()

	handler := handlers.NewHandler(weatherProvider)

	// Use the strict handler adapter to bridge between Gin and our strict typed handler
	apiGroup := s.router.Group("/api")
	strictHandler := api.NewStrictHandler(handler, nil)
	api.RegisterHandlers(apiGroup, strictHandler)
}

// Start starts the HTTP server on the configured port
func (s *Server) Start() error {
	return s.router.Run(":" + s.config.Port())
}
