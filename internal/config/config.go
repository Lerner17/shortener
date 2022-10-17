package config

import (
	"errors"
)

var ErrUnknownParam = errors.New("unknown param")

// Config project
type Config struct {
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"unknown"`
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:":8080"`
	DatabaseDsn     string `env:"DATABASE_DSN" envDefault:""`
}

const (
	BaseURL                = "BASE_URL"
	ServerAddress          = "SERVER_ADDRESS"
	FileStoragePath        = "FILE_STORAGE_PATH"
	DatabaseDsn            = "DATABASE_DSN"
	FileStoragePathDefault = "unknown"
)

// Maps for take inv params
var mapVarToInv = map[string]string{
	BaseURL:         "b",
	ServerAddress:   "a",
	FileStoragePath: "f",
	DatabaseDsn:     "d",
}

var instance *Config

// Instance new Config
func Instance() *Config {
	if instance == nil {
		instance = new(Config)
		instance.initInv()
	}
	return instance
}

// Param from configs
func (c *Config) Param(p string) (string, error) {
	switch p {
	case BaseURL:
		return c.BaseURL, nil
	case ServerAddress:
		return c.ServerAddress, nil
	case FileStoragePath:
		return c.FileStoragePath, nil
	case DatabaseDsn:
		return c.DatabaseDsn, nil
	}
	return "", ErrUnknownParam
}

// initInv check from inv
func (c *Config) initInv() {
	// Get from inv
	// bu := flag.String(mapVarToInv[BaseURL], "", "")
	// sa := flag.String(mapVarToInv[ServerAddress], "", "")
	// fs := flag.String(mapVarToInv[FileStoragePath], "", "")
	// db := flag.String(mapVarToInv[DatabaseDsn], "", "")
	// flag.Parse()

	// if *bu != "" {
	// 	c.BaseURL = *bu
	// }
	// if *sa != "" {
	// 	c.ServerAddress = *sa
	// }
	// if *fs != "" {
	// 	c.FileStoragePath = *fs
	// }
	// if *db != "" {
	// 	c.DatabaseDsn = *db
	// }
	// if err := env.Parse(c); err != nil {
	// 	return
	// }
}

// var instance *Config

// func (c Config) Parse() {
// 	if err := env.Parse(c); err != nil {
// 		fmt.Printf("Cannot parse env vars %v\n", err)
// 	}
// 	serverAddressPtr := flag.String("a", "", "")
// 	baseURLPtr := flag.String("b", "", "")
// 	fileStoragePathPtr := flag.String("f", "", "")
// 	DatabaseDsnPtr := flag.String("d", "", "")

// 	flag.Parse()

// 	if *serverAddressPtr != "" {
// 		c.ServerAddress = *serverAddressPtr
// 	}

// 	if *baseURLPtr != "" {
// 		c.BaseURL = *baseURLPtr
// 	}

// 	if *fileStoragePathPtr != "" {
// 		c.FileStoragePath = *fileStoragePathPtr
// 	}

// 	if *DatabaseDsnPtr != "" {
// 		c.DatabaseDsn = *DatabaseDsnPtr
// 	}
// }

// var once sync.Once

// func Instance() *Config {
// 	once.Do(func() {
// 		log.Println("Load config...")
// 		instance = new(Config)
// 		instance.Parse()
// 		log.Println("Successfully load config from env variables")
// 	})
// 	return instance
// }
