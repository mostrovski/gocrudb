package config

import (
	"fmt"
	"gocrudb/utils"
	"strings"
)

func PostgresDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		utils.GetEnv("DB_HOST", "localhost"),
		utils.GetEnv("DB_USER", "gocrudb"),
		utils.GetEnv("DB_PASSWORD", "gocrudb"),
		utils.GetEnv("DB_NAME", "gocrudb"),
		utils.GetEnv("DB_PORT", "5432"),
		utils.GetEnv("DB_SSL_MODE", "disable"),
		utils.GetEnv("DB_TIME_ZONE", "Europe/Berlin"),
	)
}

func AppEnv() string {
	return utils.GetEnv("APP_ENV", "dev")
}

func IsProduction() bool {
	return strings.Contains(AppEnv(), "prod")
}
