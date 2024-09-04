package config

import (
	"log"
	"os"
)

type Config struct {
	DBSource     string
	ClientSecret string
}

func LoadConfig() *Config {
	return &Config{
		DBSource: getEnv("DB_SOURCE"),
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	log.Fatalf("Required environment variable %s is not set. Please initialize it.", key)
	return ""
}
