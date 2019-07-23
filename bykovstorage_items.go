package bykovstorage

import "log"

func (storage *StorageSingleton) checkReadiness() error {
	if storage.sencObject == nil || storage.sencObject.IsReady() == false {
		return formError("Encrypter is not initialized", "checkReadiness")
	}

	if storage.sdbObject == nil {
		return formError("Database is not initialized", "checkReadiness")
	}

	return nil
}

// AddNewItem - adds new item
func (storage *StorageSingleton) AddNewItem(addItemParms JSONInputAddItem) (response JSONResponseAddItem, err error) {

	err = storage.checkReadiness()
	if err != nil {
		log.Println("Check readiness: " + err.Error())
		return response, err
	}

	encrypredItemName, err := storage.sencObject.Encrypt(addItemParms.ItemName)
	if err != nil {
		log.Println("Encryot item name: " + err.Error())
		return response, err
	}

	encrypredIconName, err := storage.sencObject.Encrypt(addItemParms.ItemIcon)
	if err != nil {
		log.Println("Encryot item icon: " + err.Error())
		return response, err
	}

	err = storage.sdbObject.StartTransaction()
	if err != nil {
		log.Println("Start DB transaction: " + err.Error())
		return response, err
	}

	response.ItemID, err = storage.sdbObject.insertItem(encrypredItemName, encrypredIconName)
	if err != nil {
		log.Println("Insert item into DB: " + err.Error())
		storage.sdbObject.EndTransaction(false)
		return response, err
	}

	err = storage.sdbObject.EndTransaction(true)
	if err != nil {
		log.Println("Finish transaction: " + err.Error())
		return response, err
	}

	response.Status = ConstSuccessResponse

	return response, nil
}

// ReadAllItems - read all not deleted items from the database and decrypt them
func (storage *StorageSingleton) ReadAllItems() (items []BSItem, err error) {
	err = storage.checkReadiness()
	if err != nil {
		return items, err
	}

	items, err = storage.sdbObject.selectAllItems()
	if err != nil {
		return items, err
	}

	for i := range items {
		items[i].Name, err = storage.sencObject.Decrypt(items[i].Name)
		if err != nil {
			return items, err
		}
		items[i].Icon, err = storage.sencObject.Decrypt(items[i].Icon)
		if err != nil {
			return items, err
		}
	}

	return items, nil
}
