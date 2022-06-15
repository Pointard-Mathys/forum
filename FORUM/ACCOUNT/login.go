package forum

import (
	"fmt"
	forum "forum/FORUM/DATABASE"
	"net/http"
	"reflect"
)

type UserLogin struct {
	Email    string
	Password string
}

func GetDataLogin() http.HandlerFunc {
	db := forum.InitDatabase("FORUM/DATABASE/databaseHolder/DATA_BASE.db")
	return func(w http.ResponseWriter, r *http.Request) {
		var structure_utiLogin UserLogin
		structure_utiLogin.Email = r.FormValue("loginEmail")
		structure_utiLogin.Password = r.FormValue("CreateAccountPassword")

		user := forum.Login(db, structure_utiLogin.Email, forum.Encoding_password(structure_utiLogin.Password))
		fmt.Println(reflect.TypeOf(user))
	}
}
