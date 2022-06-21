package forum

import (
	forum "forum/FORUM/DATABASE"
	"html/template"
	"net/http"
)

type User struct {
	Name              string
	Email             string
	Password          string
	ConfirmedPassword string
}

func SignInPage(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("././static/ACCOUNT/signin/signin.html", "././static/home/header.html"))
	if r.FormValue("Wrongnoob") == "true" {
		templates.Execute(w, struct{ Wrongnoob bool }{true})
		return
	}
	templates.Execute(w, nil)
}

func GetData() http.HandlerFunc {
	db := forum.InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	return func(w http.ResponseWriter, r *http.Request) {
		var structure_uti User
		structure_uti.Name = r.FormValue("CreateAccountName")
		structure_uti.Email = r.FormValue("CreateAccountEmail")
		structure_uti.Password = r.FormValue("CreateAccountPassword")
		structure_uti.ConfirmedPassword = r.FormValue("CreateAccountPasswordConfirmed")

		if structure_uti.Password != structure_uti.ConfirmedPassword {
			http.Redirect(w, r, "/signin?Wrongnoob=true", 301)
		} else if structure_uti.Password == structure_uti.ConfirmedPassword {
			forum.InsertIntoUsers(db, structure_uti.Name, structure_uti.Email, forum.Encoding_password(structure_uti.Password))
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
}
