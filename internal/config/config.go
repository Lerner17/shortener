package config

import (
	"log"
	"sync"

	"github.com/alexflint/go-arg"
)

type Config struct {
	ServerAddress   string `arg:"-a,env:SERVER_ADDRESS" default:"localhost:8080"`
	BaseURL         string `arg:"-b,env:BASE_URL" default:"localhost:8080"`
	FileStoragePath string `arg:"-f,env:FILE_STORAGE_PATH" default:""`
}

var instance *Config

var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Println("Load config...")
		instance = &Config{}
		arg.MustParse(instance)
	})
	log.Println("Successfully load config from env variables")
	return instance
}
