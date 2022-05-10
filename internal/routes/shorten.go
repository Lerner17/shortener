package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "github.com/Lerner17/shortener/internal/db"
)

type ShortenBody struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Result string `json:"result"`
}

func ShortenerAPIHandler(w http.ResponseWriter, r *http.Request) {
	db := database.NewURLStorage()
	var body ShortenBody
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

	key, _ := db.CreateURL(body.URL)
	response := &ShortenResponse{
		Result: fmt.Sprintf("http://localhost:8080/%s", key),
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
