package config

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	DatabaseURL        string
	RedisURL           string
	RedisPassword      string
	JWTSecret          string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
	GRPCPort           string
	HTTPPort           string
	KafkaBrokers       []string
	KafkaTopicPrefix   string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists (ignore error if file doesn't exist)
	_ = godotenv.Load()

	accessExpiry, _ := strconv.Atoi(getEnv("ACCESS_TOKEN_EXPIRY", "900"))      // 15 minutes
	refreshExpiry, _ := strconv.Atoi(getEnv("REFRESH_TOKEN_EXPIRY", "604800")) // 7 days

	accessTokenExpiry := time.Duration(accessExpiry) * time.Second
	refreshTokenExpiry := time.Duration(refreshExpiry) * time.Second

	kafkaBrokersStr := getEnv("KAFKA_BROKERS", "localhost:9092")
	var kafkaBrokers []string
	if kafkaBrokersStr != "" {
		// Split by comma if multiple brokers
		brokers := strings.Split(kafkaBrokersStr, ",")
		for _, broker := range brokers {
			if trimmed := strings.TrimSpace(broker); trimmed != "" {
				kafkaBrokers = append(kafkaBrokers, trimmed)
			}
		}
	}
	if len(kafkaBrokers) == 0 {
		kafkaBrokers = []string{"localhost:9092"} // Default
	}

	return &Config{
		DatabaseURL:        getEnv("DATABASE_URL", ""),
		RedisURL:           getEnv("REDIS_URL", "localhost:6379"),
		RedisPassword:      getEnv("REDIS_PASSWORD", ""),
		GRPCPort:           getEnv("GRPC_PORT", "50056"),
		HTTPPort:           getEnv("HTTP_PORT", "8082"),
		JWTSecret:          getEnv("JWT_SECRET", "your-secret-key"),
		AccessTokenExpiry:  accessTokenExpiry,
		RefreshTokenExpiry: refreshTokenExpiry,
		KafkaBrokers:       kafkaBrokers,
		KafkaTopicPrefix:   getEnv("KAFKA_TOPIC_PREFIX", "workforce"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
