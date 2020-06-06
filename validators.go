package bslib

import "errors"

func ValidateField(field BSField) error {
	if !CheckIfExistsFontAwesome(field.Icon) {
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

func ValidateItemBeforeUpdate(updateItemParams UpdateItemForm) error {
	if updateItemParams.ID <= 0 {
		return formError(BSERR00025ItemIdEmptyOrWrong)
	}

	if updateItemParams.Name == "" && updateItemParams.Icon == "" {
		return formError(BSERR00023UpdateFieldsEmpty)
	}
	if updateItemParams.Icon != "" {
		if !CheckIfExistsFontAwesome(updateItemParams.Icon) {
			return formError(BSERR00024FontAwesomeIconNotFound)
		}
	}
	return nil
}
