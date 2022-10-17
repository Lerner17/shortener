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

	"github.com/Lerner17/shortener/internal/config"
	"github.com/Lerner17/shortener/internal/routes"
)

func main() {
	r := routes.NewRouter()
	cfg := config.Instance()
	if err := http.ListenAndServe(cfg.ServerAddress, r); err != nil {
		log.Fatal(err)
	}
}
