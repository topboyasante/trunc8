package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string // env variables are read as strings, so we keep it as a string
}

type DatabaseConfig struct {
	Url string
}

func getRequiredEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("required environment variable %s is not set", key)
	}
	return value, nil
}

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Config vs *Config
// func LoadConfig() (Config, error) - Returns the actual struct

// Go copies the entire Config struct when you return it
// The caller gets their own independent copy
// If Config is small, this is fine and simple

// func LoadConfig() (*Config, error) - Returns a pointer to the struct

// Go returns the memory address where the Config lives
// The caller gets a reference to the original struct
// No copying happens
func LoadConfig() (*Config, error) {
	// Try to load .env file, but don't fail if it doesn't exist
	// This allows tests to work without a .env file
	_ = godotenv.Load()

	dbURL, err := getRequiredEnv("DATABASE_URL")
	if err != nil {
		return nil, err
	}

	fmt.Println("Loaded environment variables successfully")

	return &Config{
		Server: ServerConfig{
			Port: getEnvWithDefault("SERVER_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Url: dbURL,
		},
	}, nil
}
