package HtmlLink

import (
	"Portefolio/templates"
	"net/http"
)

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
