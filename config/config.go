package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if os.Getenv("DB_RESERVATIONS_DSN") != "" && os.Getenv("QUERY_RESERVATION_URL") != "" {
		return nil
	}

	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

func GetQueryReservationURL() string {
	return os.Getenv("QUERY_RESERVATION_URL")
}

func ConnectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error verifying connection to database: %w", err)
	}
	return db, nil
}

func InitDatabases() (map[string]*sql.DB, error) {
	if err := LoadEnv(); err != nil {
		return nil, fmt.Errorf("error loading environment variables: %w", err)
	}

	dbDSN := os.Getenv("DB_RESERVATIONS_DSN")
	if dbDSN == "" {
		return nil, fmt.Errorf("missing DB_RESERVATIONS_DSN in environment variables")
	}

	db, err := ConnectDB(dbDSN)
	if err != nil {
		return nil, fmt.Errorf("error connecting to reservations database: %w", err)
	}

	return map[string]*sql.DB{"reservations": db}, nil
}
