package handlers

import (
	"net/http"

	"gotth-starter/package/views"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, views.Home())
}
