package main

import (
	"encoding/json"
	"fmt"
	account "forum/FORUM/ACCOUNT"
	database "forum/FORUM/DATABASE"
	mainpage "forum/FORUM/home"
	chat "forum/FORUM/messages"
	support "forum/FORUM/support"
	"net/http"
	"regexp"

	"github.com/robfig/cron/v3"
)

func testPage(w http.ResponseWriter, r *http.Request) {
	db := database.InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	name := r.FormValue("message")

	isNotEmptyOrBlank := regexp.MustCompile(`\S`)
	if name != "" && isNotEmptyOrBlank.MatchString(name) {
		database.InsertIntoTopic(db, "Titre", name, 36, 4)
	}

	http.ServeFile(w, r, "static/testpage/testpage.html")
}

func testPage2(w http.ResponseWriter, r *http.Request) {
	db := database.InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")

	DbData, err := database.SelectArchiveFromTopic(db, 0)

	if err != nil {
		fmt.Println("Error loading DB : ", err)
	}
	data, _ := json.Marshal(DbData)
	w.Write(data)
	return
}

func createApiRep(w http.ResponseWriter, r *http.Request) {
	db := database.InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")

	DbData, _ := database.SelectReponseFromTopic(db, 1)

	fmt.Println("Réponses : ", DbData)

	data, _ := json.Marshal(DbData)
	w.Write(data)
	return
}

func main() {
	database.DataBase()
	fs := http.FileServer(http.Dir("static/"))

	c := cron.New()
	c.AddFunc("18 * * * *", func() { fmt.Println("Poogie is love, Poogie is life") })
	c.Start()

	http.HandleFunc("/", mainpage.Mainpage)
	http.HandleFunc("/messages", chat.Chat)
	http.HandleFunc("/test", testPage)
	http.HandleFunc("/test2", testPage2)
	http.HandleFunc("/login", account.ConnectionPage)
	http.HandleFunc("/signin", account.SignInPage)
	http.HandleFunc("/support", support.SupportPage)

	http.HandleFunc("/create-topic", chat.TopicPage)

	http.HandleFunc("/api-reponses", createApiRep)

	http.HandleFunc("/redirect-login", account.GetDataLogin())
	http.HandleFunc("/redirect-createaccount", account.GetData())
	//------------------------------------------------------------------
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
