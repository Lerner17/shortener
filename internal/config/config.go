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
	DatabaseDsn     string `env:"DATABASE_DSN"`
}

var instance *Config

func init() {
	instance = new(Config)
	if err := env.Parse(instance); err != nil {
		fmt.Printf("Cannot parse env vars %v\n", err)
	}
	serverAddressPtr := flag.String("a", "", "")
	baseURLPtr := flag.String("b", "", "")
	fileStoragePathPtr := flag.String("f", "", "")
	DatabaseDsnPtr := flag.String("d", "", "")
	flag.Parse()

	if *serverAddressPtr != "" {
		instance.ServerAddress = *serverAddressPtr
	}

	if *baseURLPtr != "" {
		instance.BaseURL = *baseURLPtr
	}

	if *fileStoragePathPtr != "" {
		instance.FileStoragePath = *fileStoragePathPtr
	}

	if *DatabaseDsnPtr != "" {
		instance.DatabaseDsn = *DatabaseDsnPtr
	}
}

var once sync.Once

func GetConfig() *Config {
	log.Println("Load config...")
	log.Println("Successfully load config from env variables")
	return instance
}
