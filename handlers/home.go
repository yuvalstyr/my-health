package handlers

import (
	"net/http"
	"personal/health-app/views/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, home.HomeIndex())
}
