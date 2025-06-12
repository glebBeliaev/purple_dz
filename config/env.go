package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		return
	}
	log.Println("Environment variables loaded")
}

type DataBaseConfig struct {
	url string
}

func NewDataBaseConfig() *DataBaseConfig {
	return &DataBaseConfig{
		url: getString("DATABASE_URL", "URL"),
	}
}

type LogConfig struct {
	Level  string
	Format string
}

func NewLogConfig() *LogConfig {
	return &LogConfig{
		Level:  strings.ToLower(strings.TrimSpace(getString("LOG_LEVEL", "info"))),
		Format: strings.ToLower(strings.TrimSpace(getString("LOG_FORMAT", "json"))),
	}
}

func getString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
