package main

import (
	"Portefolio/server"
	"log"
)

// http://localhost:8080/adminXP
// http://localhost:8080/adminFormations
// http://localhost:8080/index

func main() {
	log.Println("Starting server...")
	server.Start()
}
