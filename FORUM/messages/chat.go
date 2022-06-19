package forum

import (
	"fmt"
	session "forum/FORUM/ACCOUNT"
	database "forum/FORUM/DATABASE"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
)

func Chat(w http.ResponseWriter, r *http.Request) {
	session := session.Store
	sessions, _ := session.Get(r, "session")
	userID := sessions.Values["userID"].(int)
	db := database.InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	var templates = template.Must(template.ParseFiles("././static/messages/chat.html", "././static/home/home.html", "././static/home/header.html"))
	messageContent := r.FormValue("chat-window-message")
	topicId, _ := strconv.Atoi(r.FormValue("idTopic"))
	if topicId > 0 {
		fmt.Println("Topic ID : ", topicId, "\n\n\n")
		sessions.Values["topicID"] = topicId
		sessions.Save(r, w)
	}
	fmt.Println("sessions.Values Topic : ", sessions.Values["topicID"])
	fmt.Println("sessions.Values User : ", sessions.Values["userID"])
	fmt.Println("sessions.Values User name : ", sessions.Values["userName"])
	isNotEmptyOrBlank := regexp.MustCompile(`\S`)
	if messageContent != "" && isNotEmptyOrBlank.MatchString(messageContent) {
		database.InsertIntoReponse(db, messageContent, sessions.Values["topicID"].(int), userID)
	}
	// fmt.Println("Réponse : ", messageContent, " envoyée par :", ID)
	templates.Execute(w, nil)
}
