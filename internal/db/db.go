package db

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var lock = &sync.Mutex{}
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

type db struct {
	state map[string]string
}

var dbInstance *db

func init() {
	dbInstance = &db{state: make(map[string]string)}
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (d *db) getUniqueId() string {
	var uniqueID string
	for {
		var randomCandidate string = stringWithCharset(7, charset)
		if _, ok := d.state[randomCandidate]; !ok {
			uniqueID = randomCandidate
			break
		}
	}

	return uniqueID
}

func GetInstance() *db {
	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbInstance == nil {
			dbInstance = &db{}
		}
	}
	return dbInstance
}

func (d *db) Find(key string) (string, bool) {
	if result, ok := d.state[key]; ok {
		return result, ok
	}
	return "", false
}

func (d *db) InsertWithKey(key, value string) (string, error) {
	if key == "" || value == "" {
		return "", errors.New("empty key or value")
	}
	d.state[key] = value
	return value, nil
}

func (d *db) Insert(value string) (string, string) {
	uniqueID := d.getUniqueId()
	d.state[uniqueID] = value
	return uniqueID, d.state[uniqueID]
}
