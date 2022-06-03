package main

// connection "forum/FORUM/ACCOUNT/connection"
// mainpage "forum/FORUM/mainpage"
import (
	database "forum/FORUM/DATABASE"
)

func main() {
	database.DataBase()
	// fs := http.FileServer(http.Dir("FORUM/CSS/"))

	// http.HandleFunc("/", mainpage.Mainpage())
	// http.HandleFunc("/connection", connection.ConnectionPage())
	// // http.HandleFunc("/login", login())
	// // http.HandleFunc("/create-account", createAccount())
	// //------------------------------------------------------------------
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	// http.ListenAndServe(":8080", nil)
}
