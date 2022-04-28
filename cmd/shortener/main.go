// Напишите сервис для сокращения длинных URL. Требования:
// Сервер должен быть доступен по адресу: http://localhost:8080.
// Сервер должен предоставлять два эндпоинта: POST / и GET /{id}.
// Эндпоинт POST / принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле.
// Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
// Нужно учесть некорректные запросы и возвращать для них ответ с кодом 400.

package main

import (
	"log"
	"net/http"

	"github.com/Lerner17/shortener/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.RedirectHandler)
	server := &http.Server{
		Addr: "127.0.0.1:8080",
	}
	log.Fatal(server.ListenAndServe())
}
