package GestionBDD

import (
	"database/sql"
	"log"
)

func GetAllProjet(db *sql.DB) (string, error) {
	var allProjet string
	err := db.QueryRow("SELECT * FROM Projet").Scan(&allProjet)
	if err != nil {
		log.Fatal(err)
	}
	return allProjet, nil
}

func GetProjectNameById(db *sql.DB, id int) (string, error) {
	var projetName string
	err := db.QueryRow("SELECT Nom_Projet FROM Projet WHERE Id_Projet = ?", id).Scan(&projetName)
	if err != nil {
		log.Fatal(err)
	}
	return projetName, nil
}

func GetProjectDescriptionById(db *sql.DB, id int) (string, error) {
	var projetDescription string
	err := db.QueryRow("SELECT Description FROM Projet WHERE Id_Projet = ?", id).Scan(&projetDescription)
	if err != nil {
		log.Fatal(err)
	}
	return projetDescription, nil
}

func GetProjectStartDateById(db *sql.DB, id int) (string, error) {
	var projetStartDate string
	err := db.QueryRow("SELECT Date_Debut FROM Projet WHERE Id_Projet = ?", id).Scan(&projetStartDate)
	if err != nil {
		log.Fatal(err)
	}
	return projetStartDate, nil
}

func GetProjectEndDateById(db *sql.DB, id int) (string, error) {
	var projetEndDate string
	err := db.QueryRow("SELECT Date_Fin FROM Projet WHERE Id_Projet = ?", id).Scan(&projetEndDate)
	if err != nil {
		log.Fatal(err)
	}
	return projetEndDate, nil
}

func GetProjectSpanById(db *sql.DB, id int) (string, error) {
	var projetSpan string
	err := db.QueryRow("SELECT Duree FROM Projet WHERE Id_Projet = ?", id).Scan(&projetSpan)
	if err != nil {
		log.Fatal(err)
	}
	return projetSpan, nil
}