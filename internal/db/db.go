package db

import (
	"github.com/Lerner17/shortener/internal/config"
	filestorage "github.com/Lerner17/shortener/internal/db/file_storage"
	"github.com/Lerner17/shortener/internal/db/memdb"
	"github.com/Lerner17/shortener/internal/db/psql"
	"github.com/Lerner17/shortener/internal/models"
)

type URLStorage interface {
	GetURL(string, string) (string, bool)
	CreateURL(string, string) (string, string)
	GetUserURLs(string) models.URLs
}

func GetDB() URLStorage {
	cfg := config.GetConfig()

	if cfg.DatabaseDsn != "" {
		psql := psql.NewPostgres()
		psql.Migrate()
		return psql
	}
	if cfg.FileStoragePath != "" {
		return filestorage.NewFileStorage(cfg.FileStoragePath)
	}
	return memdb.NewMemDB()
}
