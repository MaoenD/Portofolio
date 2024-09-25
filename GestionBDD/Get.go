package GestionBDD

import (
	"database/sql"
	"log"
)

func GetAllProjet(db *sql.DB) ([]Projet, error) {
	rows, err := db.Query("SELECT * FROM Projet")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Projets []Projet
	for rows.Next() {
		var Projet Projet
		if err := rows.Scan(&Projet.Id_Projet, &Projet.Nom_Projet, &Projet.Description, &Projet.Date_Debut, &Projet.Date_Fin, &Projet.Durée); err != nil {
			return nil, err
		}
		Projets = append(Projets, Projet)
	}

	return Projets, nil
}

func GetProjetById(db *sql.DB, id int) (Projet, error) {
	var projet Projet
	err := db.QueryRow("SELECT * FROM Projet WHERE Id_Projet = ?", id).Scan(&projet.Id_Projet, &projet.Nom_Projet, &projet.Description, &projet.Date_Debut, &projet.Date_Fin, &projet.Durée)
	if err != nil {
		log.Fatal(err)
	}
	return projet, nil
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
