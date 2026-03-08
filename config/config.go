package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var conf = map[string]string{}
var configured = false
var defaults = map[string]string{
	"db_host":     "localhost",
	"db_port":     "5432",
	"db_ssl_mode": "disable",
	"time_zone":   "Europe/Berlin",
}

func Set() {
	if configured {
		return
	}

	godotenv.Load()
	conf["app_name"] = getEnv("APP_NAME", "gocrudb API")
	conf["app_env"] = getEnv("APP_ENV", "dev")
	conf["app_port"] = getEnv("APP_PORT", "3000")
	conf["app_db_skip_setup"] = getEnv("DB_SKIP_SETUP", "false")
	conf["app_db_host"] = getEnv("DB_HOST", defaults["db_host"])
	conf["app_db_user"] = getEnv("DB_USER", "gocrudb")
	conf["app_db_password"] = getEnv("DB_PASSWORD", "gocrudb")
	conf["app_db_name"] = getEnv("DB_NAME", "gocrudb")
	conf["app_db_port"] = getEnv("DB_PORT", defaults["db_port"])
	conf["app_db_ssl_mode"] = getEnv("DB_SSL_MODE", defaults["db_ssl_mode"])
	conf["app_db_time_zone"] = getEnv("DB_TIME_ZONE", defaults["time_zone"])
	conf["service_db_host"] = getEnv("SERVICE_DB_HOST", defaults["db_host"])
	conf["service_db_user"] = getEnv("SERVICE_DB_USER", "postgres")
	conf["service_db_password"] = getEnv("SERVICE_DB_PASSWORD", "")
	conf["service_db_name"] = getEnv("SERVICE_DB_NAME", "postgres")
	conf["service_db_port"] = getEnv("SERVICE_DB_PORT", defaults["db_port"])
	conf["service_db_ssl_mode"] = getEnv("SERVICE_DB_SSL_MODE", defaults["db_ssl_mode"])
	conf["service_db_time_zone"] = getEnv("SERVICE_DB_TIME_ZONE", defaults["time_zone"])
	conf["rate_limiter_requests_per_second"] = getEnv("RATE_LIMITER_REQUESTS_PER_SECOND", "1")
	conf["rate_limiter_requests_burst_size"] = getEnv("RATE_LIMITER_REQUESTS_BURST_SIZE", "5")

	configured = true
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

func ShoulSkipDbSetup() bool {
	return Get("app_db_skip_setup") == "true"
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
