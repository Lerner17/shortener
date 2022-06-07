package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Lerner17/shortener/internal/config"
	"github.com/Lerner17/shortener/internal/helpers"
)

type CreateShortURLHandlerURLCreator interface {
	CreateURL(string, string) (string, string)
}

func CreateShortURLHandler(db CreateShortURLHandlerURLCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cfg := config.GetConfig()
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil || string(body) == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad request"))
			return
		}
		ctx := r.Context()
		token, ok := ctx.Value("token").(string)
		if !ok {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		uuid := helpers.GetUUIDFromToken(token)
		key, _ := db.CreateURL(uuid, string(body))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("%s/%s", cfg.BaseURL, key)))
	}
}
