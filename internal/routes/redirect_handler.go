package routes

import (
	"net/http"

	database "github.com/Lerner17/shortener/internal/db"
	"github.com/go-chi/chi/v5"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetInstance()

	urlID := chi.URLParam(r, "urlID")
	if fullURL, ok := db.Find(urlID); ok {
		w.Header().Set("Content-Type", "plain/text")
		http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
	} else {
		w.Header().Set("Content-Type", "plain/text")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Not found"))
		return
	}

}
