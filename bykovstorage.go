package bykovstorage

import (
	"log"
	"sync"
	"time"
)

// StorageSingleton is an entry point to work with the storage
type StorageSingleton struct {
	sdbObject  *storageDB
	sencObject *bsEncryptor
	lastAccess time.Time
}

var instance *StorageSingleton
var once sync.Once

// GetInstance returns singleton object
func GetInstance() *StorageSingleton {
	once.Do(func() {
		log.Println("Get Instance initiated (should be once per session)")
		instance = &StorageSingleton{}
	})
	return instance
}

func (storage *StorageSingleton) storageAccess() {
	storage.lastAccess = time.Now()
}

// IsNew - checks if database was already initialized
func (storage *StorageSingleton) IsNew() bool {
	if storage.DBversion() > 0 {
		return false
	}
	return true
}

// IsActive is checking if storage is active
func (storage *StorageSingleton) IsActive() bool {
	if storage.IsNew() == false && storage.sdbObject.IsOpen() == true {
		return true
	}
	return false
}

//IsLocked - checks if the storage is locked
func (storage *StorageSingleton) IsLocked() bool {
	if storage.IsActive() == true && storage.sencObject != nil && storage.sencObject.IsReady() == true {
		return false
	}
	return true
}

// DBversion returns actual version of the database used in the file
func (storage *StorageSingleton) DBversion() int {
	if storage.sdbObject != nil {
		return storage.sdbObject.Version()
	}

	return -1
}

// GetAvailableCyphers - get available cypers as one string
func (storage *StorageSingleton) GetAvailableCyphers() (allCyphers []string) {
	for _, cypherName := range storage.sencObject.getCypherNames() {
		allCyphers = append(allCyphers, cypherName)
	}
	return allCyphers
}
