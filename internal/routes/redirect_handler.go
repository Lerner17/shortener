package routes

import (
	"fmt"
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

		ctx := r.Context()
		token, ok := ctx.Value("token").(string)
		if !ok {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		uuid := helpers.GetUUIDFromToken(token)
		fmt.Println(db.GetURL(uuid, urlID))
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
