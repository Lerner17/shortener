package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lerner17/shortener/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestRedirectHandlerSuccess(t *testing.T) {
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
		urlsMap map[string]string
		want    want
	}{
		{
			name:    "Redirect Success test #1",
			request: "/",
			urlsMap: urlKeyValue,
			want:    want{statusCode: 307, contentType: "plain/text", content: "https://youtube.com"},
		},
		{
			name:    "Redirect Success test #2",
			request: "/",
			urlsMap: urlKeyValue,
			want:    want{statusCode: 307, contentType: "plain/text", content: "https://yandex.ru"},
		},
		{
			name:    "Redirect Success test #3",
			request: "/",
			urlsMap: urlKeyValue,
			want:    want{statusCode: 307, contentType: "plain/text", content: "https://google.com"},
		},
		{
			name:    "Redirect Success test #4",
			request: "/",
			urlsMap: urlKeyValue,
			want:    want{statusCode: 307, contentType: "plain/text", content: "https://go.dev"},
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requestURL string = fmt.Sprintf("%s%s", tt.request, urlKeyValue[tt.want.content])
			request := httptest.NewRequest(http.MethodGet, requestURL, nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(RedirectHandler)
			h.ServeHTTP(w, request)
			result := w.Result()
			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
		})
	}
}

func TestRedirectHandlerUndefinded(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
	}

	tests := []struct {
		name    string
		request string
		id      string
		want    want
	}{
		{
			name:    "Bad test #1",
			request: "/aaa",
			want: want{
				contentType: "plain/text",
				statusCode:  http.StatusBadRequest,
			},
		},
		{
			name:    "Bad test #3",
			request: "/ccc",
			want: want{
				contentType: "plain/text",
				statusCode:  http.StatusBadRequest,
			},
		},
		{
			name:    "Bad test #3",
			request: "/xyz1",
			want: want{
				contentType: "plain/text",
				statusCode:  http.StatusBadRequest,
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tt.request, nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(RedirectHandler)
			h.ServeHTTP(w, request)
			result := w.Result()
			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
		})
	}

}
