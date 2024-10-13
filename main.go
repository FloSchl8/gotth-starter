package main

import (
	"log"
	"net/http"
	"os"

	"gotth-starter/package/handlers"
	l "gotth-starter/package/log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	slogchi "github.com/samber/slog-chi"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	logger := l.StructuredLogger()

	r := chi.NewMux()
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(slogchi.New(logger))

	r.Handle("/*", public())
	r.Get("/", handlers.BuildHandler(handlers.HandleHome))
	r.Get("/hello", handlers.BuildHandler(handlers.HandleHello))

	listenAddr := os.Getenv("LISTEN_ADDR")

	logger.Info("Server started", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, r)
}
