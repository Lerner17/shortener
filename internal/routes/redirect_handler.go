package routes

import (
	"net/http"

	"github.com/Lerner17/shortener/internal/logger"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type URLGetter interface {
	GetURL(string) (string, bool, bool)
}

func RedirectHandler(db URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlID := chi.URLParam(r, "urlID")
		logger.Info("URL ID:", zap.String("urlID", urlID))
		ctx := r.Context()
		session, ok := ctx.Value("ctxSession").(string)
		logger.Info("Session:", zap.String("session", session))
		if !ok {
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}
		fullURL, isDeleted, ok := db.GetURL(urlID)
		logger.Info("Get from DB status", zap.Bool("ok", ok))
		logger.Info("Value from database", zap.String("value", fullURL))
		if ok {
			if isDeleted == true {
				w.Header().Set("Content-Type", "plain/text")
				http.Redirect(w, r, fullURL, http.StatusGone)
				return
			}
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
