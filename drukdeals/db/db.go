package db

import (
	"database/sql"
	"os"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/drukdeals_db?sslmode=disable"
	}
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	return DB.Ping()
}
