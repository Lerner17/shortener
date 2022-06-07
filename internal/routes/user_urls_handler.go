package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Lerner17/shortener/internal/helpers"
	"github.com/Lerner17/shortener/internal/models"
)

type URLListGetter interface {
	GetUserURLs(string) models.URLs
}

func UserURLsAPIHandler(db URLListGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token, ok := ctx.Value("token").(string)
		if !ok {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		uuid := helpers.GetUUIDFromToken(token)

		urlsList := db.GetUserURLs(uuid)
		fmt.Printf("arrays length: %v\n", len(urlsList))
		if len(urlsList) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		serializedResponse, err := json.Marshal(&urlsList)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Cannot serialize struct to JSON"}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(serializedResponse))
	}
}
