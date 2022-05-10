package db

import (
	"errors"
	"sync"

	"github.com/Lerner17/shortener/internal/helpers"
)

var lock = &sync.Mutex{}

// var _ URLStorage = &memdb{}

type URLStorage interface {
	GetURL(string) (string, bool)
	CreateURL(string) (string, string)
}

func (d *memdb) GetURL(shortURL string) (string, bool) {
	return d.Find(shortURL)
}

func (d *memdb) CreateURL(fullURL string) (string, string) {
	return d.Insert(fullURL)
}

type memdb struct {
	state map[string]string
}

var dbInstance *memdb

func init() {
	dbInstance = &memdb{state: make(map[string]string)}
}

func NewURLStorage() URLStorage {
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

func (d *memdb) Find(key string) (string, bool) {
	if result, ok := d.state[key]; ok {
		return result, ok
	}
	return "", false
}

func (d *memdb) InsertWithKey(key, value string) (string, error) {
	if key == "" || value == "" {
		return "", errors.New("empty key or value")
	}
	d.state[key] = value
	return value, nil
}

func (d *memdb) Insert(value string) (string, string) {
	uniqueID := d.getUniqueID()
	d.state[uniqueID] = value
	return uniqueID, d.state[uniqueID]
}
