package HtmlLink

import (
	"Portefolio/GestionBDD"
	"Portefolio/templates"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var TempId string      // Declare the TempId variable as a string.
var loged bool = false // Declare the loged variable as a boolean.

// Function that handles the admin page.
func HandleAdminPage(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username") // Get the username from the form.
	password := r.FormValue("password")
	user, pass := GetLogins()                 // Get the login information.
	if username == user && password == pass { // If the username and password are correct.
		loged = true // Set the loged variable to true.
	}
	if loged { // If the user is logged in.
		fmt.Println("HandleAdminPage") // Print the function name.
		if r.Method == http.MethodGet {
			datas := GetDatas()             // Get the project data.
			data := map[string]interface{}{ // Declare the data variable as a map.
				"Projets": datas, // Declare the data variable as a map with the key "Projets" and the value datas.
			}
			idToGet := r.FormValue("id") // Get the project ID from the form.
			if idToGet != "" {
				id, err := strconv.Atoi(idToGet) // Convert the project ID to an integer.
				if err == nil {
					dataByIdTemp := GetDatasById(id)       // Get the project data by the project ID.
					data["SelectedProject"] = dataByIdTemp // Declare the data variable as a map with the key "SelectedProject" and the value dataByIdTemp.
				} else {
					log.Println("Invalid ID:", err) // Log the error.
				}
			}
			templates.RenderTemplate(w, "adminXP", data) // Render the template with the data.
			return
		}
		if r.Method == http.MethodPost {
			nom := r.FormValue("Nom_Projet") // Get the project name from the form.
			description := r.FormValue("Description")
			DateDebut := r.FormValue("Date_Debut")
			DateFin := r.FormValue("Date_Fin")
			Span := r.FormValue("Durée")
			if nom != "" && description != "" && DateDebut != "" && DateFin != "" && Span != "" { // If the project name, description, start date, end date, and span are not empty.
				PostData(nom, description, DateDebut, DateFin, Span)
			}
			id := r.FormValue("id")
			if id != "" {
				TempId = id
			}
			nomUpdate := r.FormValue("updatePfName") // Get the project name from the form.
			descriptionUpdate := r.FormValue("updatePfDescription")
			DateDebutUpdate := r.FormValue("updateStartDate")
			DateFinUpdate := r.FormValue("updateEndDate")
			SpanUpdate := r.FormValue("updateDuration")
			if nomUpdate != "" && descriptionUpdate != "" && DateDebutUpdate != "" && DateFinUpdate != "" && SpanUpdate != "" { // If the project name, description, start date, end date, and span are not empty.
				UpdateData(TempId, nomUpdate, descriptionUpdate, DateDebutUpdate, DateFinUpdate, SpanUpdate) // Update the project data.
			}
		}
	}
}

// Function that posts the project data.
func PostData(nom string, description string, DateStart string, DateFin string, Span string) {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Println(err)
	}
	GestionBDD.PostProjet(db, nom, description, DateStart, DateFin, Span) // Post the project data.
}

// Function that update the project data.
func UpdateData(id string, nom string, description string, DateStart string, DateFin string, span string) {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Println(err)
	}
	idInt, _ := strconv.Atoi(id)
	GestionBDD.UpdateDateDebutById(db, idInt, DateStart) // Update the project data.
	GestionBDD.UpdateDateFinById(db, idInt, DateFin)
	GestionBDD.UpdateDescriptionById(db, idInt, description)
	GestionBDD.UpdateProjetNameById(db, idInt, nom)
	GestionBDD.UpdateSpanById(db, idInt, span)
}

// Function that gets the project data.
func GetDatas() []GestionBDD.Projet {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Println(err)
	}
	datas, err := GestionBDD.GetAllProjet(db) // Get all the project data.
	if err != nil {
		log.Println(err)
	}
	return datas
}

// Function that gets the project data by the project ID.
func GetDatasById(id int) GestionBDD.Projet {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Println(err)
	}
	datas, err := GestionBDD.GetProjetById(db, id) // Get the project data by the project ID.
	if err != nil {
		log.Println(err)
	}
	return datas
}

// Function that gets the login information.
func GetLogins() (user string, pass string) {
	var username, password string
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()                                   // Close the database.
	username, password, err = GestionBDD.Getlogins(db) // Get the login information.
	if err != nil {
		log.Println(err)
	}
	return username, password // Return the login information.
}
