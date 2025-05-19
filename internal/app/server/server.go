package server

import (
	"log"

	"github.com/gin-gonic/gin"
	api "github.com/tropicaltux/weather-subscription-service/internal/api/http"
	"github.com/tropicaltux/weather-subscription-service/internal/config"
	serverConfig "github.com/tropicaltux/weather-subscription-service/internal/config/server"
	"github.com/tropicaltux/weather-subscription-service/internal/database"
	handlers "github.com/tropicaltux/weather-subscription-service/internal/handlers/http"
	"github.com/tropicaltux/weather-subscription-service/internal/repository/db"
	"github.com/tropicaltux/weather-subscription-service/pkg/weather"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	config *serverConfig.Config
	db     *gorm.DB
}

func New(srvConfig *serverConfig.Config) *Server {
	// Set Gin mode based on environment
	if srvConfig.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize database connection using the standalone config package
	dbConfig := config.LoadDatabaseConfig()
	dbConn, err := database.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run database migrations
	if err := database.RunMigrations(dbConn); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}
	log.Println("Database migrations completed successfully")

	s := &Server{
		router: gin.New(),
		config: srvConfig,
		db:     dbConn,
	}

	s.setupMiddleware()
	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	weatherProvider := weather.NewOpenMeteoProvider()
	subscriptionRepo := db.NewPostgresSubscriptionRepository(s.db)

	// TODO: Implement email service for sending subscription confirmations and weather updates
	// emailService := email.NewEmailService(...)

	handler := handlers.NewHandler(weatherProvider, subscriptionRepo)

	// Use the strict handler adapter to bridge between Gin and our strict typed handler
	apiGroup := s.router.Group("/api")
	strictHandler := api.NewStrictHandler(handler, nil)
	api.RegisterHandlers(apiGroup, strictHandler)

	// TODO: Add a scheduler for periodic weather updates to subscribers
}

// Start starts the HTTP server on the configured port
func (s *Server) Start() error {
	return s.router.Run(":" + s.config.Port())
}
