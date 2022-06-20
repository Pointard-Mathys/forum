package forum

import (
	session "forum/FORUM/ACCOUNT"
	database "forum/FORUM/DATABASE"
	"net/http"
	"text/template"
)

func TopicPage(w http.ResponseWriter, r *http.Request) {
	session := session.Store
	sessions, _ := session.Get(r, "session")
	ID := sessions.Values["userID"].(int)
	var templates = template.Must(template.ParseFiles("././static/create-topic/create-topic.html", "././static/home/header.html"))
	db := database.InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")

	title := r.FormValue("title")
	description := r.FormValue("description")
	if title != "" && description != "" {
		database.InsertIntoTopic(db, title, description, 0, ID)
		http.Redirect(w, r, "/", http.StatusFound)
	}
	templates.Execute(w, nil)
}
