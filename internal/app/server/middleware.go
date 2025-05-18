package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// setupMiddleware configures the middleware for the Gin router
func (s *Server) setupMiddleware() {
	// Use Gin's built-in logger middleware
	s.router.Use(gin.Logger())

	// Setup custom CORS middleware
	s.router.Use(s.CORS())

	// Use Gin's built-in recovery middleware
	s.router.Use(gin.Recovery())
}

// CORS returns a configured CORS middleware
func (s *Server) CORS() gin.HandlerFunc {
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	// Set allowed origins based on environment
	if s.config.IsDevelopment() {
		corsConfig.AllowAllOrigins = true
	} else {
		host := s.config.AllowedHost()
		if host != "" {
			corsConfig.AllowOrigins = []string{"https://" + host}
		}
	}

	return cors.New(corsConfig)
}
