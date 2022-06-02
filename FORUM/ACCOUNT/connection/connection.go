package forum

import (
	"html/template"
	"net/http"
)

func ConnectionPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./FORUM/ACCOUNT/connection/connection.html")
		tmpl.Execute(w, r)
	}
}
