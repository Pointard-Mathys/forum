package forum

import (
	"html/template"
	"net/http"
)

func TestTHEME() http.HandlerFunc {
	i := 0
	return func(w http.ResponseWriter, r *http.Request) {
		c := cron.New()
		tmpl, _ := template.ParseFiles("././static/testTHEME/testTHEME.html")
		tmpl.Execute(w, r)

	}
}
