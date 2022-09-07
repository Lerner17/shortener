package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Lerner17/shortener/internal/config"
)

type CreateShortURLHandlerURLCreator interface {
	CreateURL(string, string) (string, string)
}

func CreateShortURLHandler(db CreateShortURLHandlerURLCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cfg := config.GetConfig()
		ctx := r.Context()
		session, ok := ctx.Value("ctxSession").(string)
		if !ok {
			http.Error(w, http.StatusText(422), 422)
			return
		}
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil || string(body) == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad request"))
			return
		}
		key, _ := db.CreateURL(session, string(body))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("%s/%s", cfg.BaseURL, key)))
	}
}
