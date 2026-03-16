package database

import (
	"context"
	"fmt"
	"gocrudb/config"
	"gocrudb/resource"
	"log"

	"github.com/jackc/pgx/v5"
	"gorm.io/gorm"
)

func Setup() {
	if config.IsProduction() || config.ShoulSkipDbSetup() {
		return
	}

	fmt.Println("[DB SETUP] Start")

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, serviceDSN())
	if err != nil {
		log.Fatalf("[DB SETUP] Failed (connection): %v", err)
	}
	defer conn.Close(ctx)

	dbName := config.Get("app_db_name")
	dbUser := config.Get("app_db_user")
	dbPassword := config.Get("app_db_password")

	var userExists bool
	err = conn.QueryRow(ctx, "SELECT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = $1)", dbUser).Scan(&userExists)
	if err != nil {
		log.Fatalf("[DB SETUP] Failed (user check): %v", err)
	}

	var dbExists bool
	err = conn.QueryRow(ctx, "SELECT EXISTS (SELECT FROM pg_database WHERE datname = $1)", dbName).Scan(&dbExists)
	if err != nil {
		log.Fatalf("[DB SETUP] Failed (database check): %v", err)
	}

	if userExists {
		fmt.Println("[DB SETUP] User already exists, skip creation")
	} else {
		fmt.Println("[DB SETUP] Creating user")
		_, err = conn.Exec(ctx, fmt.Sprintf("CREATE ROLE %s WITH LOGIN PASSWORD '%s'", pgx.Identifier{dbUser}.Sanitize(), dbPassword))
		if err != nil {
			log.Fatalf("[DB SETUP] Failed (database create): %v", err)
		}
	}

	if dbExists {
		fmt.Println("[DB SETUP] Database already exists, skip creation")
	} else {
		fmt.Println("[DB SETUP] Creating database")
		_, err = conn.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s OWNER %s", pgx.Identifier{dbName}.Sanitize(), pgx.Identifier{dbUser}.Sanitize()))
		if err != nil {
			log.Fatalf("[DB SETUP] Failed (database create): %v", err)
		}
	}

	fmt.Println("[DB SETUP] Finish")
}

func Migrate(db *gorm.DB, resources ...any) {
	if config.IsProduction() {
		return
	}

	fmt.Println("[DB MIGRATE] Start")

	for _, r := range resources {
		db.AutoMigrate(&r)
	}

	fmt.Println("[DB MIGRATE] Finish")
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

	fmt.Printf("[DB SEED] Start %T\n", resources)

	manager.CreateInBatches(ctx, &resources, 100)

	fmt.Printf("[DB SEED] Finish %T\n", resources)
}

func serviceDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Get("service_db_host"),
		config.Get("service_db_user"),
		config.Get("service_db_password"),
		config.Get("service_db_name"),
		config.Get("service_db_port"),
		config.Get("service_db_ssl_mode"),
		config.Get("service_db_time_zone"),
	)
}
