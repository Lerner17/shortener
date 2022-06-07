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

		var cookie *http.Cookie
		cookie, err = r.Cookie("token")
		if err != nil || !helpers.ValidateToken(cookie) {
			cookie = helpers.CreateToken()
		}
		uuid := helpers.GetUUIDFromToken(cookie.Value)

		key, _ := db.CreateURL(uuid, string(body))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("%s/%s", cfg.BaseURL, key)))
	}
}
