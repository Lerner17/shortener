package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "github.com/Lerner17/shortener/internal/db"
)

type createShortURLBody struct {
	URL string `json:"url"`
}

func CreateShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var body createShortURLBody
	db := database.GetInstance()
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil || body.URL == "" {
		w.Header().Set("Content-Type", "plain/text")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
	}
	key, _ := db.Insert(body.URL)
	w.Header().Set("Content-Type", "plain/text")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("http://localhost:8080/%s", key)))
}
