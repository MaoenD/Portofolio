package HtmlLink

import (
	"Portefolio/templates"
	"net/http"
)

func HandleLoginPage(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "login", nil)
}
