package routes

import (
	"net/http"

	"github.com/Lerner17/shortener/internal/db/psql"
)

func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn := psql.GetInstance()
		if err := conn.Ping(); err == nil {
			w.WriteHeader(http.StatusOK)

		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
