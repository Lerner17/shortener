package fileStorage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/Lerner17/shortener/internal/helpers"
)

// var _ URLStorage = &fileStorage{}

type fileStorage struct {
	state  map[string]string
	file   *os.File
	writer *bufio.Writer
}

func NewFileStorage(dbPath string) *fileStorage {
	file, err := os.OpenFile(dbPath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(fmt.Sprintf("Cannot open db file: %v", err))
	}

	var data map[string]string

	byteValue, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		data = make(map[string]string)
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

func (fs *fileStorage) GetURL(shortURL string) (string, bool) {
	if result, ok := fs.state[shortURL]; ok {
		return result, ok
	}
	return "", false
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

func (fs *fileStorage) CreateURL(fullURL string) (string, string) {
	key := fs.getUniqueID()
	fs.state[key] = fullURL
	err := fs.writeState()
	if err != nil {
		fmt.Printf("Cannot write state to file: %v\n", err)
	}
	return key, fullURL
}
