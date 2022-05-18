package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Lerner17/shortener/internal/config"
)

type CreateShortURLHandlerUrlCreator interface {
	CreateURL(string) (string, string)
}

func CreateShortURLHandler(db CreateShortURLHandlerUrlCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cfg := config.GetConfig()
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil || string(body) == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad request"))
			return
		}
		key, _ := db.CreateURL(string(body))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("%s/%s", cfg.BaseURL, key)))
	}
}
