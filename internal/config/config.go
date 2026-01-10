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

// LoadConfig 从 config.yaml 和 环境变量加载配置
// 环境变量优先级高于 config.yaml 中的值
func LoadConfig() (*Config, error) {
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()

	// 设置环境变量绑定及其默认值
	bindEnvWithDefault("server.port", "SERVER_PORT", ":8080")
	bindEnvWithDefault("server.mode", "SERVER_MODE", "debug")
	bindEnvWithDefault("database.host", "DB_HOST", "127.0.0.1")
	bindEnvWithDefault("database.port", "DB_PORT", "3306")
	bindEnvWithDefault("database.user", "DB_USER", "root")
	bindEnvWithDefault("database.password", "DB_PASSWORD", "")
	bindEnvWithDefault("database.name", "DB_NAME", "adminpro")
	bindEnvWithDefault("jwt.secret", "JWT_SECRET", "")
	bindEnvWithDefault("jwt.expire", "JWT_EXPIRE", "24")

	// 尝试读取配置文件，如果不存在则不报错
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("警告:config.yaml 未找到，使用环境变量: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// 校验必要配置
	if cfg.Database.Password == "" {
		log.Println("警告: DB_PASSWORD 未设置")
	}
	if cfg.JWT.Secret == "" || cfg.JWT.Secret == "your_jwt_secret_key" {
		log.Println("警告: JWT_SECRET 未设置或正在使用默认值!")
	}

	return &cfg, nil
}

// bindEnvWithDefault 将环境变量绑定到配置键，并提供默认值
func bindEnvWithDefault(key, envKey, defaultValue string) {
	// 首先检查环境变量是否设置
	if value := os.Getenv(envKey); value != "" {
		viper.SetDefault(key, value)
	}
	// 然后绑定环境变量
	viper.BindEnv(key, envKey)
}
