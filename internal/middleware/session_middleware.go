package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

var key = []byte("asd")

func parseCookie(cookie *http.Cookie) ([]string, error) {
	values := strings.Split(cookie.Value, ":")
	if len(values) != 2 {
		return []string{"", ""}, errors.New("Invalid cookie")
	}
	return values, nil
}

func createNewCookie() *http.Cookie {
	session := uuid.NewString()
	h := hmac.New(sha256.New, key)
	h.Write([]byte(session))
	dst := h.Sum(nil)

	return &http.Cookie{
		Name:  "session",
		Value: fmt.Sprintf("%s:%x", session, dst),
		Path:  "/",
	}
}

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token, err := r.Cookie("session")

		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Println("No cookie")
				token = createNewCookie()
			} else {
				panic(err)
			}
		}

		fmt.Println(token)
		http.SetCookie(w, token)
		next.ServeHTTP(w, r)
	})
}
