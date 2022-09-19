package memdb

import (
	"fmt"

	"github.com/Lerner17/shortener/internal/helpers"
	"github.com/Lerner17/shortener/internal/models"
)

var DBInstance *memdb

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
	fmt.Println("================================================================")
	fmt.Println(uuid)
	fmt.Println(fullURL)
	shortURL := m.generateShortURL()

	u := models.URLEntity{
		OriginURL:   fullURL,
		ShortURL:    shortURL,
		UserSession: uuid,
	}
	m.state = append(m.state, u)

	return shortURL, fullURL, nil
}

func (m *memdb) GetURL(uuid string, shortURL string) (string, bool) {
	for _, u := range m.state {
		if u.UserSession == uuid && u.ShortURL == shortURL {
			return u.OriginURL, true
		}
	}
	return "", false
}

func (m *memdb) GetUserURLs(uuid string) models.URLs {
	result := make(models.URLs, 0)

	for _, u := range m.state {
		if u.UserSession == uuid {
			url := models.URL{
				OriginalURL: u.OriginURL,
				ShortURL:    u.ShortURL,
			}
			result = append(result, url)
		}
	}
	return result
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
