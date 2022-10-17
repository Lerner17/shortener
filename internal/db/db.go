package db

import (
	"context"

	"github.com/Lerner17/shortener/internal/config"
	filestorage "github.com/Lerner17/shortener/internal/db/file_storage"

	"github.com/Lerner17/shortener/internal/db/memdb"
	"github.com/Lerner17/shortener/internal/db/psql"
	"github.com/Lerner17/shortener/internal/logger"
	"github.com/Lerner17/shortener/internal/models"
)

type URLStorage interface {
	GetURL(string) (string, bool, bool)
	CreateURL(string, string) (string, string, error)
	GetUserURLs(string) models.URLs
	CreateBatchURL(string, models.BatchURLs) (models.BatchShortURLs, error)
	DeleteBatchURL(context.Context, []string, string) error
}

func GetDB() URLStorage {
	cfg := config.Instance()

	if cfg.DatabaseDsn != "" {
		logger.Info("using postgres database")
		psql := psql.NewPostgres()
		psql.Migrate()
		return psql
	}
	if cfg.FileStoragePath != "" {
		logger.Info("using file storage")
		return filestorage.NewFileStorage(cfg.FileStoragePath)
	}
	logger.Info("using memory database")
	return memdb.DBInstance
}
