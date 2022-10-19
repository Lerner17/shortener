package routes

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Lerner17/shortener/internal/db"
	mw "github.com/Lerner17/shortener/internal/middleware"
)

func NewRouter() chi.Router {

	var db = db.GetDB()

	r := chi.NewRouter()
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Logger)
	r.Use(mw.SessionMiddleware)
	r.Use(mw.GZIPMiddleware)
	r.Post("/api/shorten", ShortenerAPIHandler(db))
	r.Post("/api/shorten/batch", BatchAPIHandler(db))
	r.Get("/api/user/urls", UserURLsAPIHandler(db))
	r.Delete("/api/user/urls", DeleteUserURLsAPIHandler(db))
	r.Get("/{urlID}", RedirectHandler(db))
	r.Post("/", CreateShortURLHandler(db))
	r.Get("/ping", PingHandler())
	return r
}
