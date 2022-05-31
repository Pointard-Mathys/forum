package forum

import (
	"net/http"
	"text/template"
)

func Mainpage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./FORUM/mainpage/mainpage.html")
		tmpl.Execute(w, r)
	}
}
