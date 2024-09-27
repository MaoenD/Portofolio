package HtmlLink

import (
	"Portefolio/templates"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func HandleFormationPage(w http.ResponseWriter, r *http.Request) {
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
		templates.RenderTemplate(w, "admin", data)
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
