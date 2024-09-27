package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Config struct {
	DatabaseURL string
}

var DB *sql.DB

func ReadDB() Config {
	fmt.Println("Reading DB")
	return Config{
		DatabaseURL: "database/database.db",
	}
}

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
