package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	database "github.com/Lerner17/shortener/internal/db"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetInstance()

	switch r.Method {
	case "POST":
		db := database.GetInstance()
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil || string(body) == "" {
			w.Header().Set("Content-Type", "plain/text")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad request"))
			return
		}
		key, _ := db.Insert(string(body))
		w.Header().Set("Content-Type", "plain/text")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("http://localhost:8080/%s", key)))
	case "GET":
		url := strings.Split(r.URL.Path, "/")
		if fullURL, ok := db.Find(url[1]); len(url) > 1 && url[1] != "" && ok {
			w.Header().Set("Content-Type", "plain/text")
			http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
		} else {
			w.Header().Set("Content-Type", "plain/text")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Undefined redirect"))
			return
		}
	default:
		w.Header().Set("Content-Type", "plain/text")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}

}
