package oxilib

import (
	"errors"
	"github.com/oxipass/oxilib/models"
)

func ValidateField(field models.OxiField) error {
	if !CheckIfExistsInFontAwesome(field.Icon) {
		return errors.New(BSERR00022ValidationFailed + ": icon not found")
	}
	if !CheckValueType(field.ValueType) {
		return errors.New(BSERR00022ValidationFailed + ": value type not found")
	}
	if field.Name == "" {
		return errors.New(BSERR00022ValidationFailed + ": field name cannot be empty")
	}

	return nil
}

func ValidateItemBeforeUpdate(updateItemParams models.UpdateItemForm) error {
	if updateItemParams.ID <= 0 {
		return formError(BSERR00025ItemIdEmptyOrWrong)
	}

	if updateItemParams.Name == "" && updateItemParams.Icon == "" { // if both are empty then there is nothing to update
		return formError(BSERR00027ItemValidationError)
	}
	if updateItemParams.Icon != "" {
		if !CheckIfExistsInFontAwesome(updateItemParams.Icon) {
			return formError(BSERR00024FontAwesomeIconNotFound)
		}
	}
	return nil
}
