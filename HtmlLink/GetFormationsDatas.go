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

// Function that handles the formation page.
func HandleFormationPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleAdminPage") // Print the function name.
	if loged {
		if r.Method == http.MethodGet {
			datas := GetFormationsDatas()   // Get the project data.
			data := map[string]interface{}{ // Declare the data variable as a map.
				"Formations": datas, // Declare the data variable as a map with the key "Formations" and the value datas.
			}
			idToGet := r.FormValue("id") // Get the project ID from the form.
			if idToGet != "" {
				id, err := strconv.Atoi(idToGet) // Convert the project ID to an integer.
				if err == nil {
					dataByIdTemp := GetFormationsDatasById(id) // Get the project data by the project ID.
					data["SelectedFormation"] = dataByIdTemp   // Declare the data variable as a map with the key "SelectedFormation" and the value dataByIdTemp.
				} else {
					log.Println("Invalid ID:", err) // Log the error.
				}
			}
			templates.RenderTemplate(w, "adminFormations", data) // Render the template with the data.
			return
		}
		if r.Method == http.MethodPost { // If the method is POST.
			nom := r.FormValue("Nom_Projet")                                                      // Get the project name from the form.
			description := r.FormValue("Description")                                             // Get the project description from the form.
			DateDebut := r.FormValue("Date_Debut")                                                // Get the project start date from the form.
			DateFin := r.FormValue("Date_Fin")                                                    // Get the project end date from the form.
			Span := r.FormValue("Durée")                                                          // Get the project span from the form.
			if nom != "" && description != "" && DateDebut != "" && DateFin != "" && Span != "" { // If the project name, description, start date, end date, and span are not empty.
				PostFormationsData(nom, description, DateDebut, DateFin, Span) // Post the project data.
			}
			log.Print("test")       // Print "test".
			id := r.FormValue("id") // Get the project ID from the form.
			if id != "" {           // If the project ID is not empty.
				TempId = id // Set the TempId to the project ID.
			}
			nomUpdate := r.FormValue("updatePfName")
			descriptionUpdate := r.FormValue("updatePfDescription")                                                             // Get the project description from the form.
			DateDebutUpdate := r.FormValue("updateStartDate")                                                                   // Get the project start date from the form.
			DateFinUpdate := r.FormValue("updateEndDate")                                                                       // Get the project end date from the form.
			SpanUpdate := r.FormValue("updateDuration")                                                                         // Get the project span from the form.
			if nomUpdate != "" && descriptionUpdate != "" && DateDebutUpdate != "" && DateFinUpdate != "" && SpanUpdate != "" { // If the project name, description, start date, end date, and span are not empty.
				UpdateFormationsData(TempId, nomUpdate, descriptionUpdate, DateDebutUpdate, DateFinUpdate, SpanUpdate) // Update the project data.
			}
		}
	}
}

// Function that posts the formation data.
func PostFormationsData(nom string, description string, DateStart string, DateFin string, Span string) {
	db, err := sql.Open("sqlite3", "database/database.db") // Open the database.
	if err != nil {
		log.Println(err)
	}
	GestionBDD.PostFormations(db, nom, description, DateStart, DateFin, Span) // Post the project data.
}

// Function that gets the formation data.
func GetFormationsDatas() []GestionBDD.Formation {
	db, err := sql.Open("sqlite3", "database/database.db") // Open the database.
	if err != nil {
		log.Println(err) // Log the error.
	}
	datas, err := GestionBDD.GetAllFormations(db) // Get all the project data.
	if err != nil {
		log.Println(err)
	}
	return datas // Return the project data.
}

// Function that gets the formation data by the formation ID.
func GetFormationsDatasById(id int) GestionBDD.Formation {
	db, err := sql.Open("sqlite3", "database/database.db") // Open the database.
	if err != nil {
		log.Println(err)
	}
	datas, err := GestionBDD.GetFormationsById(db, id) // Get the project data by the project ID.
	if err != nil {
		log.Println(err)
	}
	return datas // Return the project data.
}

// Function that updates the formation data.²
func UpdateFormationsData(id string, nom string, description string, DateStart string, DateFin string, span string) {
	db, err := sql.Open("sqlite3", "database/database.db") // Open the database.
	if err != nil {
		log.Print(err)
	}
	idInt, _ := strconv.Atoi(id)                                  // Convert the project ID to an integer.
	GestionBDD.UpdateFormationDateDebutById(db, idInt, DateStart) // Update the project start date by the project ID.
	GestionBDD.UpdateFormationDateFinById(db, idInt, DateFin)     // Update the project end date by the project ID.
	GestionBDD.UpdateFormationDescriptionById(db, idInt, description)
	GestionBDD.UpdateFormationNameById(db, idInt, nom)
	GestionBDD.UpdateFormationSpanById(db, idInt, span)
}
