package GestionBDD

import (
	"database/sql"
	"fmt"
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
		log.Println(err)
	}
	return projet, nil
}

func GetProjectNameById(db *sql.DB, id int) (string, error) {
	var projetName string
	err := db.QueryRow("SELECT Nom_Projet FROM Projet WHERE Id_Projet = ?", id).Scan(&projetName)
	if err != nil {
		log.Println(err)
	}
	return projetName, nil
}

func GetProjectDescriptionById(db *sql.DB, id int) (string, error) {
	var projetDescription string
	err := db.QueryRow("SELECT Description FROM Projet WHERE Id_Projet = ?", id).Scan(&projetDescription)
	if err != nil {
		log.Println(err)
	}
	return projetDescription, nil
}

func GetProjectStartDateById(db *sql.DB, id int) (string, error) {
	var projetStartDate string
	err := db.QueryRow("SELECT Date_Debut FROM Projet WHERE Id_Projet = ?", id).Scan(&projetStartDate)
	if err != nil {
		log.Println(err)
	}
	return projetStartDate, nil
}

func GetProjectEndDateById(db *sql.DB, id int) (string, error) {
	var projetEndDate string
	err := db.QueryRow("SELECT Date_Fin FROM Projet WHERE Id_Projet = ?", id).Scan(&projetEndDate)
	if err != nil {
		log.Println(err)
	}
	return projetEndDate, nil
}

func GetProjectSpanById(db *sql.DB, id int) (string, error) {
	var projetSpan string
	err := db.QueryRow("SELECT Duree FROM Projet WHERE Id_Projet = ?", id).Scan(&projetSpan)
	if err != nil {
		log.Println(err)
	}
	return projetSpan, nil
}

func GetAllFormations(db *sql.DB) ([]Formation, error) {
	rows, err := db.Query("SELECT * FROM Formations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Formations []Formation
	for rows.Next() {
		var Formation Formation
		if err := rows.Scan(&Formation.Id_Formation, &Formation.Nom_Formation, &Formation.Description, &Formation.Date_Debut, &Formation.Date_Fin, &Formation.Durée); err != nil {
			return nil, err
		}
		Formations = append(Formations, Formation)
	}

	return Formations, nil
}

func GetFormationsById(db *sql.DB, id int) (Formation, error) {
	var formation Formation
	err := db.QueryRow("SELECT * FROM Formations WHERE Id_Formation = ?", id).Scan(&formation.Id_Formation, &formation.Nom_Formation, &formation.Description, &formation.Date_Debut, &formation.Date_Fin, &formation.Durée)
	if err != nil {
		if err == sql.ErrNoRows {
			return Formation{}, fmt.Errorf("aucune information trouvée pour l'ID %d", id)
		}
		return Formation{}, err
	}
	return formation, nil
}

func GetFormationsNameById(db *sql.DB, id int) (string, error) {
	var formationName string
	err := db.QueryRow("SELECT Nom_Formation FROM Formations WHERE Id_Formation = ?", id).Scan(&formationName)
	if err != nil {
		log.Println(err)
	}
	return formationName, nil
}

func GetFormationsDescriptionById(db *sql.DB, id int) (string, error) {
	var formationDescription string
	err := db.QueryRow("SELECT Description FROM Formations WHERE Id_Formation = ?", id).Scan(&formationDescription)
	if err != nil {
		log.Println(err)
	}
	return formationDescription, nil
}

func GetFormationsStartDateById(db *sql.DB, id int) (string, error) {
	var formationStartDate string
	err := db.QueryRow("SELECT Date_Debut FROM Formations WHERE Id_Formation = ?", id).Scan(&formationStartDate)
	if err != nil {
		log.Println(err)
	}
	return formationStartDate, nil
}

func GetFormationsEndDateById(db *sql.DB, id int) (string, error) {
	var formationEndDate string
	err := db.QueryRow("SELECT Date_Fin FROM Formations WHERE Id_Formation = ?", id).Scan(&formationEndDate)
	if err != nil {
		log.Println(err)
	}
	return formationEndDate, nil
}

func GetFormationsSpanById(db *sql.DB, id int) (string, error) {
	var formationSpan string
	err := db.QueryRow("SELECT Duree FROM Formations WHERE Id_Formation = ?", id).Scan(&formationSpan)
	if err != nil {
		log.Println(err)
	}
	return formationSpan, nil
}
