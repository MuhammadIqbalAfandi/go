package config

import (
	"restfull-api/internal/shared"
)

type Config struct {
	AppEnv   string
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

func Load() Config {
	env := shared.GetEnv("APP_ENV", "dev")

	cfg := Config{
		AppEnv: env,
	}

	switch env {
	case "prod":
		cfg.Database = DatabaseConfig{
			Host:    shared.GetEnv("DATABASE_HOST", "localhost"),
			Port:    shared.GetEnv("DATABASE_PORT", "5432"),
			User:    shared.GetEnv("DATABASE_USER", "postgres"),
			Pass:    shared.GetEnv("DATABASE_PASS", "postgres"),
			DBName:  shared.GetEnv("DATABASE_NAME", "golang_restful_api"),
			SSLMode: shared.GetEnv("DATABASE_SSLMODE", "disable"),
		}
	case "testing":
		cfg.Database = DatabaseConfig{
			Host:    shared.GetEnv("DATABASE_HOST", "localhost"),
			Port:    shared.GetEnv("DATABASE_PORT", "5432"),
			User:    shared.GetEnv("DATABASE_USER", "postgres"),
			Pass:    shared.GetEnv("DATABASE_PASS", "postgres"),
			DBName:  shared.GetEnv("DATABASE_NAME", "go_restfulapi_test"),
			SSLMode: "disable",
		}
	default:
		cfg.Database = DatabaseConfig{
			Host:    shared.GetEnv("DATABASE_HOST", "localhost"),
			Port:    shared.GetEnv("DATABASE_PORT", "5432"),
			User:    shared.GetEnv("DATABASE_USER", "postgres"),
			Pass:    shared.GetEnv("DATABASE_PASS", "postgres"),
			DBName:  shared.GetEnv("DATABASE_NAME", "go_restfulapi"),
			SSLMode: "disable",
		}
	}

	return cfg
}
