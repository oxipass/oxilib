package bslib

func (storage *StorageSingleton) checkReadiness() error {
	if storage.encObject == nil || !storage.encObject.IsReady() {
		return formError("Encrypter is not initialized", "checkReadiness")
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

func (storage *StorageSingleton) UpdateItem(updateItemParms UpdateItemForm) (response ItemUpdatedResponse, err error) {
	var encryptedItemName, encryptedIconName string
	err = storage.checkReadiness()
	if err != nil {
		return response, err
	}

	if updateItemParms.ID <= 0 {
		return response, formError(BSERR00025ItemIdEmptyOrWrong)
	}

	if updateItemParms.Name == "" && updateItemParms.Icon == "" {
		return response, formError(BSERR00023UpdateFieldsEmpty)
	}

	if updateItemParms.Icon != "" {
		if !CheckIfExistsFontAwesome(updateItemParms.Icon) {
			return response, formError(BSERR00024FontAwesomeIconNotFound)
		}
		encryptedIconName, err = storage.encObject.Encrypt(updateItemParms.Icon)
		if err != nil {
			return response, err
		}
	}
	if updateItemParms.Name != "" {
		encryptedItemName, err = storage.encObject.Encrypt(updateItemParms.Name)
		if err != nil {
			return response, err
		}
	}

	err = storage.dbObject.StartTX()
	if err != nil {
		return response, err
	}

	if updateItemParms.Name != "" {
		err = storage.dbObject.dbUpdateItemName(updateItemParms.ID, encryptedItemName)
		if err != nil {
			errEndTX := storage.dbObject.RollbackTX()
			if errEndTX != nil {
				return response, formError(BSERR00018DbItemNameUpdateFailed, err.Error(), errEndTX.Error())
			}
			return response, err
		}
	}

	if updateItemParms.Icon != "" {
		err = storage.dbObject.dbUpdateItemIcon(updateItemParms.ID, encryptedIconName)
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
func (storage *StorageSingleton) AddNewItem(addItemParms UpdateItemForm) (response ItemAddedResponse, err error) {

	err = storage.checkReadiness()
	if err != nil {
		return response, err
	}

	if !CheckIfExistsFontAwesome(addItemParms.Icon) {
		return response, formError(BSERR00024FontAwesomeIconNotFound)
	}

	encryptedItemName, err := storage.encObject.Encrypt(addItemParms.Name)
	if err != nil {
		return response, err
	}

	encryptedIconName, err := storage.encObject.Encrypt(addItemParms.Icon)
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
func (storage *StorageSingleton) ReadAllItems(readDeleted bool) (items []BSItem, err error) {
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
	}

	return items, nil
}

// ReadAllItems - read all not deleted items from the database and decrypt them
func (storage *StorageSingleton) ReadItemById(itemId int64) (item BSItem, err error) {
	err = storage.checkReadiness()
	if err != nil {
		return item, err
	}

	item, err = storage.dbObject.dbGetItemById(itemId)
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
