package bslib

import "sort"

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

	encryptedTag, errEnc := storage.encObject.Encrypt(addTagParam.Name)
	if errEnc != nil {
		return response, formError(BSERR00006DbInsertFailed, errEnc.Error())
	}

	response.TagId, err = storage.dbObject.dbInsertTag(encryptedTag)
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
func (storage *StorageSingleton) ReadTagsByItemID(itemId int64) (tags []BSTag, err error) {
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
func (storage *StorageSingleton) GetTags() (tags []BSTag, err error) {
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

	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Name < tags[j].Name
	}) // Sort by name before returning

	return tags, nil
}

func (storage *StorageSingleton) DecryptTag(tag BSTag) (decryptedTag BSTag, err error) {
	decryptedTag = tag
	decryptedTag.Name, err = storage.encObject.Decrypt(tag.Name)
	if err != nil {
		return decryptedTag, err
	}
	return decryptedTag, err
}
