// Напишите сервис для сокращения длинных URL. Требования:
// Сервер должен быть доступен по адресу: http://localhost:8080.
// Сервер должен предоставлять два эндпоинта: POST / и GET /{id}.
// Эндпоинт POST / принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле.
// Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
// Нужно учесть некорректные запросы и возвращать для них ответ с кодом 400.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Lerner17/shortener/internal/config"
	"github.com/Lerner17/shortener/internal/routes"
)

func parsArgs(c *config.Config) {
	serverAddressPtr := flag.String("a", "", "")
	baseURLPtr := flag.String("b", "", "")
	fileStoragePathPtr := flag.String("f", "", "")
	DatabaseDsnPtr := flag.String("d", "", "")
	flag.Parse()

	if *serverAddressPtr != "" {
		c.ServerAddress = *serverAddressPtr
	}

	if *baseURLPtr != "" {
		c.BaseURL = *baseURLPtr
	}

	if *fileStoragePathPtr != "" {
		c.FileStoragePath = *fileStoragePathPtr
	}

	if *DatabaseDsnPtr != "" {
		c.DatabaseDsn = *DatabaseDsnPtr
	}
}

func main() {
	cfg := config.GetConfig()
	parsArgs(cfg)
	r := routes.NewRouter()
	fmt.Println(cfg)
	if err := http.ListenAndServe(cfg.ServerAddress, r); err != nil {
		log.Fatal(err)
	}
}
