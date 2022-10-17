package config

import (
	"errors"
	"flag"

	"github.com/caarlos0/env/v6"
)

var ErrUnknownParam = errors.New("unknown param")

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	DatabaseDsn     string `env:"DATABASE_DSN"`
}

const (
	BaseURL                = "BASE_URL"
	ServerAddress          = "SERVER_ADDRESS"
	FileStoragePath        = "FILE_STORAGE_PATH"
	FileStoragePathDefault = "unknown"
)

// Maps for take inv params
var mapVarToInv = map[string]string{
	BaseURL:         "b",
	ServerAddress:   "a",
	FileStoragePath: "f",
}

var instance *Config

// Instance new Config
func GetConfig() *Config {
	if instance == nil {
		instance = new(Config)
		instance.initInv()
		instance.init()
	}
	return instance
}

// Param from configs
func (c *Config) Param(p string) (string, error) {
	switch p {
	case BaseURL:
		if c.BaseURL != "" {
			return c.BaseURL, nil
		}
		return c.BaseURL, nil
	case ServerAddress:
		if c.ServerAddress != "" {
			return c.ServerAddress, nil
		}
		return c.ServerAddress, nil
	case FileStoragePath:
		if c.FileStoragePath != "" {
			return c.FileStoragePath, nil
		}
		return c.FileStoragePath, nil
	}
	return "", ErrUnknownParam
}

// initInv check from inv
func (c *Config) initInv() {
	// Get from inv
	if err := env.Parse(c); err != nil {
		return
	}
}

// initParams from cli params
func (c *Config) init() {
	bu := flag.String(mapVarToInv[BaseURL], "", "")
	sa := flag.String(mapVarToInv[ServerAddress], "", "")
	fs := flag.String(mapVarToInv[FileStoragePath], "", "")
	flag.Parse()

	if *bu != "" {
		c.BaseURL = *bu
	}
	if *sa != "" {
		c.ServerAddress = *sa
	}
	if *fs != "" {
		c.FileStoragePath = *fs
	}
}

// var Instance *Config

// func init() {
// 	Instance = &Config{}
// }

// func (c Config) ParseConfig() {
// 	if err := env.Parse(Instance); err != nil {
// 		fmt.Printf("Cannot parse env vars %v\n", err)
// 	}
// 	serverAddressPtr := flag.String("a", "", "")
// 	baseURLPtr := flag.String("b", "", "")
// 	fileStoragePathPtr := flag.String("f", "", "")
// 	DatabaseDsnPtr := flag.String("d", "", "")
// 	flag.Parse()
// 	if *serverAddressPtr != "" {
// 		Instance.ServerAddress = *serverAddressPtr
// 	}

// 	if *baseURLPtr != "" {
// 		Instance.BaseURL = *baseURLPtr
// 	}

// 	if *fileStoragePathPtr != "" {
// 		Instance.FileStoragePath = *fileStoragePathPtr
// 	}

// 	if *DatabaseDsnPtr != "" {
// 		Instance.DatabaseDsn = *DatabaseDsnPtr
// 	}

// 	logger.Info("load configuration for application",
// 		zap.String("server address", *serverAddressPtr),
// 		zap.String("base URL", *baseURLPtr),
// 		zap.String("file storage path", *fileStoragePathPtr),
// 		zap.String("database dns url", *DatabaseDsnPtr),
// 	)
// }

// func GetConfig() *Config {

// 	return Instance
// }
