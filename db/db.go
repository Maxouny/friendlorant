package db

import (
	"database/sql"

	"friendlorant/internal/config"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	cfg := config.LoadConfig()

	db, err := sql.Open("postgres", cfg.Database.DSN)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
