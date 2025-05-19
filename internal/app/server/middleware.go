package server

import (
	"net/http"
	"strconv"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// setupMiddleware configures the middleware for the Gin router
func (s *Server) setupMiddleware() {
	// Use Gin's built-in logger middleware
	s.router.Use(gin.Logger())

	// Setup custom CORS middleware
	s.router.Use(s.CORS())

	// Setup rate limiting middleware
	s.router.Use(s.RateLimiter())

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
		allowOrigin := s.config.AllowOrigin()
		if allowOrigin != "" {
			corsConfig.AllowOrigins = []string{allowOrigin}
		}
	}

	return cors.New(corsConfig)
}

// RateLimiter returns a middleware that limits request rates by IP address
func (s *Server) RateLimiter() gin.HandlerFunc {
	// Create a rate limit store - using the in-memory store
	// TODO: Use a more persistent store like Redis when scaling horizontally to ensure consistent rate limiting across instances
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 10, // 10 requests per second
	})

	// TODO: Configure trusted proxies when deploying behind a reverse proxy
	s.router.SetTrustedProxies(nil)

	// Return the middleware
	return ratelimit.RateLimiter(store, &ratelimit.Options{
		KeyFunc: func(c *gin.Context) string { return c.ClientIP() },
		ErrorHandler: func(c *gin.Context, info ratelimit.Info) {
			sec := int(time.Until(info.ResetTime).Seconds())
			c.Header("Retry-After", strconv.Itoa(sec))
			c.AbortWithStatusJSON(http.StatusTooManyRequests,
				gin.H{"message": "Too many requests. Please try again later."})
		},
	})
}
