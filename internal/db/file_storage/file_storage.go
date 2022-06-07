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
	"github.com/Lerner17/shortener/internal/models"
)

// var _ URLStorage = &fileStorage{}
var cfg *config.Config = config.GetConfig()

type fileStorage struct {
	state  map[string]map[string]string
	file   *os.File
	writer *bufio.Writer
}

func NewFileStorage(dbPath string) *fileStorage {
	file, err := os.OpenFile(dbPath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(fmt.Sprintf("Cannot open db file: %v", err))
	}

	var data map[string]map[string]string

	byteValue, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		data = make(map[string]map[string]string)
	}

	fmt.Println(data)

	return &fileStorage{
		state:  data,
		file:   file,
		writer: bufio.NewWriter(file),
	}

}

func (fs *fileStorage) Close() error {
	return fs.file.Close()
}

func (fs *fileStorage) GetURL(uuid string, shortURL string) (string, bool) {
	userState := fs.state[uuid]
	if result, ok := userState[shortURL]; ok {
		return result, ok
	}
	return "", false
}

func (fs *fileStorage) GetUserURLs(uuid string) models.URLs {
	rawUrls := fs.state[uuid]
	var urls models.URLs
	for k, v := range rawUrls {
		urls = append(urls, models.URL{
			ShortURL:    fmt.Sprintf("%s/%s", cfg.BaseURL, k),
			OriginalURL: v,
		})
	}
	return urls
}

func (fs *fileStorage) getUniqueID() string {
	var uniqueID string
	for {
		randomCandidate := helpers.StringWithCharset(7)
		if _, ok := fs.state[randomCandidate]; !ok {
			uniqueID = randomCandidate
			break
		}
	}

	return uniqueID
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

func (fs *fileStorage) CreateURL(uuid string, fullURL string) (string, string) {
	urls := make(map[string]string)

	_, ok := fs.state[uuid]
	if ok {
		urls = fs.state[uuid]
	}

	uniqueID := fs.getUniqueID()
	urls[uniqueID] = fullURL
	fs.state[uuid] = urls
	err := fs.writeState()
	if err != nil {
		fmt.Printf("Cannot write state to file: %v\n", err)
	}
	return uniqueID, fullURL
}
