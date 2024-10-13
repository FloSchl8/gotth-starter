package handlers

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

func BuildHandler(h HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			http.Error(w, err.Error(), 500)
			slog.Error(err.Error(), "path", r.URL.Path)
		}
	}
}

func render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}
