package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config holds all application configurations.
type Config struct {
	ServerPort   string
	KafkaBroker  string
	StripeSecret string
	PostgresURL  string
}

// LoadConfig reads environment variables and loads them into a Config struct.
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	if err := godotenv.Load(filepath.Join("src", "payment-service", ".env")); err != nil {
		log.Println("No .env file found, relying on environment variables.")
	}

	cfg := &Config{
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		KafkaBroker:  getEnv("KAFKA_BROKER", ""),
		StripeSecret: getEnv("STRIPE_SECRET", ""),
		PostgresURL:  getEnv("POSTGRES_URL", ""),
	}

	if cfg.PostgresURL == "" || cfg.KafkaBroker == "" || cfg.StripeSecret == "" {
		log.Println("Warning: Some environment variables are missing!")
	}

	return cfg, nil
}

// getEnv fetches environment variables with a fallback default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
