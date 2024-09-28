package main

import (
	"Portefolio/server"
	"log"
)

// http://localhost:8080/index
// login: Guillaume pwd: 1234
// double clic arrow to admin

func main() {
	log.Println("Starting server...")
	server.Start()
}
