package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func InitDatabase() (*sql.DB, error) {
	c := os.Getenv("DB_URL")
	if c == "" {
		log.Fatal("DB_URL not found in environment variable")
	}

	db, err = sql.Open("postgres", c)
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	return db, nil
}

func GetDB() *sql.DB {
	return db
}
