package psql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Lerner17/shortener/internal/config"
	"github.com/Lerner17/shortener/internal/logger"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var Instance *sql.DB

func init() {
	dsn := config.GetConfig().DatabaseDsn
	if dsn == "" {
		fmt.Fprint(os.Stderr, "Cannot connect to database")
	}

	inst, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}
	Instance = inst
	logger.Info("Connect to database")

}
