package config

import (
	"os"
	"testing"
)

func TestLoadConfig_Success(t *testing.T) {
	// Set required env var
	os.Setenv("DATABASE_URL", "postgres://localhost:5432/testdb")
	defer os.Unsetenv("DATABASE_URL")

	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if config.Database.Url != "postgres://localhost:5432/testdb" {
		t.Errorf("Expected DATABASE_URL to be 'postgres://localhost:5432/testdb', got '%s'", config.Database.Url)
	}

	// Should use default port
	if config.Server.Port != "8080" {
		t.Errorf("Expected default port '8080', got '%s'", config.Server.Port)
	}
}

func TestLoadConfig_CustomPort(t *testing.T) {
	// Set both env vars
	os.Setenv("DATABASE_URL", "postgres://localhost:5432/testdb")
	os.Setenv("SERVER_PORT", "3000")
	defer func() {
		os.Unsetenv("DATABASE_URL")
		os.Unsetenv("SERVER_PORT")
	}()

	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if config.Server.Port != "3000" {
		t.Errorf("Expected custom port '3000', got '%s'", config.Server.Port)
	}
}

func TestLoadConfig_MissingDatabaseURL(t *testing.T) {
	// Make sure DATABASE_URL is not set
	os.Unsetenv("DATABASE_URL")

	config, err := LoadConfig()
	if err == nil {
		t.Fatal("Expected error when DATABASE_URL is missing, got nil")
	}

	if config != nil {
		t.Error("Expected config to be nil when error occurs")
	}

	expectedError := "required environment variable DATABASE_URL is not set"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}
}

func TestGetEnvWithDefault(t *testing.T) {
	// Test with env var set
	os.Setenv("TEST_VAR", "custom_value")
	defer os.Unsetenv("TEST_VAR")

	result := getEnvWithDefault("TEST_VAR", "default_value")
	if result != "custom_value" {
		t.Errorf("Expected 'custom_value', got '%s'", result)
	}

	// Test with env var not set
	result = getEnvWithDefault("NONEXISTENT_VAR", "default_value")
	if result != "default_value" {
		t.Errorf("Expected 'default_value', got '%s'", result)
	}
}

func TestGetRequiredEnv(t *testing.T) {
	// Test with env var set
	os.Setenv("REQUIRED_TEST_VAR", "test_value")
	defer os.Unsetenv("REQUIRED_TEST_VAR")

	value, err := getRequiredEnv("REQUIRED_TEST_VAR")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if value != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", value)
	}

	// Test with env var not set
	_, err = getRequiredEnv("NONEXISTENT_REQUIRED_VAR")
	if err == nil {
		t.Fatal("Expected error for missing required env var, got nil")
	}
}
