package routes

import (
	"github.com/Lerner17/shortener/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() chi.Router {

	var db = db.GetDB()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/api/shorten", ShortenerAPIHandler(db))
	r.Get("/{urlID}", RedirectHandler(db))
	r.Post("/", CreateShortURLHandler(db))
	return r
}
