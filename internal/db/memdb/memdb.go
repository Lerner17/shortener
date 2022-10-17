package memdb

import (
	"context"
	"fmt"

	"github.com/Lerner17/shortener/internal/config"
	"github.com/Lerner17/shortener/internal/helpers"
	"github.com/Lerner17/shortener/internal/models"
)

var DBInstance *memdb
var cfg = config.Instance()

type memdb struct {
	state []models.URLEntity
}

func init() {
	DBInstance = &memdb{state: make([]models.URLEntity, 0)}
}

func (m *memdb) generateShortURL() string {
	return helpers.StringWithCharset(7)
}

func (m *memdb) CreateURL(uuid string, fullURL string) (string, string, error) {
	shortURL := m.generateShortURL()

	u := models.URLEntity{
		OriginURL:   fullURL,
		ShortURL:    shortURL,
		UserSession: uuid,
		IsDeleted:   false,
	}
	m.state = append(m.state, u)

	return shortURL, fullURL, nil
}

func (m *memdb) GetURL(shortURL string) (string, bool, bool) {
	for _, u := range m.state {
		if u.ShortURL == shortURL {
			return u.OriginURL, u.IsDeleted, true
		}
	}
	return "", false, false
}

func (m *memdb) GetUserURLs(uuid string) models.URLs {
	result := make(models.URLs, 0)

	for _, u := range m.state {
		if u.UserSession == uuid {
			url := models.URL{
				OriginalURL: u.OriginURL,
				ShortURL:    fmt.Sprintf("%s/%s", cfg.BaseURL, u.ShortURL),
			}
			result = append(result, url)
		}
	}
	return result
}

func (m *memdb) DeleteBatchURL(ctx context.Context, shortURLs []string, uuid string) error {

	for _, sh := range shortURLs {
		for i := 0; i < len(m.state); i++ {
			if m.state[i].ShortURL == sh {
				m.state[i].IsDeleted = true
			}
		}
	}
	return nil
}

func (m *memdb) CreateBatchURL(uuid string, urls models.BatchURLs) (models.BatchShortURLs, error) {
	result := make(models.BatchShortURLs, 0)
	for _, u := range urls {
		shortURL := m.generateShortURL()
		m.state = append(m.state, models.URLEntity{
			OriginURL:     u.OriginalURL,
			ShortURL:      shortURL,
			CorrelationID: u.CorrelationID,
		})
		result = append(result, models.BatchShortURL{
			CorrelationID: u.CorrelationID,
			ShortURL:      shortURL,
		})
	}
	return result, nil
}
