package bslib

import "errors"

func ValidateField(field BSField) error {
	if CheckIfExistsFontAwesome(field.Icon) == false {
		return errors.New(BSERR00022ValidationFailed + ": icon not found")
	}
	if CheckValueType(field.ValueType) == false {
		return errors.New(BSERR00022ValidationFailed + ": value type not found")
	}
	if field.Name == "" {
		return errors.New(BSERR00022ValidationFailed + ": field name cannot be empty")
	}

	return nil
}