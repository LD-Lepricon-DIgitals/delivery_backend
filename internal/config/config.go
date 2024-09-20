package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// TODO: define Config struct
type Config struct {
	HostPort   string
	DBUsername string
	DBPassword string
	DBHostname string
	DBName     string
	DBPort     string
	SSLMode    string
}

//TODO: define

func NewConfig() *Config {
	err := godotenv.Load("config/local.env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	cfg := &Config{
		HostPort:   getEnv("HOST_PORT", "8080"),
		DBUsername: getEnv("DB_USERNAME", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBHostname: getEnv("DB_HOSTNAME", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "db"),
		SSLMode:    getEnv("SSL_MODE", "disable"),
	}

	if cfg.DBPassword == "" {
		log.Fatal("DB_PASSWORD environment variable is required")
	}

	return cfg
}
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
