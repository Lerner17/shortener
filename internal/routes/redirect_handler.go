package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UrlGetter interface {
	GetURL(string) (string, bool)
}

func RedirectHandler(db UrlGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlID := chi.URLParam(r, "urlID")
		if fullURL, ok := db.GetURL(urlID); ok {
			w.Header().Set("Content-Type", "plain/text")
			http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
		} else {
			w.Header().Set("Content-Type", "plain/text")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Not found"))
			return
		}

	}
}
