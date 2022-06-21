package forum

import (
	"fmt"
	session "forum/FORUM/ACCOUNT"
	database "forum/FORUM/DATABASE"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
)

func Chat(w http.ResponseWriter, r *http.Request) {
	session := session.Store
	sessions, _ := session.Get(r, "session")
	userID := 0
	if sessions.Values["Connected?"] == true {
		userID = sessions.Values["userID"].(int)
		fmt.Println(userID)
	}
	db := database.InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	var templates = template.Must(template.ParseFiles("././static/messages/chat.html", "././static/home/home.html", "././static/home/header.html"))
	messageContent := r.FormValue("chat-window-message")
	topicId, _ := strconv.Atoi(URLParams(r))
	if topicId > 0 {
		fmt.Println("Topic ID : ", topicId, "\n\n\n")
		sessions.Values["topicID"] = topicId
		sessions.Save(r, w)
	}
	fmt.Println("\n\nsessions.Values Topic : ", sessions.Values["topicID"])
	fmt.Println("sessions.Values User : ", sessions.Values["userID"])
	fmt.Println("sessions.Values User name : ", sessions.Values["userName"])
	isNotEmptyOrBlank := regexp.MustCompile(`\S`)
	if messageContent != "" && isNotEmptyOrBlank.MatchString(messageContent) {
		fmt.Println(messageContent)
		fmt.Println("topicID :", sessions.Values["topicID"].(int))
		database.InsertIntoReponse(db, messageContent, sessions.Values["topicID"].(int), userID)
	}
	URLParams(r)
	templates.Execute(w, topicId)
}

func URLParams(r *http.Request) string {
	keys, ok := r.URL.Query()["topicId"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'topicId' is missing")
		return ""
	}

	key := keys[0]

	return key
}
