package oxilib

import (
	"github.com/oxipass/oxilib/assets"
	"github.com/oxipass/oxilib/internal/pkg/db"
	"github.com/oxipass/oxilib/internal/pkg/security"
	"github.com/oxipass/oxilib/models"
	"log"
	"sync"
	"time"
)

// StorageSingleton is an entry point to work with the storage
type StorageSingleton struct {
	dbObject   *db.StorageDB
	encObject  *security.OxiEncryptor
	lastAccess time.Time
	language   string
}

var instance *StorageSingleton
var once sync.Once

// GetInstance returns singleton object
func GetInstance() *StorageSingleton {
	once.Do(func() {
		log.Println("Get Instance initiated (should be once per session)")
		instance = &StorageSingleton{}
	})
	if instance != nil {
		instance.language = "en" // Initiate by default with English, can be changed later with SetLang
	}
	return instance
}

func (storage *StorageSingleton) storageAccess() {
	storage.lastAccess = time.Now()
}

// IsNew - checks if db was already initialized
func (storage *StorageSingleton) IsNew() bool {
	return storage.DBVersion() <= 0
}

// IsActive is checking if storage is active
func (storage *StorageSingleton) IsActive() bool {
	if !storage.IsNew() && storage.dbObject.IsOpen() {
		return true
	}
	return false
}

// IsLocked - checks if the storage is locked
func (storage *StorageSingleton) IsLocked() bool {
	if storage.IsActive() && storage.encObject != nil && storage.encObject.IsReady() {
		return false
	}
	return true
}

// DBVersion returns actual version of the db used in the file
func (storage *StorageSingleton) DBVersion() int {
	if storage.dbObject != nil {
		return storage.dbObject.Version()
	}

	return -1
}

// GetAvailableCyphers - get available cyphers as one string
func (storage *StorageSingleton) GetAvailableCyphers() (allCyphers []string) {
	return append(allCyphers, storage.encObject.GetCypherNames()...)
}

// Open - opens the db on specified file
func (storage *StorageSingleton) Open(filePath string) error {
	if storage.dbObject != nil && storage.dbObject.IsOpen() && storage.encObject != nil {
		return nil
	}
	storage.dbObject = new(db.StorageDB)
	err := storage.dbObject.Open(filePath)
	if err != nil {
		return err
	}

	storage.encObject = new(security.OxiEncryptor)
	if !storage.IsNew() {
		err := storage.encObject.Init(storage.dbObject.CryptID)
		if err != nil {
			return err
		}
	}

	storage.storageAccess()
	return nil

}

// Close - closes the storage, cypher and db
func (storage *StorageSingleton) Close() error {
	if storage.dbObject != nil && storage.dbObject.IsOpen() {
		err := storage.dbObject.Close()
		if err != nil {
			return err
		}

		storage.dbObject = nil
		storage.encObject = nil
		return nil
	}
	storage.encObject = nil
	return nil
}

// SupportedLangs - returns list of supported languages
func (storage *StorageSingleton) SupportedLangs() []models.Lang {
	return assets.GetLangs()
}

// SetLang - sets language for the storage
func (storage *StorageSingleton) SetLang(lang string) {
	err := assets.SetLang(lang)
	if err != nil {
		log.Println(err)
	}
	storage.language = lang
}

// T Returning translation for a specific key
func (storage *StorageSingleton) T(key string) string {
	return assets.T(key)
}

// TODO: QR code generation support
