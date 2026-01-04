package config

import (
	"fmt"
	"os"
	"strings"
)

func PostgresDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "gocrudb"),
		getEnv("DB_PASSWORD", "gocrudb"),
		getEnv("DB_NAME", "gocrudb"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_SSL_MODE", "disable"),
		getEnv("DB_TIME_ZONE", "Europe/Berlin"),
	)
}

func AppEnv() string {
	return getEnv("APP_ENV", "dev")
}

func AppPort() string {
	return getEnv("APP_PORT", ":3000")
}

func IsProduction() bool {
	return strings.Contains(AppEnv(), "prod")
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
