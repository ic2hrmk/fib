package model

import (
	"os"
	"sync"
	"path/filepath"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"github.com/ic2hrmk/fib/common"
)

type Storage struct {
	storagePath string

	//	Encryption properties
	isEncrypted   bool
	encryptionKey string

	mutex sync.Mutex
}

func NewStorage(filePath string) *Storage {
	return &Storage{
		storagePath:   filePath,
		isEncrypted:   false,
	}
}

func NewEncryptedStorage(filePath string, key string) *Storage {
	return &Storage{
		storagePath:   filePath,
		isEncrypted:   true,
		encryptionKey: key,
	}
}

func initStorage(filePath string) {
	dirPath, _ := filepath.Split(filePath)
	os.MkdirAll(dirPath, 777)
}

func (s *Storage) Add(record interface{}) (err error){
	s.mutex.Lock()
	defer s.mutex.Unlock()

	f, err := os.OpenFile(s.storagePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 777)
	defer f.Close()
	if err != nil {
		err = fmt.Errorf("failed to open file on write, %s", err.Error())
		return
	}

	var serializedData []byte
	serializedData, err = serialize(record)
	if err != nil {
		err = fmt.Errorf("failed to serialize data, %s", err.Error())
		return
	}

	if s.isEncrypted {
		serializedData, err = common.AESEncrypt([]byte(s.encryptionKey), serializedData)
		if err != nil {
			err = fmt.Errorf("failed to encrypt data, %s", err.Error())
			return
		}
	}

	_, err = f.Write(serializedData)
	if err != nil {
		err = fmt.Errorf("failed to write data, %s", err.Error())
		return
	}

	return
}

func serialize(record interface{}) (data []byte, err error) {
	return json.Marshal(record)
}