package forum

import (
	"fmt"
	forum "forum/FORUM/DATABASE"
	"html/template"
	"net/http"
)

type UserLogin struct {
	Email    string
	Password string
}

func ConnectionPage(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("././static/ACCOUNT/connection/connection.html", "././static/home/header.html"))
	templates.Execute(w, nil)
}

func GetDataLogin() http.HandlerFunc {
	db := forum.InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	return func(w http.ResponseWriter, r *http.Request) {
		var structure_utiLogin UserLogin
		structure_utiLogin.Email = r.FormValue("loginEmail")
		structure_utiLogin.Password = r.FormValue("loginPassword")

		user := forum.Login(db, structure_utiLogin.Email, structure_utiLogin.Password)
		//VERIFICATION D'OBJET VIDE
		if user.Id >= 1 {
			session, _ := Store.Get(r, "session")
			// CECI DOIT ETRE EGALE AU USER ID DE LA BDD C'EST LOGIQUE CONNARD
			session.Values["userName"] = user.Name
			session.Values["userID"] = user.Id
			// SAUVEGARDE DE LA SESSION
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			fmt.Println("CACA DANS LE CULCUL")
		}
	}
}
