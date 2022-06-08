package main

import (
	// connection "forum/FORUM/ACCOUNT/connection"
	database "forum/FORUM/DATABASE"
	mainpage "forum/FORUM/mainpage"
	theme "forum/FORUM/testTHME"
	"net/http"
)

func testPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/testpage/testpage.html")
}

func main() {
	database.DataBase()
	fs := http.FileServer(http.Dir("static/"))

	http.HandleFunc("/", mainpage.Mainpage())
	http.HandleFunc("/test", testPage)
	http.HandleFunc("/testpage", theme.TestTHEME())
	// http.HandleFunc("/connection", connection.ConnectionPage())
	// http.HandleFunc("/login", login())
	// http.HandleFunc("/create-account", createAccount())
	//------------------------------------------------------------------
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
