package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	DB = Connect()
	fmt.Println("\n\tDatabase connected")
}

func Connect() *sql.DB {
	connStr := "dbname=flashcard user=flashcard password=minh2801 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
