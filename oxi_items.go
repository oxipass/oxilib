package oxilib

import (
	"github.com/oxipass/oxilib/assets"
	"github.com/oxipass/oxilib/consts"
	"github.com/oxipass/oxilib/internal/pkg/oxierr"
	"github.com/oxipass/oxilib/internal/pkg/validators"
	"github.com/oxipass/oxilib/models"
)

func (storage *StorageSingleton) checkReadiness() error {
	if storage.encObject == nil || !storage.encObject.IsReady() {
		return oxierr.FormError("Encryptor is not initialized", "checkReadiness")
	}

	if storage.dbObject == nil {
		return oxierr.FormError("Database is not initialized", "checkReadiness")
	}

	return nil
}

func (storage *StorageSingleton) DeleteItem(deleteItemParams models.UpdateItemForm) (response models.CommonResponse, err error) {
	err = storage.checkReadiness()
	if err != nil {
		return response, err
	}
	err = storage.dbObject.StartTX()
	if err != nil {
		return response, err
	}

	err = storage.dbObject.DbDeleteItem(deleteItemParams.ID)
	if err != nil {
		errEndTX := storage.dbObject.RollbackTX()
		if errEndTX != nil {
			return response, oxierr.FormError(oxierr.BSERR00016DbDeleteFailed, err.Error(), errEndTX.Error())
		}
		return response, err
	}

	err = storage.dbObject.CommitTX()
	if err != nil {
		return response, err
	}

	response.Status = consts.CSuccessResponse

	return response, nil
}

func (storage *StorageSingleton) UpdateItem(updateItemParams models.UpdateItemForm) (response models.ItemUpdatedResponse, err error) {
	var encryptedItemName, encryptedIconName string
	err = storage.checkReadiness()
	if err != nil {
		return response, err
	}

	err = validators.ValidateItemBeforeUpdate(updateItemParams)
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
		err = storage.dbObject.DbUpdateItemName(updateItemParams.ID, encryptedItemName)
		if err != nil {
			errEndTX := storage.dbObject.RollbackTX()
			if errEndTX != nil {
				return response, oxierr.FormError(oxierr.BSERR00018DbItemNameUpdateFailed, err.Error(), errEndTX.Error())
			}
			return response, err
		}
	}

	if updateItemParams.Icon != "" {
		err = storage.dbObject.DbUpdateItemIcon(updateItemParams.ID, encryptedIconName)
		if err != nil {
			errEndTX := storage.dbObject.RollbackTX()
			if errEndTX != nil {
				return response, oxierr.FormError(oxierr.BSERR00026DbItemIconUpdateFailed, err.Error(), errEndTX.Error())
			}
			return response, err
		}
	}

	err = storage.dbObject.CommitTX()
	if err != nil {
		return response, err
	}

	response.Status = consts.CSuccessResponse

	return response, nil
}

// AddNewItem - adds new item
func (storage *StorageSingleton) AddNewItem(addItemParams models.UpdateItemForm) (response models.ItemAddedResponse, err error) {

	err = storage.checkReadiness()
	if err != nil {
		return response, err
	}

	if !assets.CheckIfExistsInFontAwesome(addItemParams.Icon) {
		return response, oxierr.FormError(oxierr.BSERR00024FontAwesomeIconNotFound)
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

	response.ItemID, err = storage.dbObject.DbInsertItem(encryptedItemName, encryptedIconName)
	if err != nil {
		errEndTX := storage.dbObject.RollbackTX()
		if errEndTX != nil {
			return response, oxierr.FormError(oxierr.BSERR00006DbInsertFailed, err.Error(), errEndTX.Error())
		}
		return response, err
	}

	err = storage.dbObject.CommitTX()
	if err != nil {
		return response, err
	}

	response.Status = consts.CSuccessResponse

	return response, nil
}

// ReadAllItems TODO: Sort items before returning them
// ReadAllItems - read all not deleted items from the db and decrypt them
func (storage *StorageSingleton) ReadAllItems(readTags bool, readDeleted bool) (items []models.OxiItem, err error) {
	err = storage.checkReadiness()
	if err != nil {
		return items, err
	}

	items, err = storage.dbObject.DbSelectAllItems(readDeleted)
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
func (storage *StorageSingleton) ReadItemById(itemId int64, withDeleted bool) (item models.OxiItem, err error) {
	err = storage.checkReadiness()
	if err != nil {
		return item, err
	}

	item, err = storage.dbObject.DbGetItemById(itemId, withDeleted)
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
