package HtmlLink

import (
	"Portefolio/templates"
	"net/http"
)

func HandleIndexPage(w http.ResponseWriter, r *http.Request) {
	loged = false
	datas := GetDatas()
	FDatas := GetFormationsDatas()
	data := map[string]interface{}{
		"Projets":   datas,
		"Formation": FDatas,
	}
	if r.Method == http.MethodGet {
	}
	templates.RenderTemplate(w, "index", data)
	return
}
