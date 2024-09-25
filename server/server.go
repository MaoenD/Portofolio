package server

import (
	"Portefolio/GestionBDD"
	"Portefolio/database"
	"database/sql"
	"log"
	"net/http"
)

// Function that initializes and starts the server
func Start() {
	cfg := database.ReadDB()
	database.ConnectDB(cfg.DatabaseURL)
	database.CreateTable()
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	ProjetId, err := GestionBDD.GetProjetById(db, 1)
	log.Println("ID : ", ProjetId.Id_Projet, ProjetId.Nom_Projet)
	log.Println("Listening on :8080...")
	// Starting the server on port 8080

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	// Starting the HTTP server.
}
