package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Lerner17/shortener/internal/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ContextType string

var key = []byte("haskjdhkjdhakjsd@12")

var UserIDCtxName ContextType = "ctxSession"

func createNewCookie(session string) *http.Cookie {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(session))
	dst := h.Sum(nil)

	return &http.Cookie{
		Name:  "session",
		Value: fmt.Sprintf("%s:%x", session, dst),
		Path:  "/",
	}
}

func validateSession(uuid string, stringSign string) bool {
	decodedSign, err := hex.DecodeString(stringSign)
	if err != nil {
		return false
	}
	h := hmac.New(sha256.New, key)
	h.Write([]byte(uuid))
	sign := h.Sum(nil)
	return hmac.Equal(sign, decodedSign)
}

func getSessionFromCookie(cookieValue string, session *string) error {
	logger.Info("Session by link", zap.String("s", *session))
	values := strings.Split(cookieValue, ":")
	if len(values) != 2 {
		return errors.New("INVALID COOKIE")
	}

	uuid, sign := values[0], values[1]
	isSessionValid := validateSession(uuid, sign)
	logger.Info("Validate signed session:", zap.Bool("result", isSessionValid))
	if isSessionValid {
		*session = uuid
	}
	return nil
}

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session := uuid.New().String()
		logger.Info("Generate session", zap.String("session:", session))
		sessionCookie, err := r.Cookie("session")

		if err == nil {
			_ = getSessionFromCookie(sessionCookie.Value, &session)
		}

		ctx := context.WithValue(r.Context(), "ctxSession", session)
		logger.Info("", zap.Reflect("a", ctx))
		http.SetCookie(w, createNewCookie(session))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
