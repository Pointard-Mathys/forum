package forum

import (
	"fmt"
	"net/http"
)

type User struct {
	Name              string
	Email             string
	Password          string
	ConfirmedPassword string
}

func GetData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var structure_uti User
		fmt.Println("ICI")
		structure_uti.Name = r.FormValue("CreateAccountName")
		structure_uti.Email = r.FormValue("CreateAccountEmail")
		structure_uti.Password = r.FormValue("CreateAccountPassword")
		structure_uti.ConfirmedPassword = r.FormValue("CreateAccountPasswordConfirmed")
		if structure_uti.Password != structure_uti.ConfirmedPassword {
			http.Redirect(w, r, "/create-account?wrong2pass=true", 301)
		}
		fmt.Println(structure_uti.Name)
		fmt.Println(structure_uti.Email)
		fmt.Println(structure_uti.Password)
		fmt.Println(structure_uti.ConfirmedPassword)
	}
}
