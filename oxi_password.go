package oxilib

import (
	"log"
	"strconv"
)

// SetNewPassword sets new master password for existing database or generates all the data for the new one
func (storage *StorageSingleton) SetNewPassword(newPassword string, encMethod string) error {
	if storage.IsNew() {
		if len(encMethod) == 0 {
			return formError(BSERR00011EncCypherNotProvided)
		}
		return storage.initStorage(newPassword, encMethod, storage.language)
	}
	return storage.changePassword(newPassword)
}

// Unlock sets & checks the password for the open storage
func (storage *StorageSingleton) Unlock(password string) error {
	if !storage.IsActive() {
		return formError(BSERR00009DbNotOpen, "Unlock", "IsActive", strconv.FormatBool(storage.IsActive()))
	}
	if storage.encObject == nil {
		storage.encObject = new(bsEncryptor)
	}
	err := storage.encObject.Init(storage.dbObject.cryptID)
	if err != nil {
		return err
	}

	passErr := storage.encObject.SetPassword(password)
	if passErr != nil {
		return passErr
	}

	dbID, encErr := storage.encObject.Decrypt(storage.dbObject.keyWord)
	if encErr != nil {
		return encErr
	}

	if dbID != storage.dbObject.dbID {
		err := storage.encObject.Init(storage.dbObject.cryptID) //if password is wrong, clean everything
		if err != nil {
			return formError(BSERR00010EncWrongPassword, err.Error())
		}
		return formError(BSERR00010EncWrongPassword)
	}

	return nil
}

func (storage *StorageSingleton) initStorage(newPassword string, encMethod string, lang string) error {

	storage.encObject = new(bsEncryptor)
	cryptID, err := storage.encObject.getCryptIDbyName(encMethod)
	if err != nil {
		return err
	}

	err = storage.encObject.Init(cryptID)
	if err != nil {
		return err
	}

	dbID := generateRandomString(DatabaseIDLength)

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

	err = storage.dbObject.initSettings(dbID, keyWord, storage.encObject.cryptID, lang)
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

	err = storage.dbObject.getSettings()
	if err != nil {
		return err
	}

	return nil
}

func (storage *StorageSingleton) changePassword(newPassword string) error {
	return nil
}

// Lock - locks the storage, delete the key and password
func (storage *StorageSingleton) Lock() error {
	if storage.encObject != nil && storage.dbObject != nil {
		return storage.encObject.Init(storage.dbObject.cryptID)
	}
	return formError(BSERR00009DbNotOpen, "func Lock()")
}
