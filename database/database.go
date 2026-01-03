package database

import (
	"context"
	"gocrudb/config"
	"gocrudb/resource"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.PostgresDSN()), &gorm.Config{})
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

func Migrate(db *gorm.DB, resources ...any) {
	if config.IsProduction() {
		return
	}

	for _, r := range resources {
		db.AutoMigrate(&r)
	}
}

func Seed[I resource.IdType, R resource.Resource[I]](db *gorm.DB, resources []R) {
	if config.IsProduction() {
		return
	}

	ctx := context.Background()
	manager := gorm.G[R](db)

	count, err := manager.Count(ctx, "")
	if err != nil {
		return
	}
	if count > 0 {
		return
	}

	manager.CreateInBatches(ctx, &resources, 100)
}
