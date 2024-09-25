package server

import (
	"log"
	"net/http"
)

// Start initializes and starts the server
func Start() {

	log.Println("Listening on :8080...")
	// Starting the server on port 8080

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	// Starting the HTTP server.
}
