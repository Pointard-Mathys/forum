package forum

import (
	"net/http"
	"text/template"
)

func Chat(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("././static/messages/chat.html", "././static/home/header.html"))
	templates.Execute(w, nil)
}
