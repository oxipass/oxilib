package bykovstorage

// Open - opens the database on specified file
func (storage *StorageSingleton) Open(filePath string) error {
	if storage.sdbObject != nil && storage.sdbObject.IsOpen() == true && storage.sencObject != nil {
		return nil
	}
	storage.sdbObject = new(storageDB)
	err := storage.sdbObject.Open(filePath)
	if err != nil {
		return err
	}

	storage.sencObject = new(bsEncryptor)
	if storage.IsNew() == false {
		err := storage.sencObject.Init(storage.sdbObject.cryptID)
		if err != nil {
			return err
		}
	}

	storage.storageAccess()
	return nil

}

// Lock - locks the storage, delete the key and password
func (storage *StorageSingleton) Lock() {
	if storage.sencObject != nil && storage.sdbObject != nil {
		storage.sencObject.Init(storage.sdbObject.cryptID)
	}
}
