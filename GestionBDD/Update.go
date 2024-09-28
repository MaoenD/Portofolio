package GestionBDD

import (
	"database/sql"
	"log"
)

// Function that updates the project name by the project ID.
func UpdateProjetNameById(db *sql.DB, id int, newProjetName string) {
	_, err := db.Exec("UPDATE Projet SET Nom_Projet = ? WHERE Id_Projet = ?", newProjetName, id) // Execute the SQL query to update the project name by the project ID.
	if err != nil {                                                                              // If there is an error, log the error.
		log.Println(err) // Log the error.
	}
}

// Function that updates the project description by the project ID.
func UpdateDescriptionById(db *sql.DB, id int, newDescription string) {
	_, err := db.Exec("UPDATE Projet SET Description = ? WHERE Id_Projet = ?", newDescription, id)
	if err != nil {
		log.Println(err)
	}
}

// Function that updates the project start date by the project ID.
func UpdateDateDebutById(db *sql.DB, id int, newDateStart string) {
	_, err := db.Exec("UPDATE Projet SET Date_Debut = ? WHERE Id_Projet = ?", newDateStart, id)
	if err != nil {
		log.Println(err)
	}
}

// Function that updates the project end date by the project ID.
func UpdateDateFinById(db *sql.DB, id int, newDateFin string) {
	_, err := db.Exec("UPDATE Projet SET Date_Fin = ? WHERE Id_Projet = ?", newDateFin, id)
	if err != nil {
		log.Println(err)
	}
}

// Function that updates the project span by the project ID.
func UpdateSpanById(db *sql.DB, id int, newSpan string) {
	_, err := db.Exec("UPDATE Projet SET Duree = ? WHERE Id_Projet = ?", newSpan, id)
	if err != nil {
		log.Println(err)
	}
}

// Function that updates the formation name by the formation ID.
func UpdateFormationNameById(db *sql.DB, id int, newProjetName string) {
	_, err := db.Exec("UPDATE Formations SET Nom_Formation = ? WHERE Id_Formation = ?", newProjetName, id)
	if err != nil {
		log.Println(err)
	}
}

// Function that updates the formation description by the formation ID.
func UpdateFormationDescriptionById(db *sql.DB, id int, newDescription string) {
	_, err := db.Exec("UPDATE Formations SET Description = ? WHERE Id_Formation = ?", newDescription, id)
	if err != nil {
		log.Println(err)
	}
}

// Function that updates the formation start date by the formation ID.
func UpdateFormationDateDebutById(db *sql.DB, id int, newDateStart string) {
	_, err := db.Exec("UPDATE Formations SET Date_Debut = ? WHERE Id_Formation = ?", newDateStart, id)
	if err != nil {
		log.Println(err)
	}
}

// Function that updates the formation end date by the formation ID.
func UpdateFormationDateFinById(db *sql.DB, id int, newDateFin string) {
	_, err := db.Exec("UPDATE Formations SET Date_Fin = ? WHERE Id_Formation = ?", newDateFin, id)
	if err != nil {
		log.Println(err)
	}
}

// Function that updates the formation span by the formation ID.
func UpdateFormationSpanById(db *sql.DB, id int, newSpan string) {
	_, err := db.Exec("UPDATE Formations SET Duree = ? WHERE Id_Formation = ?", newSpan, id)
	if err != nil {
		log.Println(err)
	}
}
