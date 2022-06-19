package forum

import (
	"fmt"
	database "forum/FORUM/DATABASE"
	"net/http"
	"text/template"
)

func Chat(w http.ResponseWriter, r *http.Request) {
	db := database.InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	database.InsertIntoTopic(db, "Poogie", "Message de Sa Sainteté", 0, 1)
	var templates = template.Must(template.ParseFiles("././static/messages/chat.html", "././static/home/header.html"))
	userId := 1
	messageContent := r.FormValue("chat-window-message")
	database.InsertIntoReponse(db, messageContent, 1, userId)
	fmt.Println("Réponse : ", messageContent)
	templates.Execute(w, nil)
}
