package GestionBDD

import (
	"database/sql"
	"log"
)

func UpdateProjetNameById(db *sql.DB, id int, newProjetName string) {
	_, err := db.Exec("UPDATE Projet SET Nom_Projet = ? WHERE Id_Projet = ?", newProjetName, id)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateDescriptionById(db *sql.DB, id int, newDescription string) {
	_, err := db.Exec("UPDATE Projet SET Description = ? WHERE Id_Projet = ?", newDescription, id)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateDateDebutById(db *sql.DB, id int, newDateStart string) {
	_, err := db.Exec("UPDATE Projet SET Date_Debut = ? WHERE Id_Projet = ?", newDateStart, id)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateDateFinById(db *sql.DB, id int, newDateFin string) {
	_, err := db.Exec("UPDATE Projet SET Date_Fin = ? WHERE Id_Projet = ?", newDateFin, id)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateSpanById(db *sql.DB, id int, newSpan string) {
	_, err := db.Exec("UPDATE Projet SET Duree = ? WHERE Id_Projet = ?", newSpan, id)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateFormationNameById(db *sql.DB, id int, newProjetName string) {
	_, err := db.Exec("UPDATE Formations SET Nom_Formation = ? WHERE Id_Formation = ?", newProjetName, id)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateFormationDescriptionById(db *sql.DB, id int, newDescription string) {
	_, err := db.Exec("UPDATE Formations SET Description = ? WHERE Id_Formation = ?", newDescription, id)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateFormationDateDebutById(db *sql.DB, id int, newDateStart string) {
	_, err := db.Exec("UPDATE Formations SET Date_Debut = ? WHERE Id_Formation = ?", newDateStart, id)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateFormationDateFinById(db *sql.DB, id int, newDateFin string) {
	_, err := db.Exec("UPDATE Formations SET Date_Fin = ? WHERE Id_Formation = ?", newDateFin, id)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateFormationSpanById(db *sql.DB, id int, newSpan string) {
	_, err := db.Exec("UPDATE Formations SET Duree = ? WHERE Id_Formation = ?", newSpan, id)
	if err != nil {
		log.Fatal(err)
	}
}
