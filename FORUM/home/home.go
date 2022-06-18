package forum

import (
	"net/http"
	"text/template"
)

func Mainpage(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("././static/home/home.html", "././static/home/header.html"))
	templates.Execute(w, nil)
}
