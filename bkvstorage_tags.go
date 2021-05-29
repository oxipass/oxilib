package bslib

import "sort"

// AssignTag - assign a tag to item
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

// AddNewTag - adds new tag
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

// ReadTagsByItemID - real all the tags by ItemId
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

// GetTags - real all the available tags
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

// DecryptTag - decrypt the tag
func (storage *StorageSingleton) DecryptTag(tag BSTag) (decryptedTag BSTag, err error) {
	decryptedTag = tag
	decryptedTag.Name, err = storage.encObject.Decrypt(tag.Name)
	if err != nil {
		return decryptedTag, err
	}
	return decryptedTag, err
}
