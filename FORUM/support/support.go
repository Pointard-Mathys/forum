package forum

import (
	"net/http"
)

func SupportPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "././static/support/support.html")
}
