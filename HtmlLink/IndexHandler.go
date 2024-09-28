package HtmlLink

import (
	"Portefolio/templates"
	"net/http"
)

// Function that handles the index page.
func HandleIndexPage(w http.ResponseWriter, r *http.Request) {
	loged = false       // Set the loged variable to false.
	datas := GetDatas() // Declare as getdata.
	FDatas := GetFormationsDatas()
	data := map[string]interface{}{ // Declare the data variable as a map.
		"Projets":   datas, // Declare the data variable as a map with the key "Projets" and the value datas.
		"Formation": FDatas,
	}
	if r.Method == http.MethodGet {
	}
	templates.RenderTemplate(w, "index", data) // Render the template with the data.
	return
}
