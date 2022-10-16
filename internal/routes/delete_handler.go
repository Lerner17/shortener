package routes

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Lerner17/shortener/internal/helpers"
	"github.com/Lerner17/shortener/internal/logger"
	"go.uber.org/zap"
)

type URLDeleter interface {
	DeleteBatchURL(context.Context, []string, string) error
}

type URLsList []string

func DeleteUserURLsAPIHandler(db URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := helpers.GetSessionFromContext(w, r)
		body, err := io.ReadAll(r.Body)
		if err != nil || string(body) == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad request"))
			return
		}

		var shortURLs URLsList
		if err = json.Unmarshal(body, &shortURLs); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad request"))
			return
		}
		ctx := context.Background()
		go func() {
			if err = db.DeleteBatchURL(ctx, shortURLs, session); err != nil {
				logger.Error("cannot delete batch", zap.Error(err))
			}
		}()

		w.WriteHeader(http.StatusAccepted)
	}
}
