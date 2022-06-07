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
		var cookie *http.Cookie
		cookie, err := r.Cookie("token")
		if err != nil || !helpers.ValidateToken(cookie) {
			cookie = helpers.CreateToken()
		}
		uuid := helpers.GetUUIDFromToken(cookie.Value)

		urlsList := db.GetUserURLs(uuid)
		fmt.Printf("arrays length: %v\n", len(urlsList))
		fmt.Printf("arrays: %v\n", urlsList)
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
