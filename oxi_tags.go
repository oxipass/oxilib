package oxilib

func (storage *StorageSingleton) AssignTag(updateTagForm UpdateTagForm) (response TagAssignedResponse, err error) {
	err = storage.checkReadiness()
	if err != nil {
		return response, err
	}

	err = storage.dbObject.StartTX()
	if err != nil {
		return response, formError(BSERR00006DbInsertFailed, err.Error())
	}

	response.ItemTagId, err = storage.dbObject.dbAssignTag(updateTagForm.ID, updateTagForm.ItemID)
	if err != nil {
		errEndTX := storage.dbObject.RollbackTX()
		if errEndTX != nil {
			return response, formError(BSERR00006DbInsertFailed, err.Error(), errEndTX.Error())
		}
		return response, formError(BSERR00006DbInsertFailed, err.Error())
	}

	err = storage.dbObject.CommitTX()
	if err != nil {
		return response, formError(BSERR00006DbInsertFailed, err.Error())
	}

	response.Status = ConstSuccessResponse

	return response, nil
}

// AddNewItem - adds new item
func (storage *StorageSingleton) AddNewTag(addTagParam UpdateTagForm) (response TagAddedResponse, err error) {

	err = storage.checkReadiness()
	if err != nil {
		return response, err
	}

	err = storage.dbObject.StartTX()
	if err != nil {
		return response, formError(BSERR00006DbInsertFailed, err.Error())
	}

	encryptedTag, errEncT := storage.encObject.Encrypt(addTagParam.Name)
	if errEncT != nil {
		return response, formError(BSERR00006DbInsertFailed, errEncT.Error())
	}

	encryptedColor, errEncC := storage.encObject.Encrypt(addTagParam.Color)
	if errEncC != nil {
		return response, formError(BSERR00006DbInsertFailed, errEncC.Error())
	}

	encryptedExtId, errEncE := storage.encObject.Encrypt(addTagParam.ExtId)
	if errEncE != nil {
		return response, formError(BSERR00006DbInsertFailed, errEncE.Error())
	}

	response.TagId, err = storage.dbObject.dbInsertTag(encryptedTag, encryptedColor, encryptedExtId)
	if err != nil {
		errEndTX := storage.dbObject.RollbackTX()
		if errEndTX != nil {
			return response, formError(BSERR00006DbInsertFailed, err.Error(), errEndTX.Error())
		}
		return response, formError(BSERR00006DbInsertFailed, err.Error())
	}

	err = storage.dbObject.CommitTX()
	if err != nil {
		return response, formError(BSERR00006DbInsertFailed, err.Error())
	}

	response.Status = ConstSuccessResponse

	return response, nil
}

// ReadFieldsByItemID - real all the fields by ItemId
func (storage *StorageSingleton) ReadTagsByItemID(itemId int64) (tags []OxiTag, err error) {
	fieldsEncrypted, err := storage.dbObject.dbSelectItemTags(itemId)
	if err != nil {
		return tags, err
	}

	for _, field := range fieldsEncrypted {
		tagReady, err := storage.DecryptTag(field)
		if err != nil {
			return tags, err
		}
		tags = append(tags, tagReady)
	}

	return tags, nil
}

// ReadFieldsByItemID - real all the fields by ItemId
func (storage *StorageSingleton) GetTags() (tags []OxiTag, err error) {
	fieldsEncrypted, err := storage.dbObject.dbSelectTags()
	if err != nil {
		return tags, err
	}

	for _, field := range fieldsEncrypted {
		tagReady, err := storage.DecryptTag(field)
		if err != nil {
			return tags, err
		}
		tags = append(tags, tagReady)
	}

	return tags, nil
}

func (storage *StorageSingleton) DecryptTag(tag OxiTag) (decryptedTag OxiTag, err error) {
	decryptedTag = tag
	decryptedTag.Name, err = storage.encObject.Decrypt(tag.Name)
	if err != nil {
		return decryptedTag, err
	}
	decryptedTag.Color, err = storage.encObject.Decrypt(tag.Color)
	if err != nil {
		return decryptedTag, err
	}
	decryptedTag.ExtId, err = storage.encObject.Decrypt(tag.ExtId)
	if err != nil {
		return decryptedTag, err
	}
	return decryptedTag, err
}
