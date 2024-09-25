package GestionBDD

import (
	"database/sql"
	"log"
)

func PostProjet(db *sql.DB, NomProjet string, Description string, DateDebut string, DateFin string, Duree string) {
	_, err := db.Exec("INSERT INTO Projet ( Nom_Projet, Description, Date_Debut, Date_Fin, Duree) VALUES ( ?, ?, ?, ?, ?)", NomProjet, Description, DateDebut, DateFin, Duree)
	if err != nil {
		log.Fatal(err)
	}
}
