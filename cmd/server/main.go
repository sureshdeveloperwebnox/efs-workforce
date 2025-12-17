package main

import (
	"efs-workforce/internal/adapters/outbound/kafka"
	"efs-workforce/internal/adapters/outbound/postgres"
	httpHandler "efs-workforce/internal/adapters/inbound/http"
	"efs-workforce/internal/application"
	"efs-workforce/internal/config"
	"efs-workforce/internal/database"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Println("Starting Workforce Service with Gin Framework...")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := database.InitGORM(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize repositories
	roleRepo := postgres.NewRoleRepository(db)

	// Initialize Kafka event publisher
	var eventPublisher *kafka.EventPublisher
	ep, err := kafka.NewEventPublisher(cfg.KafkaBrokers, cfg.KafkaTopicPrefix)
	if err != nil {
		log.Printf("Warning: Failed to initialize Kafka publisher: %v. Continuing without event publishing.", err)
		eventPublisher = nil
	} else {
		eventPublisher = ep
	}

	// Initialize services
	roleService := application.NewRoleService(roleRepo, eventPublisher)

	// Initialize Gin HTTP server
	ginServer := httpHandler.NewGinServer(roleService, cfg.HTTPPort)
	ginServer.SetupRouter()

	// Initialize gRPC server (keeping existing gRPC functionality)
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// Start gRPC server in goroutine
	grpcListener, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen on gRPC port: %v", err)
	}

	go func() {
		log.Printf("Starting gRPC server on port %s", cfg.GRPCPort)
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Start Gin HTTP server in goroutine
	go func() {
		if err := ginServer.Start(); err != nil {
			log.Fatalf("Failed to start Gin HTTP server: %v", err)
		}
	}()

	// Wait for interrupt signal for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown

	log.Println("Shutting down HTTP server...")
	// Shutdown gRPC server
	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()

	// Close Kafka publisher if initialized
	if eventPublisher != nil {
		log.Println("Closing Kafka publisher...")
		eventPublisher.Close()
	}

	// Close database connection
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.Close()
	}

	log.Println("Server exited gracefully")
}
