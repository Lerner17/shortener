package helpers

import (
	"net/http"

	"github.com/Lerner17/shortener/internal/logger"
)

func GetSessionFromContext(w http.ResponseWriter, r *http.Request) string {
	ctx := r.Context()
	session, ok := ctx.Value("ctxSession").(string)
	logger.Info("session from context " + session)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return ""
	}
	return session
}
