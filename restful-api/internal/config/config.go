package config

import "os"

type Config struct {
	AppPort  string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host    string
	Port    string
	User    string
	Pass    string
	DBName  string
	SSLMode string
}

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}

func Load() Config {
	return Config{
		AppPort: getEnv("APP_PORT", "8080"),
		Database: DatabaseConfig{
			Host:    getEnv("DATABASE_HOST", "localhost"),
			Port:    getEnv("DATABASE_PORT", "5432"),
			User:    getEnv("DATABASE_USER", "postgres"),
			Pass:    getEnv("DATABASE_PASS", "postgres"),
			DBName:  getEnv("DATABASE_NAME", "golang_restful_api"),
			SSLMode: getEnv("DATABASE_SSLMODE", "disable"),
		},
	}
}
