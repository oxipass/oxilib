package bykovstorage

import (
	"log"
	"sync"
	"time"
)

// StorageSingleton is an entry point to work with the storage
type StorageSingleton struct {
	dbObject   *storageDB
	encObject  *bsEncryptor
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
	if storage.DBVersion() > 0 {
		return false
	}
	return true
}

// IsActive is checking if storage is active
func (storage *StorageSingleton) IsActive() bool {
	if storage.IsNew() == false && storage.dbObject.IsOpen() == true {
		return true
	}
	return false
}

//IsLocked - checks if the storage is locked
func (storage *StorageSingleton) IsLocked() bool {
	if storage.IsActive() == true && storage.encObject != nil && storage.encObject.IsReady() == true {
		return false
	}
	return true
}

// DBVersion returns actual version of the database used in the file
func (storage *StorageSingleton) DBVersion() int {
	if storage.dbObject != nil {
		return storage.dbObject.Version()
	}

	return -1
}

// GetAvailableCyphers - get available cypers as one string
func (storage *StorageSingleton) GetAvailableCyphers() (allCyphers []string) {
	for _, cypherName := range storage.encObject.getCypherNames() {
		allCyphers = append(allCyphers, cypherName)
	}
	return allCyphers
}

// Open - opens the database on specified file
func (storage *StorageSingleton) Open(filePath string) error {
	if storage.dbObject != nil && storage.dbObject.IsOpen() == true && storage.encObject != nil {
		return nil
	}
	storage.dbObject = new(storageDB)
	err := storage.dbObject.Open(filePath)
	if err != nil {
		return err
	}

	storage.encObject = new(bsEncryptor)
	if storage.IsNew() == false {
		err := storage.encObject.Init(storage.dbObject.cryptID)
		if err != nil {
			return err
		}
	}

	storage.storageAccess()
	return nil

}
