package GestionBDD

import (
	"database/sql"
	"log"
)

// PostProjet is the function that posts the project information to the database.
func PostProjet(db *sql.DB, NomProjet string, Description string, DateDebut string, DateFin string, Duree string) {
	_, err := db.Exec("INSERT INTO Projet ( Nom_Projet, Description, Date_Debut, Date_Fin, Duree) VALUES ( ?, ?, ?, ?, ?)", NomProjet, Description, DateDebut, DateFin, Duree) // Execute the SQL query to insert the project information into the database.
	if err != nil {                                                                                                                                                            // If there is an error, log the error.
		log.Println(err) // Log the error.
	}
}

// PostFormations is the function that posts the formation information to the database.
func PostFormations(db *sql.DB, NomProjet string, Description string, DateDebut string, DateFin string, Duree string) {
	_, err := db.Exec("INSERT INTO Formations (Nom_Formation, Description, Date_Debut, Date_Fin, Duree) VALUES ( ?, ?, ?, ?, ?)", NomProjet, Description, DateDebut, DateFin, Duree)
	if err != nil {
		log.Println(err)
	}
}
