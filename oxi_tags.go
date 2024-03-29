package oxilib

import (
	"github.com/oxipass/oxilib/assets"
	"github.com/oxipass/oxilib/consts"
	"github.com/oxipass/oxilib/internal/pkg/oxierr"
	"github.com/oxipass/oxilib/models"
)

func (storage *StorageSingleton) AssignTag(updateTagForm models.UpdateTagForm) (response models.TagAssignedResponse, err error) {
	err = storage.checkReadiness()
	if err != nil {
		return response, err
	}

	err = storage.dbObject.StartTX()
	if err != nil {
		return response, oxierr.FormError(oxierr.BSERR00006DbInsertFailed, err.Error())
	}

	response.ItemTagId, err = storage.dbObject.DbAssignTag(updateTagForm.ID, updateTagForm.ItemID)
	if err != nil {
		errEndTX := storage.dbObject.RollbackTX()
		if errEndTX != nil {
			return response, oxierr.FormError(oxierr.BSERR00006DbInsertFailed, err.Error(), errEndTX.Error())
		}
		return response, oxierr.FormError(oxierr.BSERR00006DbInsertFailed, err.Error())
	}

	err = storage.dbObject.CommitTX()
	if err != nil {
		return response, oxierr.FormError(oxierr.BSERR00006DbInsertFailed, err.Error())
	}

	response.Status = consts.CSuccessResponse

	return response, nil
}

// AddNewItem - adds new item
func (storage *StorageSingleton) AddNewTag(addTagParam models.UpdateTagForm) (response models.TagAddedResponse, err error) {

	err = storage.checkReadiness()
	if err != nil {
		return response, err
	}

	err = storage.dbObject.StartTX()
	if err != nil {
		return response, oxierr.FormError(oxierr.BSERR00006DbInsertFailed, err.Error())
	}

	encryptedTag, errEncT := storage.encObject.Encrypt(addTagParam.Name)
	if errEncT != nil {
		return response, oxierr.FormError(oxierr.BSERR00006DbInsertFailed, errEncT.Error())
	}

	encryptedColor, errEncC := storage.encObject.Encrypt(addTagParam.Color)
	if errEncC != nil {
		return response, oxierr.FormError(oxierr.BSERR00006DbInsertFailed, errEncC.Error())
	}

	encryptedExtId, errEncE := storage.encObject.Encrypt(addTagParam.ExtId)
	if errEncE != nil {
		return response, oxierr.FormError(oxierr.BSERR00006DbInsertFailed, errEncE.Error())
	}

	response.TagId, err = storage.dbObject.DbInsertTag(encryptedTag, encryptedColor, encryptedExtId)
	if err != nil {
		errEndTX := storage.dbObject.RollbackTX()
		if errEndTX != nil {
			return response, oxierr.FormError(oxierr.BSERR00006DbInsertFailed, err.Error(), errEndTX.Error())
		}
		return response, oxierr.FormError(oxierr.BSERR00006DbInsertFailed, err.Error())
	}

	err = storage.dbObject.CommitTX()
	if err != nil {
		return response, oxierr.FormError(oxierr.BSERR00006DbInsertFailed, err.Error())
	}

	response.Status = consts.CSuccessResponse

	return response, nil
}

// ReadFieldsByItemID - real all the fields by ItemId
func (storage *StorageSingleton) ReadTagsByItemID(itemId int64) (tags []models.OxiTag, err error) {
	fieldsEncrypted, err := storage.dbObject.DbSelectItemTags(itemId)
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
func (storage *StorageSingleton) GetTags() (tags []models.OxiTag, err error) {
	fieldsEncrypted, err := storage.dbObject.DbSelectTags()
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

func (storage *StorageSingleton) DecryptTag(tag models.OxiTag) (decryptedTag models.OxiTag, err error) {
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

func (storage *StorageSingleton) AddDefaultTags() error {
	tags, err := assets.GetTagsTemplate()
	if err != nil {
		return err
	}

	for _, templateTag := range tags.Tags {
		var tag models.UpdateTagForm
		tag.Name = storage.T(templateTag.ID)
		tag.Color = templateTag.Color
		tag.ExtId = templateTag.ID

		_, err := storage.AddNewTag(tag)
		if err != nil {
			return err
		}
	}
	return nil
}
