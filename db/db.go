package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}
	log.Println("Connected to database")

	return db, nil
}

func Migrate(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create driver: %v", err)
	}

	mgr, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations/",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create manager: %v", err)
	}

	if err := mgr.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %v", err)
	}
	log.Println("Database migrated")

	return nil
}
