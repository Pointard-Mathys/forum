package forum

import (
	"fmt"
	session "forum/FORUM/ACCOUNT"
	"html/template"
	"net/http"
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
	return func(w http.ResponseWriter, r *http.Request) {
		IdToDelet := r.FormValue("IdToDelet")
		IDtype := r.FormValue("IDtype")
		fmt.Println(IdToDelet)
		fmt.Println(IDtype)
		http.Redirect(w, r, "/admin", http.StatusFound)
	}
}

func AddTheme() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Theme := r.FormValue("Theme")
		fmt.Println("ici")
		fmt.Println(Theme)
		http.Redirect(w, r, "/admin", http.StatusFound)
	}
}
