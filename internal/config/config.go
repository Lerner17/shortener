package config

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	SecretKey       string `env:"SECRET_KEY" envDefault:"qwerty"`
}

var instance *Config

func (c *Config) init() {
	if err := env.Parse(c); err != nil {
		fmt.Printf("Cannot parse env vars %v\n", err)
	}
	serverAddressPtr := flag.String("a", "", "")
	baseURLPtr := flag.String("b", "", "")
	fileStoragePathPtr := flag.String("f", "", "")
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
}

var once sync.Once

func GetConfig() *Config {
	log.Println("Load config...")
	once.Do(func() {
		instance = new(Config)
		instance.init()
	})
	log.Println("Successfully load config from env variables")
	return instance
}
