package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	Secret string
	Expire int // hours
}

// LoadConfig loads configuration from config.yaml and environment variables
// Environment variables take precedence over config.yaml values
func LoadConfig() (*Config, error) {
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()

	// Set up environment variable bindings with defaults
	bindEnvWithDefault("server.port", "SERVER_PORT", ":8080")
	bindEnvWithDefault("server.mode", "SERVER_MODE", "debug")
	bindEnvWithDefault("database.host", "DB_HOST", "127.0.0.1")
	bindEnvWithDefault("database.port", "DB_PORT", "3306")
	bindEnvWithDefault("database.user", "DB_USER", "root")
	bindEnvWithDefault("database.password", "DB_PASSWORD", "")
	bindEnvWithDefault("database.name", "DB_NAME", "adminpro")
	bindEnvWithDefault("jwt.secret", "JWT_SECRET", "")
	bindEnvWithDefault("jwt.expire", "JWT_EXPIRE", "24")

	// Try to read config file, but don't fail if it doesn't exist
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: config.yaml not found, using environment variables: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Validate required settings
	if cfg.Database.Password == "" {
		log.Println("Warning: DB_PASSWORD is not set")
	}
	if cfg.JWT.Secret == "" || cfg.JWT.Secret == "your_jwt_secret_key" {
		log.Println("Warning: JWT_SECRET is not set or using default value!")
	}

	return &cfg, nil
}

// bindEnvWithDefault binds an environment variable to a config key with a fallback default
func bindEnvWithDefault(key, envKey, defaultValue string) {
	// First check if environment variable is set
	if value := os.Getenv(envKey); value != "" {
		viper.SetDefault(key, value)
	}
	// Then bind the environment variable
	viper.BindEnv(key, envKey)
}
