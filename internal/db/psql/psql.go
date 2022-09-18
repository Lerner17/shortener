package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Lerner17/shortener/internal/config"
	er "github.com/Lerner17/shortener/internal/errors"
	"github.com/Lerner17/shortener/internal/helpers"
	"github.com/Lerner17/shortener/internal/logger"
	"github.com/Lerner17/shortener/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
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

func (d *Database) CreateURL(uuid string, fullURL string) (string, string, error) {
	ctx := context.Background()
	shortURL := d.getUniqueID()
	query := "INSERT INTO short_links(short_url, full_url, user_session) VALUES($1, $2, $3)"
	logger.Info(
		"Try to insert URL into table",
		zap.String("shortURL", shortURL),
		zap.String("uuid", uuid),
		zap.String("fullURL", fullURL))
	_, err := d.cursor.ExecContext(ctx, query, shortURL, fullURL, uuid)
	if err != nil {
		logger.Error("Cannot insert to table", zap.Error(err))
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return shortURL, fullURL, er.ErrorShortLinkAlreadyExists
		}
		return shortURL, fullURL, err
	}
	return shortURL, fullURL, nil

}

func (d *Database) CreateBatchURL(uuid string, urls models.BatchURLs) (models.BatchShortURLs, error) {
	result := make(models.BatchShortURLs, 0)
	cfg := config.GetConfig()
	tx, err := d.cursor.Begin()

	if err != nil {
		return result, err
	}
	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	stmt, err := tx.PrepareContext(context.Background(), `
		INSERT INTO
			short_links(
				full_url,
				correlation_id,
				short_url,
				user_session
			) VALUES($1, $2, $3, $4)`)
	if err != nil {
		fmt.Println("PrepareContext ", err)
		return result, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			logger.Info("Close statement error", zap.Error(err))
		}
	}(stmt)

	for _, u := range urls {
		shortURL := d.getUniqueID()
		if _, err := stmt.ExecContext(context.Background(), u.OriginalURL, u.CorrelationId, shortURL, uuid); err == nil {
			result = append(result, models.BatchShortURL{
				CorrelationId: u.CorrelationId,
				ShortURL:      fmt.Sprintf("%s/%s", cfg.BaseURL, shortURL),
			})
		} else {
			fmt.Println("ExecContext ", err.Error())
			return result, err
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatalf("Failed to commit! Panic! %v", err)
	}
	return result, nil
}

func (d *Database) GetUserURLs(uuid string) models.URLs {

	urls := make(models.URLs, 0)
	ctx := context.Background()
	query := "SELECT full_url, short_url FROM short_links WHERE user_session = $1;"

	rows, err := d.cursor.QueryContext(ctx, query, uuid)

	if err != nil {
		logger.Error("Failed to get URL from database", zap.Error(err), zap.String("uuid", uuid))
		return urls
	}
	defer rows.Close()

	for rows.Next() {
		var u models.URL

		err = rows.Scan(&u.OriginalURL, &u.ShortURL)

		if err != nil {
			return urls
		}
		urls = append(urls, u)
	}

	return urls
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

		ALTER TABLE short_links ADD COLUMN IF NOT EXISTS  correlation_id VARCHAR(255) null;
		alter table short_links ADD UNIQUE (user_session, full_url);
	`
	logger.Info("Try to make migration", zap.String("query", query))
	_, err := d.cursor.ExecContext(context, query)

	if err != nil {
		logger.Error("Migration failed", zap.Error(err))
	}
}

func (d *Database) GetURL(uuid string, shortURL string) (string, bool) {

	var url string

	query := "SELECT full_url FROM short_links WHERE short_url = $2"
	err := d.cursor.QueryRow(query, shortURL).Scan(&url)
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
