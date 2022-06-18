package forum

import (
	"net/http"
)

func SignInPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "././static/ACCOUNT/signin/signin.html")
}
