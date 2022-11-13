package oxilib

import (
	"github.com/oxipass/oxilib/consts"
	"github.com/oxipass/oxilib/internal/pkg/oxierr"
	"github.com/oxipass/oxilib/internal/pkg/security"
	"github.com/oxipass/oxilib/internal/pkg/utils"
	"log"
	"strconv"
)

// SetNewPassword sets new master password for existing db or generates all the data for the new one
func (storage *StorageSingleton) SetNewPassword(newPassword string, encMethod string) error {
	if storage.IsNew() {
		if len(encMethod) == 0 {
			return oxierr.FormError(oxierr.BSERR00011EncCypherNotProvided)
		}
		return storage.initStorage(newPassword, encMethod, storage.language)
	}
	return storage.changePassword(newPassword)
}

// Unlock sets & checks the password for the open storage
func (storage *StorageSingleton) Unlock(password string) error {
	if !storage.IsActive() {
		return oxierr.FormError(oxierr.BSERR00009DbNotOpen, "Unlock", "IsActive", strconv.FormatBool(storage.IsActive()))
	}
	if storage.encObject == nil {
		storage.encObject = new(security.OxiEncryptor)
	}
	err := storage.encObject.Init(storage.dbObject.CryptID)
	if err != nil {
		return err
	}

	passErr := storage.encObject.SetPassword(password)
	if passErr != nil {
		return passErr
	}

	dbID, encErr := storage.encObject.Decrypt(storage.dbObject.KeyWord)
	if encErr != nil {
		return encErr
	}

	if dbID != storage.dbObject.DbID {
		err := storage.encObject.Init(storage.dbObject.CryptID) //if password is wrong, clean everything
		if err != nil {
			return oxierr.FormError(oxierr.BSERR00010EncWrongPassword, err.Error())
		}
		return oxierr.FormError(oxierr.BSERR00010EncWrongPassword)
	}

	return nil
}

func (storage *StorageSingleton) initStorage(newPassword string, encMethod string, lang string) error {

	storage.encObject = new(security.OxiEncryptor)
	cryptID, err := storage.encObject.GetCryptIDbyName(encMethod)
	if err != nil {
		return err
	}

	err = storage.encObject.Init(cryptID)
	if err != nil {
		return err
	}

	dbID := utils.GenerateRandomString(consts.DatabaseIDLength)

	err = storage.encObject.SetPassword(newPassword)
	if err != nil {
		return err
	}
	keyWord, encErr := storage.encObject.Encrypt(dbID)
	if encErr != nil {
		return encErr
	}

	err = storage.dbObject.StartTX()
	if err != nil {
		return err
	}

	err = storage.dbObject.InitSettings(dbID, keyWord, storage.encObject.CryptID, lang)
	if err != nil {
		errRollback := storage.dbObject.RollbackTX()
		if errRollback != nil {
			log.Println(errRollback.Error())
		}
		return err
	}

	err = storage.dbObject.CommitTX()
	if err != nil {
		return err
	}

	err = storage.dbObject.GetSettings()
	if err != nil {
		return err
	}

	return nil
}

func (storage *StorageSingleton) FillEmptyStorage() error {
	errTags := storage.AddDefaultTags()
	if errTags != nil {
		return errTags
	}
	errTemplates := storage.AddDefaultItemsTemplates()
	if errTemplates != nil {
		return errTemplates
	}
	return nil
}

func (storage *StorageSingleton) changePassword(newPassword string) error {
	return nil
}

// Lock - locks the storage, delete the key and password
func (storage *StorageSingleton) Lock() error {
	if storage.encObject != nil && storage.dbObject != nil {
		return storage.encObject.Init(storage.dbObject.CryptID)
	}
	return oxierr.FormError(oxierr.BSERR00009DbNotOpen, "func Lock()")
}
