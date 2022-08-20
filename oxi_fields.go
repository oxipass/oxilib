package oxilib

import "github.com/oxipass/oxilib/models"

// DeleteField - delete existing field
func (storage *StorageSingleton) DeleteField(deleteFieldForm models.UpdateFieldForm) (response models.CommonResponse, err error) {
	err = storage.checkReadiness()
	if err != nil {
		return response, err
	}

	err = storage.dbObject.StartTX()
	if err != nil {
		return response, err
	}
	err = storage.dbObject.dbDeleteField(deleteFieldForm.ID)
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

// AddNewItem - adds new item
func (storage *StorageSingleton) AddNewField(addFieldForm models.UpdateFieldForm) (response models.FieldAddedResponse, err error) {
	var field models.OxiField

	err = storage.checkReadiness()
	if err != nil {
		return response, err
	}
	if err := ValidateField(addFieldForm.OxiField); err != nil {
		return response, formError(BSERR00006DbInsertFailed, err.Error())

	}

	field.Name, err = storage.encObject.Encrypt(addFieldForm.Name)
	if err != nil {
		return response, formError(BSERR00006DbInsertFailed, err.Error())
	}

	field.Icon, err = storage.encObject.Encrypt(addFieldForm.Icon)
	if err != nil {
		return response, formError(BSERR00006DbInsertFailed, err.Error())
	}

	field.ValueType, err = storage.encObject.Encrypt(addFieldForm.ValueType)
	if err != nil {
		return response, formError(BSERR00006DbInsertFailed, err.Error())
	}

	field.Value, err = storage.encObject.Encrypt(addFieldForm.Value)
	if err != nil {
		return response, formError(BSERR00006DbInsertFailed, err.Error())
	}

	err = storage.dbObject.StartTX()
	if err != nil {
		return response, err
	}

	fieldId, err := storage.dbObject.dbInsertField(addFieldForm.ItemID, field)
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
	response.FieldID = fieldId

	return response, nil
}

// ReadFieldsByItemID - real all the fields by ItemId
func (storage *StorageSingleton) ReadFieldsByItemID(itemId int64) (fields []models.OxiField, err error) {
	fieldsEncrypted, err := storage.dbObject.dbSelectAllItemFields(itemId)
	if err != nil {
		return fields, err
	}

	for _, field := range fieldsEncrypted {
		fieldReady, err := storage.DecryptField(field)
		if err != nil {
			return fields, err
		}
		fields = append(fields, fieldReady)
	}

	return fields, nil
}

// AddNewItem - adds new item
func (storage *StorageSingleton) ReadFieldsByFieldID(fieldId int64) (field models.OxiField, err error) {
	fieldEncrypted, err := storage.dbObject.dbGetFieldById(fieldId)
	if err != nil {
		return field, err
	}

	fieldReady, err := storage.DecryptField(fieldEncrypted)
	if err != nil {
		return field, err
	}
	return fieldReady, nil
}

func (storage *StorageSingleton) DecryptField(fieldEncrypted models.OxiField) (field models.OxiField, err error) {
	field.Value, err = storage.encObject.Decrypt(fieldEncrypted.Value)
	if err != nil {
		return field, err
	}
	field.Name, err = storage.encObject.Decrypt(fieldEncrypted.Name)
	if err != nil {
		return field, err
	}
	field.ValueType, err = storage.encObject.Decrypt(fieldEncrypted.ValueType)
	if err != nil {
		return field, err
	}
	field.Icon, err = storage.encObject.Decrypt(fieldEncrypted.Icon)
	if err != nil {
		return field, err
	}
	field.ID = fieldEncrypted.ID
	field.Created = fieldEncrypted.Created
	field.Updated = fieldEncrypted.Updated
	field.Deleted = fieldEncrypted.Deleted
	return field, nil
}
