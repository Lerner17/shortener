package middleware

import (
	"context"
	"net/http"

	"github.com/Lerner17/shortener/internal/helpers"
)

type ContextType string

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// var token string
		var cookie *http.Cookie

		cookie, err := r.Cookie("token")
		if err != nil || !helpers.ValidateToken(cookie) {
			cookie = helpers.CreateToken()
		}
		http.SetCookie(w, cookie)

		ctx := context.WithValue(r.Context(), "token", cookie.Value)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
