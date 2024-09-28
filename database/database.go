package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct { // Config is the struct that holds the database configuration.
	DatabaseURL string
}

var DB *sql.DB // DB is the global variable that holds the database connection.

// ReadDB reads and returns the database configuration, which is read from the environment variables.
func ReadDB() Config {
	fmt.Println("Reading DB")
	return Config{
		DatabaseURL: "database/database.db",
	}
}

// ConnectDB connects to the database using the provided database URL and stores the connection in the global variable DB.
func ConnectDB(databaseURL string) {
	var err error
	DB, err = sql.Open("sqlite3", databaseURL)
	if err != nil {
		log.Println(err)
	}

	if err = DB.Ping(); err != nil {
		log.Println(err)
	}

	log.Println("Database connected!")
}

// Creates the table in the database if it does not exist.
func CreateTable() {
	fmt.Println("Creating table")
	query := `
    CREATE TABLE IF NOT EXISTS Projet(
        Id_Projet INTEGER PRIMARY KEY AUTOINCREMENT,
        Nom_Projet TEXT NOT NULL,
        Description TEXT,
        Date_Debut TEXT,
        Date_Fin TEXT,
    	Duree TEXT
    );
		
	CREATE TABLE IF NOT EXISTS Formations (
        id_Formation INTEGER PRIMARY KEY AUTOINCREMENT,
        Nom_Formation TEXT NOT NULL,
        description TEXT,
        Date_Debut TEXT,
        Date_Fin TEXT,
        Duree TEXT
    );

	CREATE TABLE IF NOT EXISTS Logins (
	    id INTEGER PRIMARY KEY,
	    username TEXT,
	    password TEXT
	);
`

	_, err := DB.Exec(query)
	if err != nil {
		log.Println(err)
	}

	log.Println("Table created!")
}
