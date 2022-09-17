package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Lerner17/shortener/internal/helpers"
	"github.com/Lerner17/shortener/internal/logger"
	"github.com/Lerner17/shortener/internal/models"
	"go.uber.org/zap"
)

type URLBatchCreator interface {
	CreateBatchURL(string, models.BatchURLs) (models.BatchShortURLs, error)
}

func BatchAPIHandler(db URLBatchCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var urls models.BatchURLs
		if err := json.NewDecoder(r.Body).Decode(&urls); err != nil {
			panic(err)
		}
		logger.Info("[POST /api/shorten/batch] Body:", zap.Reflect("body", urls))
		session := helpers.GetSessionFromContext(w, r)

		results, err := db.CreateBatchURL(session, urls)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}

		serializedResponse, err := json.Marshal(&results)
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
}
