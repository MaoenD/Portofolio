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
var loged bool = false

func HandleAdminPage(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	user, pass := GetLogins()
	if username == user && password == pass {
		loged = true
	}
	if loged {
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
			nom := r.FormValue("Nom_Projet")
			description := r.FormValue("Description")
			DateDebut := r.FormValue("Date_Debut")
			DateFin := r.FormValue("Date_Fin")
			Span := r.FormValue("Durée")
			if nom != "" && description != "" && DateDebut != "" && DateFin != "" && Span != "" {
				PostData(nom, description, DateDebut, DateFin, Span)
			}
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
				UpdateData(TempId, nomUpdate, descriptionUpdate, DateDebutUpdate, DateFinUpdate, SpanUpdate)
			}
		}
	}
}

func PostData(nom string, description string, DateStart string, DateFin string, Span string) {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Println(err)
	}
	GestionBDD.PostProjet(db, nom, description, DateStart, DateFin, Span)
}

func UpdateData(id string, nom string, description string, DateStart string, DateFin string, span string) {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Println(err)
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
		log.Println(err)
	}
	datas, err := GestionBDD.GetAllProjet(db)
	if err != nil {
		log.Println(err)
	}
	return datas
}

func GetDatasById(id int) GestionBDD.Projet {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Println(err)
	}
	datas, err := GestionBDD.GetProjetById(db, id)
	if err != nil {
		log.Println(err)
	}
	return datas
}

func GetLogins() (user string, pass string) {
	var username, password string
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	username, password, err = GestionBDD.Getlogins(db)
	if err != nil {
		log.Println(err)
	}
	return username, password
}
