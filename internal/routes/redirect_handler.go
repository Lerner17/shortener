package routes

import (
	"net/http"

	"github.com/Lerner17/shortener/internal/helpers"
	"github.com/go-chi/chi/v5"
)

type URLGetter interface {
	GetURL(string, string) (string, bool)
}

func RedirectHandler(db URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlID := chi.URLParam(r, "urlID")

		var cookie *http.Cookie
		cookie, err := r.Cookie("token")
		if err != nil || !helpers.ValidateToken(cookie) {
			cookie = helpers.CreateToken()
		}
		uuid := helpers.GetUUIDFromToken(cookie.Value)

		if fullURL, ok := db.GetURL(uuid, urlID); ok {
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
