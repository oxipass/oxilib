package oxilib

func (storage *StorageSingleton) checkReadiness() error {
	if storage.encObject == nil || !storage.encObject.IsReady() {
		return formError("Encryptor is not initialized", "checkReadiness")
	}

	if storage.dbObject == nil {
		return formError("Database is not initialized", "checkReadiness")
	}

	return nil
}

func (storage *StorageSingleton) DeleteItem(deleteItemParams UpdateItemForm) (response CommonResponse, err error) {
	err = storage.checkReadiness()
	if err != nil {
		return response, err
	}
	err = storage.dbObject.StartTX()
	if err != nil {
		return response, err
	}

	err = storage.dbObject.dbDeleteItem(deleteItemParams.ID)
	if err != nil {
		errEndTX := storage.dbObject.RollbackTX()
		if errEndTX != nil {
			return response, formError(BSERR00016DbDeleteFailed, err.Error(), errEndTX.Error())
		}
		return response, err
	}

	err = storage.dbObject.CommitTX()
	if err != nil {
		return response, err
	}

	response.Status = ConstSuccessResponse

	return response, nil
}

func (storage *StorageSingleton) UpdateItem(updateItemParams UpdateItemForm) (response ItemUpdatedResponse, err error) {
	var encryptedItemName, encryptedIconName string
	err = storage.checkReadiness()
	if err != nil {
		return response, err
	}

	err = ValidateItemBeforeUpdate(updateItemParams)
	if err != nil {
		return response, err
	}

	if updateItemParams.Icon != "" {
		encryptedIconName, err = storage.encObject.Encrypt(updateItemParams.Icon)
		if err != nil {
			return response, err
		}
	}
	if updateItemParams.Name != "" {
		encryptedItemName, err = storage.encObject.Encrypt(updateItemParams.Name)
		if err != nil {
			return response, err
		}
	}

	err = storage.dbObject.StartTX()
	if err != nil {
		return response, err
	}

	if updateItemParams.Name != "" {
		err = storage.dbObject.dbUpdateItemName(updateItemParams.ID, encryptedItemName)
		if err != nil {
			errEndTX := storage.dbObject.RollbackTX()
			if errEndTX != nil {
				return response, formError(BSERR00018DbItemNameUpdateFailed, err.Error(), errEndTX.Error())
			}
			return response, err
		}
	}

	if updateItemParams.Icon != "" {
		err = storage.dbObject.dbUpdateItemIcon(updateItemParams.ID, encryptedIconName)
		if err != nil {
			errEndTX := storage.dbObject.RollbackTX()
			if errEndTX != nil {
				return response, formError(BSERR00026DbItemIconUpdateFailed, err.Error(), errEndTX.Error())
			}
			return response, err
		}
	}

	err = storage.dbObject.CommitTX()
	if err != nil {
		return response, err
	}

	response.Status = ConstSuccessResponse

	return response, nil
}

// AddNewItem - adds new item
func (storage *StorageSingleton) AddNewItem(addItemParams UpdateItemForm) (response ItemAddedResponse, err error) {

	err = storage.checkReadiness()
	if err != nil {
		return response, err
	}

	if !CheckIfExistsFontAwesome(addItemParams.Icon) {
		return response, formError(BSERR00024FontAwesomeIconNotFound)
	}

	encryptedItemName, err := storage.encObject.Encrypt(addItemParams.Name)
	if err != nil {
		return response, err
	}

	encryptedIconName, err := storage.encObject.Encrypt(addItemParams.Icon)
	if err != nil {
		return response, err
	}

	err = storage.dbObject.StartTX()
	if err != nil {
		return response, err
	}

	response.ItemID, err = storage.dbObject.dbInsertItem(encryptedItemName, encryptedIconName)
	if err != nil {
		errEndTX := storage.dbObject.RollbackTX()
		if errEndTX != nil {
			return response, formError(BSERR00006DbInsertFailed, err.Error(), errEndTX.Error())
		}
		return response, err
	}

	err = storage.dbObject.CommitTX()
	if err != nil {
		return response, err
	}

	response.Status = ConstSuccessResponse

	return response, nil
}

// ReadAllItems - read all not deleted items from the database and decrypt them
func (storage *StorageSingleton) ReadAllItems(readTags bool, readDeleted bool) (items []OxiItem, err error) {
	err = storage.checkReadiness()
	if err != nil {
		return items, err
	}

	items, err = storage.dbObject.dbSelectAllItems(readDeleted)
	if err != nil {
		return items, err
	}

	for i := range items {
		items[i].Name, err = storage.encObject.Decrypt(items[i].Name)
		if err != nil {
			return items, err
		}
		items[i].Icon, err = storage.encObject.Decrypt(items[i].Icon)
		if err != nil {
			return items, err
		}
		if readTags {
			items[i].Tags, err = storage.ReadTagsByItemID(items[i].ID)
			if err != nil {
				return items, err
			}
		}
	}

	return items, nil
}

// ReadItemById - read item by its Id
func (storage *StorageSingleton) ReadItemById(itemId int64, withDeleted bool) (item OxiItem, err error) {
	err = storage.checkReadiness()
	if err != nil {
		return item, err
	}

	item, err = storage.dbObject.dbGetItemById(itemId, withDeleted)
	if err != nil {
		return item, err
	}

	item.Name, err = storage.encObject.Decrypt(item.Name)
	if err != nil {
		return item, err
	}
	item.Icon, err = storage.encObject.Decrypt(item.Icon)
	if err != nil {
		return item, err
	}

	item.Fields, err = storage.ReadFieldsByItemID(itemId)
	if err != nil {
		return item, err
	}

	return item, nil
}
