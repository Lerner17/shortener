package psql

import (
	"database/sql"
	"errors"

	"github.com/Lerner17/shortener/internal/config"
	"github.com/Lerner17/shortener/internal/logger"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var Instance *sql.DB

func init() {
	dsn := config.GetConfig().DatabaseDsn
	if dsn == "" {
		panic(errors.New("DB error"))
	}

	inst, err := sql.Open("pgx", dsn+"?sslmode=disable")
	if err != nil {
		panic(err)
	}
	Instance = inst
	logger.Info("Connect to database")

}
