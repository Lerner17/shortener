package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type URLGetter interface {
	GetURL(string, string) (string, bool)
}

func RedirectHandler(db URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlID := chi.URLParam(r, "urlID")
		ctx := r.Context()
		session, ok := ctx.Value("ctxSession").(string)
		if !ok {
			http.Error(w, http.StatusText(422), 422)
			return
		}
		if fullURL, ok := db.GetURL(session, urlID); ok {
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
