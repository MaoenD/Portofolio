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

func HandleFormationPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleAdminPage")
	if r.Method == http.MethodGet {
		datas := GetFormationsDatas()
		data := map[string]interface{}{
			"Formations": datas,
		}
		idToGet := r.FormValue("id")
		if idToGet != "" {
			id, err := strconv.Atoi(idToGet)
			if err == nil {
				dataByIdTemp := GetFormationsDatasById(id)
				data["SelectedFormation"] = dataByIdTemp
			} else {
				log.Println("Invalid ID:", err)
			}
		}
		templates.RenderTemplate(w, "adminFormations", data)
		return
	}
	if r.Method == http.MethodPost {
		fmt.Println("Post method")
		nom := r.FormValue("Nom_Projet")
		description := r.FormValue("Description")
		DateDebut := r.FormValue("Date_Debut")
		DateFin := r.FormValue("Date_Fin")
		Span := r.FormValue("Durée")
		fmt.Println("caca")
		if nom != "" && description != "" && DateDebut != "" && DateFin != "" && Span != "" {
			PostFormationsData(nom, description, DateDebut, DateFin, Span)
		}
		fmt.Println("caca2")
		log.Print("test")
		id := r.FormValue("id")
		if id != "" {
			TempId = id
		}
		nomUpdate := r.FormValue("updatePfName")
		descriptionUpdate := r.FormValue("updatePfDescription")
		DateDebutUpdate := r.FormValue("updateStartDate")
		DateFinUpdate := r.FormValue("updateEndDate")
		SpanUpdate := r.FormValue("updateDuration")
		if nomUpdate != "" && descriptionUpdate != "" && DateDebutUpdate != "" && DateFinUpdate != "" && SpanUpdate != "" {
			UpdateFormationsData(TempId, nomUpdate, descriptionUpdate, DateDebutUpdate, DateFinUpdate, SpanUpdate)
		}
		fmt.Println("caca3")
	}
}

func PostFormationsData(nom string, description string, DateStart string, DateFin string, Span string) {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	GestionBDD.PostFormations(db, nom, description, DateStart, DateFin, Span)
}

func GetFormationsDatas() []GestionBDD.Formation {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	datas, err := GestionBDD.GetAllFormations(db)
	if err != nil {
		log.Fatal(err)
	}
	return datas
}

func GetFormationsDatasById(id int) GestionBDD.Formation {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	datas, err := GestionBDD.GetFormationsById(db, id)
	if err != nil {
		log.Fatal(err)
	}
	return datas
}

func UpdateFormationsData(id string, nom string, description string, DateStart string, DateFin string, span string) {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	idInt, _ := strconv.Atoi(id)
	GestionBDD.UpdateFormationDateDebutById(db, idInt, DateStart)
	GestionBDD.UpdateFormationDateFinById(db, idInt, DateFin)
	GestionBDD.UpdateFormationDescriptionById(db, idInt, description)
	GestionBDD.UpdateFormationNameById(db, idInt, nom)
	GestionBDD.UpdateFormationSpanById(db, idInt, span)
}
