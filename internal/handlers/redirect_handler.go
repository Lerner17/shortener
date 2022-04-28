package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	database "github.com/Lerner17/shortener/internal/db"
)

type createShortURLBody struct {
	URL string `json:"url"`
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetInstance()
	if r.Method == "GET" {
		url := strings.Split(r.URL.Path, "/")
		if fullURL, ok := db.Find(url[1]); len(url) > 1 && url[1] != "" && ok {
			w.Header().Set("Content-Type", "plain/text")
			http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
		} else {
			w.Header().Set("Content-Type", "plain/text")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Undefined redirect"))
		}
	} else if r.Method == "POST" {
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

}
