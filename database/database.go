package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func SqlTable() {
	query := `
    CREATE TABLE IF NOT EXISTS users(
        id TEXT PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        admin BOOLEAN NOT NULL,
        modo BOOLEAN NOT NULL
    );`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database migrated!")
}
