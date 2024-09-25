package server

import (
	"Portefolio/database"
	"log"
	"net/http"
)

// Function that initializes and starts the server
func Start() {
	cfg := database.ReadDB()
	database.ConnectDB(cfg.DatabaseURL)
	database.CreateTable()

	log.Println("Listening on :8080...")
	// Starting the server on port 8080

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	// Starting the HTTP server.
}
