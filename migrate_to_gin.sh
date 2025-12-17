#!/bin/bash

# Gin Migration Script for Workforce Service
# This script helps migrate from Gorilla Mux to Gin framework

echo "=== Workforce Service - Gin Framework Migration ==="
echo ""
echo "This script will help you integrate Gin framework into your workforce service."
echo ""

# Backup existing main.go
echo "Step 1: Backing up existing main.go..."
if [ -f "cmd/server/main.go" ]; then
    cp cmd/server/main.go cmd/server/main.go.backup
    echo "✓ Backup created: cmd/server/main.go.backup"
else
    echo "✗ main.go not found"
    exit 1
fi

# Create new main.go with Gin
echo ""
echo "Step 2: Creating new main.go with Gin framework..."

cat > cmd/server/main.go << 'EOF'
package main

import (
	"context"
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
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

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
EOF

echo "✓ New main.go created with Gin framework"

echo ""
echo "Step 3: Running go mod tidy..."
go mod tidy

echo ""
echo "=== Migration Complete ==="
echo ""
echo "The following changes have been made:"
echo "  1. Gin framework and middleware installed"
echo "  2. New HTTP handlers created in internal/adapters/inbound/http/"
echo "  3. main.go updated to use Gin server"
echo "  4. Old main.go backed up to main.go.backup"
echo ""
echo "To start the server:"
echo "  go run cmd/server/main.go"
echo ""
echo "To rollback:"
echo "  mv cmd/server/main.go.backup cmd/server/main.go"
echo ""
EOF
