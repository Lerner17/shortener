package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Lerner17/shortener/internal/models"
)

type URLListGetter interface {
	GetUserURLs(string) models.URLs
}

func UserURLsAPIHandler(db URLListGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		session, ok := ctx.Value("ctxSession").(string)
		if !ok {
			http.Error(w, http.StatusText(422), 422)
			return
		}
		urlsList := db.GetUserURLs(session)
		serializedResponse, err := json.Marshal(&urlsList)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Cannot serialize struct to JSON"}`))
			return
		}

		if len(urlsList) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(serializedResponse))

	}
}
