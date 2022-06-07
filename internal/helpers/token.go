package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Lerner17/shortener/internal/config"
	"github.com/google/uuid"
)

var cfg *config.Config = config.GetConfig()

var key = []byte(cfg.SecretKey)

func CreateToken() *http.Cookie {
	token := uuid.NewString()

	log.Println("Create new token ", token)

	h := hmac.New(sha256.New, key)
	h.Write([]byte(token))
	dst := h.Sum(nil)

	cookie := &http.Cookie{
		Name:  "token",
		Value: fmt.Sprintf("%s:%x", token, dst),
		Path:  "/",
	}
	return cookie
}

func ValidateToken(cookie *http.Cookie) bool {
	values := strings.Split(cookie.Value, ":")
	data, err := hex.DecodeString(values[1])
	if err != nil {
		log.Fatal(err)
		return false
	}

	uuid := values[0]
	h := hmac.New(sha256.New, key)
	h.Write([]byte(uuid))
	sign := h.Sum(nil)
	return hmac.Equal(sign, data)
}

func GetUUIDFromToken(token string) string {
	values := strings.Split(token, ":")
	return values[0]
}
