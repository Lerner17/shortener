package memdb

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Lerner17/shortener/internal/config"
	"github.com/Lerner17/shortener/internal/helpers"
	"github.com/Lerner17/shortener/internal/models"
)

var lock = &sync.Mutex{}

var cfg *config.Config = config.GetConfig()

type memdb struct {
	state map[string]map[string]string
}

// {
// 	data: {
// 		uuid: {
// 			short_url: value,
// 			short_url: value,
// 		}
// 	}
// }

func (d *memdb) GetURL(uuid string, shortURL string) (string, bool) {
	return d.Find(uuid, shortURL)
}

func (d *memdb) CreateURL(uuid string, fullURL string) (string, string) {
	return d.Insert(uuid, fullURL)
}

func (d *memdb) GetUserURLs(uuid string) models.URLs {
	rawUrls := d.state[uuid]
	var urls models.URLs
	for k, v := range rawUrls {
		urls = append(urls, models.URL{
			ShortURL:    fmt.Sprintf("%s/%s", cfg.BaseURL, k),
			OriginalURL: v,
		})
	}
	return urls
}

var dbInstance *memdb

func init() {
	dbInstance = &memdb{state: make(map[string]map[string]string)}
}

func NewMemDB() *memdb {
	us := GetInstance()
	return us
}

func (d *memdb) getUniqueID() string {
	var uniqueID string
	for {
		randomCandidate := helpers.StringWithCharset(7)
		if _, ok := d.state[randomCandidate]; !ok {
			uniqueID = randomCandidate
			break
		}
	}

	return uniqueID
}

func GetInstance() *memdb {
	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbInstance == nil {
			dbInstance = &memdb{}
		}
	}
	return dbInstance
}

func (d *memdb) Find(uuid string, key string) (string, bool) {

	if _, ok := d.state[uuid]; !ok {
		d.state[uuid] = map[string]string{}
	}

	if result, ok := d.state[uuid][key]; ok {
		return result, ok
	}
	return "", false
}

func (d *memdb) InsertWithKey(uuid string, key, value string) (string, error) {
	if key == "" || value == "" {
		return "", errors.New("empty key or value")
	}
	url := make(map[string]string)
	url[key] = value
	d.state[uuid] = url
	return value, nil
}

func (d *memdb) Insert(uuid string, value string) (string, string) {

	if _, ok := d.state[uuid]; !ok {
		d.state[uuid] = map[string]string{}
	}

	uniqueID := d.getUniqueID()
	d.state[uuid][uniqueID] = value

	return uniqueID, d.state[uuid][uniqueID]
}
