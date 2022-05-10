package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Lerner17/shortener/internal/config"
	database "github.com/Lerner17/shortener/internal/db"
)

func CreateShortURLHandler(w http.ResponseWriter, r *http.Request) {

	db := database.NewURLStorage()
	cfg := config.GetConfig()
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil || string(body) == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	key, _ := db.CreateURL(string(body))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("%s/%s", cfg.BaseURL, key)))
}
