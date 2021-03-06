package handlers

import (
	"net/http"
	"strings"

	database "github.com/Lerner17/shortener/internal/db"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetInstance()
	url := strings.Split(r.URL.Path, "/")
	if fullURL, ok := db.Find(url[1]); len(url) > 1 && url[1] != "" && ok {
		http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Undefined redirect"))
	}

}
