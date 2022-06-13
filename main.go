package main

import (
	// connection "forum/FORUM/ACCOUNT/connection"

	database "forum/FORUM/DATABASE"
	mainpage "forum/FORUM/mainpage"
	"net/http"
)

// func ConnectionPage(w http.ResponseWriter, r http.Request) {
// 	http.ServeFile(w, r, "static/ACCOUNT/form/indexform.html")
// }

// func CreateAccountPage(w http.ResponseWriter, rhttp.Request) {
// 	http.ServeFile(w, r, "static/ACCOUNT/sign-in/indexsign.html")
// }

// func testPage(w http.ResponseWriter, r *http.Request) {
// 	db := database.InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
// 	name := r.FormValue("message")

// 	isNotEmptyOrBlank := regexp.MustCompile(`\S`)
// 	if name != "" && isNotEmptyOrBlank.MatchString(name) {
// 		// database.InsertIntoTopic(db, "Titre", name, 36, 4)
// 	}
// 	http.ServeFile(w, r, "static/testpage/testpage.html")
// }

func main() {
	database.DataBase()
	fs := http.FileServer(http.Dir("static/"))

	http.HandleFunc("/", mainpage.Mainpage())
	// http.HandleFunc("/test", testPage)

	// http.HandleFunc("/redirect-login", account.GetData())
	// http.HandleFunc("/redirect-createaccount", account.GetData())

	// http.HandleFunc("/create-account", CreateAccountPage)
	// http.HandleFunc("/connection", ConnectionPage)

	// http.HandleFunc("/connection", connection.ConnectionPage())
	// http.HandleFunc("/login", login())
	// http.HandleFunc("/create-account", createAccount())
	//------------------------------------------------------------------
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
