package routes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lerner17/shortener/internal/config"
	"github.com/Lerner17/shortener/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func TestShortenerAPIHandler(t *testing.T) {

	config.GetConfig()
	r := NewRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, _ := helpers.TestRequest(t, ts, http.MethodPost, "/api/shorten", bytes.NewReader([]byte(`{"rl": "https://example.com"}`)))
	defer resp.Body.Close()
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, helpers.ContentTypeJSON, resp.Header.Get("Content-Type"))

	resp, _ = helpers.TestRequest(t, ts, http.MethodPost, "/api/shorten", bytes.NewReader([]byte(`{"url": "https://example.com"}`)))
	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, helpers.ContentTypeJSON, resp.Header.Get("Content-Type"))

	resp, _ = helpers.TestRequest(t, ts, http.MethodPost, "/api/shorten", bytes.NewReader([]byte(`{"url": ""}`)))
	defer resp.Body.Close()
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, helpers.ContentTypeJSON, resp.Header.Get("Content-Type"))

}
