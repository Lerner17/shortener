package psql

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/Lerner17/shortener/internal/config"
	"github.com/Lerner17/shortener/internal/helpers"
	"github.com/Lerner17/shortener/internal/logger"
	"github.com/Lerner17/shortener/internal/models"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
)

var instance *Database

var GetInstance = NewPostgres

type NullString struct {
	String string
	Valid  bool
}
type Database struct {
	cursor *sql.DB
}

func (d *Database) getUniqueID() string {
	return helpers.StringWithCharset(7)
}

func (d *Database) CreateURL(uuid string, fullURL string) (string, string) {
	ctx := context.Background()
	shortURL := d.getUniqueID()
	query := "INSERT INTO short_links(short_url, full_url, user_session) VALUES($1, $2, $3) returning id"
	logger.Info(
		"Try to insert URL into table",
		zap.String("shortURL", shortURL),
		zap.String("uuid", uuid),
		zap.String("fullURL", fullURL))
	_, err := d.cursor.ExecContext(ctx, query, shortURL, fullURL, uuid)
	if err != nil {
		logger.Error("Cannot insert to table", zap.Error(err))
	}
	return shortURL, fullURL

}

func (d *Database) GetUserURLs(string) models.URLs {
	return models.URLs{}
}

func NewPostgres() *Database {
	return instance
}

func (d *Database) Ping() error {
	return d.cursor.Ping()
}

func (d *Database) Migrate() {

	context := context.Background()

	query := `
		CREATE TABLE IF NOT EXISTS short_links(
			id serial PRIMARY KEY,
			short_url VARCHAR ( 16 ) UNIQUE NOT NULL,
			full_url TEXT NOT NULL,
			user_session UUID
		);
	`
	logger.Info("Try to make migration", zap.String("query", query))
	_, err := d.cursor.ExecContext(context, query)

	if err != nil {
		logger.Error("Migration failed", zap.Error(err))
	}
}

func (d *Database) GetURL(uuid string, shortURL string) (string, bool) {

	var url string

	query := "SELECT full_url FROM short_links WHERE user_session = $1 AND short_url = $2"
	err := d.cursor.QueryRow(query, uuid, shortURL).Scan(&url)
	if err != nil {
		logger.Error("Failed to get URL from database", zap.Error(err), zap.String("shortURL", shortURL), zap.String("uuid", uuid))
		return "", false
	}
	return url, true
}

func init() {
	dsn := config.GetConfig().DatabaseDsn
	if dsn == "" {
		fmt.Fprint(os.Stderr, "Cannot connect to database")
	}
	cursor, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}
	instance = &Database{
		cursor: cursor,
	}
	logger.Info("Connect to database")

}
