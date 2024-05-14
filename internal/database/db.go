package database

import (
	"context"
	"fmt"

	"friendlorant/internal/config"

	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jackc/pgx/v5"
)

func ConnectDB() (*pgx.Conn, error) {
	dbCfg, err := config.LoadDBConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}

	dbConn, err := pgx.Connect(
		context.Background(),
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Name, dbCfg.Password))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return dbConn, nil
}
