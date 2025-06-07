package config

import (
	"log"
	"os"

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

func getString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func NewDataBaseConfig() *DataBaseConfig {
	return &DataBaseConfig{
		url: getString("DATABASE_URL", "URL"),
	}
}
