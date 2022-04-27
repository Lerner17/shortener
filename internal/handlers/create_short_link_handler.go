package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "github.com/Lerner17/shortener/internal/db"
)

type createShortUrlBody struct {
	URL string `json:"url"`
}

func CreateShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var body createShortUrlBody
	db := database.GetInstance()
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil || body.URL == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	key, _ := db.Insert(body.URL)
	w.Write([]byte(fmt.Sprintf("http://localhost:8080/%s", key)))
}
