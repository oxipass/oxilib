package oxilib

import (
	"github.com/oxipass/oxilib/assets"
	"github.com/oxipass/oxilib/models"
)

// TODO: Get items templates
// TODO: Store the item as template
// TODO: Get fields templates

func (storage *StorageSingleton) SaveItemAsTemplate(item models.OxiItem) (err error) {

	return nil
}

func (storage *StorageSingleton) SaveItemTemplateAsItem(item models.OxiItemTemplate) (err error) {

	return nil
}

func (storage *StorageSingleton) AddItemTemplate(itemTemplate models.ItemTemplateJSON) error {
	err := storage.dbObject.DbInsertItemTemplate(itemTemplate.ID,
		storage.T(itemTemplate.ID),
		itemTemplate.Icon)
	if err != nil {
		return err
	}
	ft, errF := assets.GetFieldsTemplate()
	if errF != nil {
		return errF
	}

	for _, fieldId := range itemTemplate.FieldsIds {
		var oxiField models.OxiField
		foundTemplate := false
		for _, fTemplate := range ft.Fields {
			if fTemplate.ID == fieldId {
				foundTemplate = true
				oxiField.Name = storage.T(fTemplate.ID) // Retrieve translation for field name
				oxiField.ValueType = fTemplate.FieldType
				oxiField.Icon = fTemplate.Icon
				break
			}
		}
		if foundTemplate {
			errField := storage.dbObject.DbInsertFieldTemplate(itemTemplate.ID, fieldId, oxiField)
			if errField != nil {
				return errField
			}
		}
	}
	return nil
}

func (storage *StorageSingleton) AddDefaultItemsTemplates() error {
	items, err := assets.GetItemsTemplate()
	if err != nil {
		return err
	}

	for _, item := range items.Items {
		err = storage.AddItemTemplate(item)
		if err != nil {
			return err
		}
	}

	return nil
}
