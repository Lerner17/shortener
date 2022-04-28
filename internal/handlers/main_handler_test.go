package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lerner17/shortener/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestMainHandlerSuccess(t *testing.T) {
	database := db.GetInstance()
	var urlKeyValue map[string]string = make(map[string]string)
	urlList := []string{"https://youtube.com", "https://yandex.ru", "https://google.com", "https://go.dev"}

	type want struct {
		contentType string
		statusCode  int
		content     string
	}

	for index := range urlList {
		dbID, _ := database.Insert(urlList[index])
		urlKeyValue[urlList[index]] = dbID
	}

	tests := []struct {
		name    string
		request string
		method  string
		body    string
		want    want
	}{
		{
			name:    "Redirect Success test #1",
			request: "/",
			method:  http.MethodGet,
			body:    "",
			want:    want{statusCode: 307, contentType: "plain/text", content: "https://youtube.com"},
		},
		{
			name:    "Redirect Success test #2",
			request: "/",
			method:  http.MethodGet,
			body:    "",
			want:    want{statusCode: 307, contentType: "plain/text", content: "https://yandex.ru"},
		},
		{
			name:    "Redirect Success test #3",
			request: "/",
			method:  http.MethodGet,
			body:    "",
			want:    want{statusCode: 307, contentType: "plain/text", content: "https://google.com"},
		},
		{
			name:    "Redirect Success test #4",
			request: "/",
			method:  http.MethodGet,
			body:    "",
			want:    want{statusCode: 307, contentType: "plain/text", content: "https://go.dev"},
		},
		{
			name:    "Bad test #1",
			request: "/aaa",
			method:  http.MethodGet,
			body:    "",
			want: want{
				contentType: "plain/text",
				statusCode:  http.StatusBadRequest,
			},
		},
		{
			name:    "Bad test #3",
			request: "/ccc",
			method:  http.MethodGet,
			body:    "",
			want: want{
				contentType: "plain/text",
				statusCode:  http.StatusBadRequest,
			},
		},
		{
			name:    "Bad test #3",
			request: "/xyz1",
			method:  http.MethodGet,
			body:    "",
			want: want{
				contentType: "plain/text",
				statusCode:  http.StatusBadRequest,
			},
		},
		{
			name:    "Create short url #1",
			request: "/",
			method:  http.MethodPost,
			body:    "http://vk.com",
			want: want{
				contentType: "plain/text",
				statusCode:  http.StatusCreated,
			},
		},
		{
			name:    "Create short url #2",
			request: "/",
			method:  http.MethodPost,
			body:    "http://meduza.io",
			want: want{
				contentType: "plain/text",
				statusCode:  http.StatusCreated,
			},
		},
		{
			name:    "Create short url #3",
			request: "/",
			method:  http.MethodPost,
			body:    "",
			want: want{
				contentType: "plain/text",
				statusCode:  http.StatusBadRequest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestURL := fmt.Sprintf("%s%s", tt.request, urlKeyValue[tt.want.content])

			request := httptest.NewRequest(tt.method, requestURL, bytes.NewReader([]byte(tt.body)))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(MainHandler)
			h.ServeHTTP(w, request)
			result := w.Result()
			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
		})
	}
}
