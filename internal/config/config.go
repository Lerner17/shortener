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

func (c Config) Parse() {
	if err := env.Parse(c); err != nil {
		fmt.Printf("Cannot parse env vars %v\n", err)
	}
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

var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Println("Load config...")
		instance = new(Config)
		instance.Parse()
		log.Println("Successfully load config from env variables")
	})
	return instance
}
