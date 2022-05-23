package config

import (
	"flag"
	"log"
	"os"
	"sync"
)

type Config struct {
	ServerAddress   string
	BaseURL         string
	FileStoragePath string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var instance *Config

var once sync.Once

func GetConfig() *Config {

	once.Do(func() {
		log.Println("Load config...")
		addressPtr := flag.String("a", getEnv("SERVER_ADDRESS", "localhost:8080"), "SERVER_ADDRESS")
		baseURLPtr := flag.String("b", getEnv("BASE_URL", "localhost:8080"), "BASE_URL")
		fileStoragePathPtr := flag.String("f", getEnv("FILE_STORAGE_PATH", ""), "FILE_STORAGE_PATH")
		flag.Parse()

		instance = &Config{
			ServerAddress:   *addressPtr,
			BaseURL:         *baseURLPtr,
			FileStoragePath: *fileStoragePathPtr,
		}
	})
	log.Println("Successfully load config from env variables")
	return instance
}
