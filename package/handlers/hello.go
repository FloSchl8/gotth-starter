package handlers

import (
	"net/http"

	"gotth-starter/package/views/hello"
)

func HandleHello(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, hello.Index())
}
