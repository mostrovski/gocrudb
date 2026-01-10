package database

import (
	"fmt"
	"gocrudb/config"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(appDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connection failed: %s", err.Error())
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("sql.DB retrieval failed: %s", err.Error())
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxIdleTime(10 * time.Minute)
	sqlDb.SetConnMaxLifetime(45 * time.Minute)

	return db
}

func appDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Get("app_db_host"),
		config.Get("app_db_user"),
		config.Get("app_db_password"),
		config.Get("app_db_name"),
		config.Get("app_db_port"),
		config.Get("app_db_ssl_mode"),
		config.Get("app_db_time_zone"),
	)
}
