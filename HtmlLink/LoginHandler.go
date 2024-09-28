package HtmlLink

import (
	"Portefolio/templates"
	"net/http"
)

// Function that handles the login page.
func HandleLoginPage(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "login", nil) // Render the template with the data.
}
