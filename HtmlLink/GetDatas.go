package HtmlLink

import (
	"Portefolio/GestionBDD"
	"Portefolio/templates"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func HandleAdminPage(w http.ResponseWriter, r *http.Request) {
	datas := GetDatas(w, r)
	data := map[string]interface{}{
		"Projets": datas,
	}
	fmt.Println("HandleAdminPage", datas)
	if r.Method == http.MethodGet {
		templates.RenderTemplate(w, "admin", data)
		return
	}
}

func GetDatas(w http.ResponseWriter, r *http.Request) []GestionBDD.Projet {

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
