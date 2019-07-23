package bykovstorage

import "strconv"

// SetNewPassword sets new master password for existing database or generates all the data for the new one
func (storage *StorageSingleton) SetNewPassword(newPassword string, encMethod string) error {
	if storage.IsNew() {
		if len(encMethod) == 0 {
			return formError(BSERR00011EncCypherNotProvided)
		}
		return storage.initStorage(newPassword, encMethod)
	}
	return storage.changePassword(newPassword)
}

// SetPassword sets & checks the password for the open storage
func (storage *StorageSingleton) SetPassword(password string) error {
	if storage.IsActive() == false {
		return formError(BSERR00009DbNotOpen, "SetPassword", "IsActive", strconv.FormatBool(storage.IsActive()))
	}
	if storage.sencObject == nil {
		storage.sencObject = new(bsEncryptor)
	}
	err := storage.sencObject.Init(storage.sdbObject.cryptID)
	if err != nil {
		return err
	}

	passErr := storage.sencObject.SetPassword(password)
	if passErr != nil {
		return passErr
	}

	dbID, encErr := storage.sencObject.Decrypt(storage.sdbObject.keyWord)
	if encErr != nil {
		return encErr
	}

	if dbID != storage.sdbObject.dbID {
		storage.sencObject.Init(storage.sdbObject.cryptID) //if password is wrong, clean everything
		return formError(BSERR00010EncWrongPassword)
	}

	return nil
}

func (storage *StorageSingleton) initStorage(newPassword string, encMethod string) error {

	storage.sencObject = new(bsEncryptor)
	cryptID, err := storage.sencObject.getCryptIDbyName(encMethod)
	if err != nil {
		return err
	}

	err = storage.sencObject.Init(cryptID)
	if err != nil {
		return err
	}

	dbID := generateRandomString(DatabaseIDLength)

	err = storage.sencObject.SetPassword(newPassword)
	if err != nil {
		return err
	}
	keyWord, encErr := storage.sencObject.Encrypt(dbID)
	if encErr != nil {
		return encErr
	}

	err = storage.sdbObject.StartTransaction()
	if err != nil {
		return err
	}

	err = storage.sdbObject.initSettings(dbID, keyWord, storage.sencObject.cryptID)
	if err != nil {
		storage.sdbObject.EndTransaction(false)
		return err
	}

	err = storage.sdbObject.EndTransaction(true)
	if err != nil {
		return err
	}

	err = storage.sdbObject.checkIntegrityGetSettings()
	if err != nil {
		return err
	}

	return nil
}

func (storage *StorageSingleton) changePassword(newPassword string) error {
	return nil
}
