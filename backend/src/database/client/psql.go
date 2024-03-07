package client

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func GetPsqlClient() (*sql.DB, error) {
	if db == nil {
		uri := os.Getenv("POSTGRES_URI")
		if uri == "" {
			return nil, fmt.Errorf("POSTGRES_URI environment variable not set")
		}
		database, err := sql.Open("postgres", uri)
		if err != nil {
			return nil, err
		}
		db = database
	}
	return db, nil
}
