package server

import (
	"Portefolio/HtmlLink"
	"Portefolio/database"
	"log"
	"net/http"
)

// Function that initializes and starts the server
func Start() {
	cfg := database.ReadDB()
	database.ConnectDB(cfg.DatabaseURL)
	database.CreateTable()

	http.Handle("/Style/", http.StripPrefix("/Style/", http.FileServer(http.Dir("Style"))))
	http.Handle("/ressource/", http.StripPrefix("/ressource/", http.FileServer(http.Dir("ressource"))))

	http.HandleFunc("/adminXP", HtmlLink.HandleAdminPage)
	http.HandleFunc("/adminFormations", HtmlLink.HandleFormationPage)
	http.HandleFunc("/index", HtmlLink.HandleIndexPage)

	log.Println("Listening on :8080...")
	// Starting the server on port 8080

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	// Starting the HTTP server.
}
