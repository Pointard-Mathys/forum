package forum

import (
	session "forum/FORUM/ACCOUNT"
	"net/http"
	"text/template"
)

func Mainpage(w http.ResponseWriter, r *http.Request) {
	session := session.Store
	sessions, _ := session.Get(r, "session")
	var templates = template.Must(template.ParseFiles("././static/home/home.html", "././static/home/header.html"))
	templates.ExecuteTemplate(w, "header.html", sessions.Values["userName"])
	// templates.Execute(w, nil)
}
