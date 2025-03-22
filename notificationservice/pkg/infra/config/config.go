package config

import (
	"os"
)

type DataBaseConfig struct {
	Port     string
	Host     string
	Username string
	DBName   string
	SSLMode  string
	Password string
}

type Config struct {
	HandlerPort string
	ServerPort  string
	DBConfig    DataBaseConfig
}

func NewConfig() *Config {
	return &Config{
		HandlerPort: getEnv("APP_PORT", "8082"),
		ServerPort:  getEnv("SERVER_ADDRESS", "50502"),
		DBConfig: DataBaseConfig{
			Port:     getEnv("DB_PORT", "5435"),
			Host:     getEnv("DB_HOST", "localhost"),
			Username: getEnv("DB_USERNAME", ""),
			DBName:   getEnv("DB_NAME", ""),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
			Password: getEnv("DB_PASSWORD", ""),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
