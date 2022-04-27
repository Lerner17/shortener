package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateShortURLHandler(t *testing.T) {

	type want struct {
		contentType string
		statusCode  int
	}

	tests := []struct {
		name    string
		request string
		payload map[string]interface{}
		want    want
	}{
		{
			name:    "Create url success test #1",
			request: "/createShortURL",
			payload: map[string]interface{}{"url": "http://example.com"},
			want: want{
				contentType: "plain/text",
				statusCode:  http.StatusCreated,
			},
		},
		{
			name:    "Create url success test #2",
			request: "/createShortURL",
			payload: map[string]interface{}{"url": "http://yandex.ru"},
			want: want{
				contentType: "plain/text",
				statusCode:  http.StatusCreated,
			},
		},
		{
			name:    "Create url bad test #1",
			request: "/createShortURL",
			payload: map[string]interface{}{},
			want: want{
				contentType: "plain/text",
				statusCode:  http.StatusBadRequest,
			},
		},
		{
			name:    "Create url bad test #2",
			request: "/createShortURL",
			payload: map[string]interface{}{"u": "http://example.com"},
			want: want{
				contentType: "plain/text",
				statusCode:  http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requestURL string = tt.request
			fmt.Println(requestURL)
			body, err := json.Marshal(tt.payload)
			require.NoError(t, err)
			request := httptest.NewRequest(http.MethodPost, requestURL, bytes.NewReader(body))
			request.Header.Add("Content-Type", "application/json")
			w := httptest.NewRecorder()
			h := http.HandlerFunc(CreateShortURLHandler)
			h.ServeHTTP(w, request)
			result := w.Result()
			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
		})
	}
}
