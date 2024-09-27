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

var TempId string

func HandleAdminPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleAdminPage")
	if r.Method == http.MethodGet {
		datas := GetDatas()
		data := map[string]interface{}{
			"Projets": datas,
		}
		idToGet := r.FormValue("id")
		if idToGet != "" {
			id, err := strconv.Atoi(idToGet)
			if err == nil {
				dataByIdTemp := GetDatasById(id)
				data["SelectedProject"] = dataByIdTemp
			} else {
				log.Println("Invalid ID:", err)
			}
		}
		templates.RenderTemplate(w, "adminXP", data)
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
		PostData(nom, description, DateDebut, DateFin, Span)
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
		fmt.Println(TempId, nomUpdate, descriptionUpdate, DateDebutUpdate, DateFinUpdate, SpanUpdate, "caca4")
		UpdateData(TempId, nomUpdate, descriptionUpdate, DateDebutUpdate, DateFinUpdate, SpanUpdate)
		fmt.Println("caca3")
	}
}

func PostData(nom string, description string, DateStart string, DateFin string, Span string) {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	GestionBDD.PostProjet(db, nom, description, DateStart, DateFin, Span)
}

func UpdateData(id string, nom string, description string, DateStart string, DateFin string, span string) {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	idInt, _ := strconv.Atoi(id)
	GestionBDD.UpdateDateDebutById(db, idInt, DateStart)
	GestionBDD.UpdateDateFinById(db, idInt, DateFin)
	GestionBDD.UpdateDescriptionById(db, idInt, description)
	GestionBDD.UpdateProjetNameById(db, idInt, nom)
	GestionBDD.UpdateSpanById(db, idInt, span)
}

func GetDatas() []GestionBDD.Projet {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	datas, err := GestionBDD.GetAllProjet(db)
	if err != nil {
		log.Fatal(err)
	}
	return datas
}

func GetDatasById(id int) GestionBDD.Projet {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	datas, err := GestionBDD.GetProjetById(db, id)
	if err != nil {
		log.Fatal(err)
	}
	return datas
}
