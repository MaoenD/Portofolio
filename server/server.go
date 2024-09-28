package server

import (
	"Portefolio/HtmlLink"
	"Portefolio/database"
	"log"
	"net/http"
)

// Function that initializes and starts the server
func Start() {
	cfg := database.ReadDB()            // Read the database configuration.
	database.ConnectDB(cfg.DatabaseURL) // Connect to the database.
	database.CreateTable()              // Create the table.

	http.Handle("/Style/", http.StripPrefix("/Style/", http.FileServer(http.Dir("Style")))) // Handle the "/Style/" path.
	http.Handle("/ressource/", http.StripPrefix("/ressource/", http.FileServer(http.Dir("ressource"))))
	http.HandleFunc("/adminXP", HtmlLink.HandleAdminPage)
	http.HandleFunc("/adminFormations", HtmlLink.HandleFormationPage)
	http.HandleFunc("/index", HtmlLink.HandleIndexPage)
	http.HandleFunc("/login", HtmlLink.HandleLoginPage)

	log.Println("Listening on :8080...") // Log "Listening on :8080...".

	if err := http.ListenAndServe(":8080", nil); err != nil { // If there is an error.
		log.Fatal(err) // Log the error.
	}
}
