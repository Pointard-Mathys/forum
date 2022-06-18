package forum

import (
	"net/http"
)

func ConnectionPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "././static/ACCOUNT/connection/connection.html")
}
