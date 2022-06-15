package main

import (
	// connection "forum/FORUM/ACCOUNT/connection"
	account "forum/FORUM/ACCOUNT"
	database "forum/FORUM/DATABASE"
	mainpage "forum/FORUM/mainpage"
	"net/http"
	"regexp"
)

func ConnectionPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/ACCOUNT/form/indexform.html")
}

func CreateAccountPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/ACCOUNT/sign-in/indexsign.html")
}

func testPage(w http.ResponseWriter, r *http.Request) {
	db := database.InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	name := r.FormValue("message")

	isNotEmptyOrBlank := regexp.MustCompile(`\S`)
	if name != "" && isNotEmptyOrBlank.MatchString(name) {
		database.InsertIntoTopic(db, "Titre", name, 36, 4, 0)
	}
	http.ServeFile(w, r, "static/testpage/testpage.html")
}

func main() {
	database.DataBase()
	fs := http.FileServer(http.Dir("static/"))

	http.HandleFunc("/redirect-login", account.GetDataLogin())
	http.HandleFunc("/redirect-createaccount", account.GetData())

	http.HandleFunc("/create-account", CreateAccountPage)
	http.HandleFunc("/connection", ConnectionPage)
	http.HandleFunc("/", mainpage.Mainpage())
	// http.HandleFunc("/test", testPage)

	//------------------------------------------------------------------
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// NE PAS OU BLIER DE CHANGER AVEC LE "nil" AVEC "context.ClearHandler(http.DefaultServeMux)" si jamais il y a des soucis
	http.ListenAndServe(":8080", nil)
}
