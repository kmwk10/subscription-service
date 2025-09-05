package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kmwk10/subscription-service/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("db is unreachable: %v", err)
	}

	log.Println("Connected to database")
	return db
}
