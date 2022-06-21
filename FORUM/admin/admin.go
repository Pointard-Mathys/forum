package forum

import (
	"fmt"
	session "forum/FORUM/ACCOUNT"
	forum "forum/FORUM/DATABASE"
	"html/template"
	"net/http"
	"strconv"
)

func AdminHolder(w http.ResponseWriter, r *http.Request) {
	session := session.Store
	sessions, _ := session.Get(r, "session")
	var templates = template.Must(template.ParseFiles("././static/admin/adminform.html", "././static/home/header.html"))
	if sessions.Values["userID"] != 1 {
		http.Redirect(w, r, "/", http.StatusFound)
		fmt.Println("Accès refusé")
	}
	templates.Execute(w, nil)
}

func DeletID() http.HandlerFunc {
	db := forum.InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	return func(w http.ResponseWriter, r *http.Request) {
		IdToDelet := r.FormValue("IdToDelet")
		NewIdToDelet, _ := strconv.Atoi(IdToDelet)
		IDtype := r.FormValue("IDtype")
		fmt.Println(NewIdToDelet)
		fmt.Println(IDtype)
		if IDtype == "utilisateur" {
			fmt.Println("Bannissement de l'utilisateur")
			forum.UpDate(db, "users", "password", "", NewIdToDelet)
		}
		if IDtype == "topic" {
			fmt.Println("Suppression du Topic")
			forum.ArchiveSpecificTopic(db, NewIdToDelet)
		}
		if IDtype == "reponse" {
			fmt.Println("Suppression de la reponse")
			forum.UpDate(db, "reponses", "content", "This Response has been deleted", NewIdToDelet)
		}
		http.Redirect(w, r, "/admin", http.StatusFound)
	}
}

func AddTheme() http.HandlerFunc {
	db := forum.InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	return func(w http.ResponseWriter, r *http.Request) {
		Theme := r.FormValue("Theme")
		fmt.Println("Thème ajouté : ", Theme)
		forum.InsertIntoTheme(db, Theme)
		http.Redirect(w, r, "/admin", http.StatusFound)
	}
}
