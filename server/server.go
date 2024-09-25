package server

import (
	"Portefolio/GestionBDD"
	"Portefolio/HtmlLink"
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

	GestionBDD.PostProjet(db, "test", "test", "test", "test", "test")

	http.HandleFunc("/admin", HtmlLink.HandleAdminPage)

	log.Println("Listening on :8080...")
	// Starting the server on port 8080

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	// Starting the HTTP server.
}
