package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lerner17/shortener/internal/db"
	"github.com/Lerner17/shortener/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func TestRedirectHandler(t *testing.T) {
	r := NewRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	database := db.GetInstance()

	resp, _ := helpers.TestRequest(t, ts, http.MethodGet, fmt.Sprintf("/%s", "abc331"), nil)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp.Body.Close()

	resp, _ = helpers.TestRequest(t, ts, http.MethodGet, fmt.Sprintf("/%s", "asdsadad"), nil)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp.Body.Close()

	resp, _ = helpers.TestRequest(t, ts, http.MethodGet, fmt.Sprintf("/%s", "asdsadad"), nil)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp.Body.Close()

	id, _ := database.Insert("http://yandex.ru")
	resp, _ = helpers.TestRequest(t, ts, http.MethodGet, fmt.Sprintf("/%s", id), nil)
	assert.Equal(t, http.StatusTemporaryRedirect, resp.StatusCode)
	resp.Body.Close()

}
