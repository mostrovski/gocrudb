package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var conf = map[string]string{}

func Set() {
	if len(conf) > 0 {
		return
	}
	godotenv.Load()
	conf["app_env"] = getEnv("APP_ENV", "dev")
	conf["app_port"] = getEnv("APP_PORT", "3000")
	conf["app_dsn"] = postgresDSN()
}

func Get(key string) string {
	if value, exists := conf[key]; exists {
		return value
	}
	return ""
}

func IsProduction() bool {
	return strings.Contains(Get("app_env"), "prod")
}

func postgresDSN() string {
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

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
