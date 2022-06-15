package forum

import (
	forum "forum/FORUM/DATABASE"
	"net/http"
)

type User struct {
	Name              string
	Email             string
	Password          string
	ConfirmedPassword string
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

		} else if structure_uti.Password == structure_uti.ConfirmedPassword {
			forum.InsertIntoUsers(db, structure_uti.Name, structure_uti.Email, forum.Encoding_password(structure_uti.Password))
		}
	}
}
