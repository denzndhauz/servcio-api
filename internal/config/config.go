package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct to hold configuration values
type Config struct {
	DBType     string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	JWTSecret  string
}

// LoadConfig loads environment variables and returns a Config struct
func LoadConfig() Config {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	config := Config{
		DBType:     getEnv("DB_TYPE", "postgres"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
		DBPort:     getEnv("DB_PORT", "5432"),
		JWTSecret:  getEnv("JWT_SECRET", "defaultsecret"),
	}

	return config
}

// getEnv reads an environment variable or returns a default value
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
