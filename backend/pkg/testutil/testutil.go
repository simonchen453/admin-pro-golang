package testutil

import (
	"os"
	"testing"

	"admin-pro/internal/config"
)

// GetTestConfig returns a test configuration
func GetTestConfig() *config.Config {
	return &config.Config{
		Server: config.ServerConfig{
			Port: ":8081",
			Mode: "test",
		},
		Database: config.DatabaseConfig{
			Host:     os.Getenv("TEST_DB_HOST"),
			Port:     os.Getenv("TEST_DB_PORT"),
			User:     os.Getenv("TEST_DB_USER"),
			Password: os.Getenv("TEST_DB_PASSWORD"),
			Name:     os.Getenv("TEST_DB_NAME"),
		},
		JWT: config.JWTConfig{
			Secret: "test_jwt_secret_key",
			Expire: 24,
		},
	}
}

// SkipIfShort skips the test if -short flag is provided
func SkipIfShort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
}
