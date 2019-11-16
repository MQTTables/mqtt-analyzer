package methods

import (
	"html/template"
	"log"
	"net/http"
)

//Index - Handles main page
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Template parsing error: %s", err)
		return
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}

//Index - Handles view page
func View(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Template parsing error: %s", err)
		return
	}
	tmpl.ExecuteTemplate(w, "view", nil)
}
