package config

import (
	"log"
	"os"
	"strconv"

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
	Url string
}

func NewDataBaseConfig() *DataBaseConfig {
	return &DataBaseConfig{
		Url: getString("DATABASE_URL_JOB", "URL"),
	}
}

type LogConfig struct {
	Level  int
	Format string
}

func NewLogConfig() *LogConfig {
	return &LogConfig{
		Level:  getInt("LOG_LEVEL", 1),
		Format: getString("LOG_FORMAT", "json"),
	}
}

// Всмпомогательные функции
func getString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}
