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

func HandleAdminPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleAdminPage")
	datas := GetDatas()
	data := map[string]interface{}{
		"Projets": datas,
	}
	if r.Method == http.MethodGet {
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
		templates.RenderTemplate(w, "admin", data)
		return
	}
	if r.Method == http.MethodPost {
		nom := r.FormValue("Nom_Projet")
		description := r.FormValue("Description")
		DateDebut := r.FormValue("Date_Debut")
		DateFin := r.FormValue("Date_Fin")
		Span := r.FormValue("Durée")
		PostData(nom, description, DateDebut, DateFin, Span)
	}
}

func PostData(nom string, description string, DateStart string, DateFin string, Span string) {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	GestionBDD.PostProjet(db, nom, description, DateStart, DateFin, Span)
}

func HandleIndexPage(w http.ResponseWriter, r *http.Request) {
	datas := GetDatas()
	data := map[string]interface{}{
		"Projets": datas,
	}
	if r.Method == http.MethodGet {
	}
	templates.RenderTemplate(w, "index", data)
	return
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
