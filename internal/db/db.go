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
	CreateURL(string, string) (string, string, error)
	GetUserURLs(string) models.URLs
	CreateBatchURL(string, models.BatchURLs) (models.BatchShortURLs, error)
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
	return memdb.DBInstance
}
