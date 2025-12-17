package http

import (
	"efs-workforce/internal/adapters/inbound/http/handlers"
	"efs-workforce/internal/adapters/inbound/http/middleware"
	"efs-workforce/internal/adapters/inbound/http/routes"
	"efs-workforce/internal/application"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GinServer represents the Gin HTTP server
type GinServer struct {
	router      *gin.Engine
	roleService *application.RoleService
	port        string
}

// NewGinServer creates a new Gin HTTP server
func NewGinServer(roleService *application.RoleService, port string) *GinServer {
	// Set Gin mode based on environment
	// gin.SetMode(gin.ReleaseMode) // Uncomment for production

	return &GinServer{
		roleService: roleService,
		port:        port,
	}
}

// SetupRouter configures the Gin router with all middleware and routes
func (s *GinServer) SetupRouter() *gin.Engine {
	router := gin.New()

	// Global middleware
	router.Use(middleware.Recovery())        // Panic recovery
	router.Use(middleware.Logger())          // Request logging
	router.Use(middleware.CORS())            // CORS handling
	router.Use(middleware.RequestID())       // Request ID tracking
	router.Use(middleware.SecurityHeaders()) // Security headers

	// Health check endpoint (outside versioned API)
	router.GET("/healthz", middleware.HealthCheck)
	router.GET("/health", middleware.HealthCheck)

	// Initialize handlers
	roleHandler := handlers.NewRoleHandler(s.roleService)

	// Setup versioned routes
	routes.SetupV1Routes(router, roleHandler)

	s.router = router
	return router
}

// Start starts the HTTP server
func (s *GinServer) Start() error {
	if s.router == nil {
		s.SetupRouter()
	}

	// Configure HTTP server with timeouts
	server := &http.Server{
		Addr:           ":" + s.port,
		Handler:        s.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	log.Printf("Starting Gin HTTP server on port %s", s.port)
	log.Printf("Health check available at http://localhost:%s/healthz", s.port)
	log.Printf("API v1 available at http://localhost:%s/api/v1", s.port)

	return server.ListenAndServe()
}

// GetRouter returns the configured Gin router
func (s *GinServer) GetRouter() *gin.Engine {
	if s.router == nil {
		s.SetupRouter()
	}
	return s.router
}
