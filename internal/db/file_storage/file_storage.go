package filestorage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/Lerner17/shortener/internal/config"
	"github.com/Lerner17/shortener/internal/helpers"
	"github.com/Lerner17/shortener/internal/logger"
	"github.com/Lerner17/shortener/internal/models"
	"go.uber.org/zap"
)

var cfg = config.GetConfig()

type fileStorage struct {
	state  []models.URLEntity
	file   *os.File
	writer *bufio.Writer
}

func (fs *fileStorage) writeState() error {
	fs.file.Truncate(0)
	fs.file.Seek(0, io.SeekStart)
	data, err := json.Marshal(fs.state)
	if err != nil {
		return err
	}

	if _, err := fs.writer.Write(data); err != nil {
		return err
	}
	return fs.writer.Flush()
}

func NewFileStorage(dbPath string) *fileStorage {
	file, err := os.OpenFile(dbPath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(fmt.Sprintf("Cannot open db file: %v", err))
	}

	var data []models.URLEntity

	byteValue, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		data = make([]models.URLEntity, 0)
	}

	return &fileStorage{
		state:  data,
		file:   file,
		writer: bufio.NewWriter(file),
	}

}

func (fs *fileStorage) Close() error {
	return fs.file.Close()
}

func (fs *fileStorage) generateShortURL() string {
	return helpers.StringWithCharset(7)
}

func (fs *fileStorage) CreateURL(uuid string, fullURL string) (string, string, error) {

	key := fs.generateShortURL()

	url := models.URLEntity{
		OriginURL:     fullURL,
		ShortURL:      key,
		UserSession:   uuid,
		CorrelationID: "",
	}

	fs.state = append(fs.state, url)
	err := fs.writeState()
	if err != nil {
		logger.Error("Cannot write state to file", zap.Error(err))
		return fmt.Sprintf("%s/%s", cfg.BaseURL, key), fullURL, err
	}

	return key, fullURL, nil
}

func (fs *fileStorage) GetURL(uuid string, shortURL string) (string, bool) {
	for _, u := range fs.state {
		if u.UserSession == uuid && u.ShortURL == shortURL {
			return u.OriginURL, true
		}
	}
	return "", false
}

func (fs *fileStorage) GetUserURLs(uuid string) models.URLs {
	result := make(models.URLs, 0)

	for _, u := range fs.state {
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

func (fs *fileStorage) CreateBatchURL(uuid string, urls models.BatchURLs) (models.BatchShortURLs, error) {
	result := make(models.BatchShortURLs, 0)
	for _, u := range urls {
		shortURL := fs.generateShortURL()
		fs.state = append(fs.state, models.URLEntity{
			OriginURL:     u.OriginalURL,
			ShortURL:      shortURL,
			CorrelationID: u.CorrelationID,
		})
		result = append(result, models.BatchShortURL{
			CorrelationID: u.CorrelationID,
			ShortURL:      shortURL,
		})
		err := fs.writeState()
		if err != nil {
			logger.Error("Cannot write state to file", zap.Error(err))
			return result, err
		}
	}
	return result, nil
}
