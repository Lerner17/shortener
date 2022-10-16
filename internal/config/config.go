package config

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/Lerner17/shortener/internal/logger"
	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
)

var once sync.Once

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	DatabaseDsn     string `env:"DATABASE_DSN"`
}

var Instance *Config

func (c *Config) ParseConfig() {
	if err := env.Parse(Instance); err != nil {
		fmt.Printf("Cannot parse env vars %v\n", err)
	}
	serverAddressPtr := flag.String("a", "", "")
	baseURLPtr := flag.String("b", "", "")
	fileStoragePathPtr := flag.String("f", "", "")
	DatabaseDsnPtr := flag.String("d", "", "")
	flag.Parse()
	if *serverAddressPtr != "" {
		Instance.ServerAddress = *serverAddressPtr
	}

	if *baseURLPtr != "" {
		Instance.BaseURL = *baseURLPtr
	}

	if *fileStoragePathPtr != "" {
		Instance.FileStoragePath = *fileStoragePathPtr
	}

	if *DatabaseDsnPtr != "" {
		Instance.DatabaseDsn = *DatabaseDsnPtr
	}

	logger.Info("load configuration for application",
		zap.String("server address", *serverAddressPtr),
		zap.String("base URL", *baseURLPtr),
		zap.String("file storage path", *fileStoragePathPtr),
		zap.String("database dns url", *DatabaseDsnPtr),
	)
}

func GetConfig() *Config {
	log.Println("Load config...")
	once.Do(func() {
		Instance = new(Config)
		Instance.ParseConfig()
	})
	log.Println("Successfully load config from env variables")
	return Instance
}
