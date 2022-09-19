package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Lerner17/shortener/internal/config"
	er "github.com/Lerner17/shortener/internal/errors"
	"github.com/Lerner17/shortener/internal/logger"
)

type ShortenBody struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Result string `json:"result"`
}

type URLCreator interface {
	CreateURL(string, string) (string, string, error)
}

func ShortenerAPIHandler(db URLCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cfg := config.GetConfig()
		var body ShortenBody
		ctx := r.Context()
		session, ok := ctx.Value("ctxSession").(string)
		logger.Info("session from context " + session)
		if !ok {
			http.Error(w, http.StatusText(422), 422)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Cannot parse JSON"}`))
			return
		}

		if body.URL == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "URL param is empty"}`))
			return
		}

		key, _, err := db.CreateURL(session, body.URL)

		if err != nil {
			if errors.Is(err, er.ErrorShortLinkAlreadyExists) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusConflict)
				response := &ShortenResponse{
					Result: fmt.Sprintf("%s/%s", cfg.BaseURL, key),
				}
				serializedResponse, err := json.Marshal(&response)
				if err != nil {
					http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
				}
				w.Write([]byte(serializedResponse))
				return
			}
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}

		response := &ShortenResponse{
			Result: fmt.Sprintf("%s/%s", cfg.BaseURL, key),
		}
		serializedResponse, err := json.Marshal(&response)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Cannot serialize struct to JSON"}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(serializedResponse))
	}
}
