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
