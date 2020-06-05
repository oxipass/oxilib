package bslib

import "testing"

const cValidateFieldIcon01 = "fab fa-android"
const cValidateFieldIcon02wrong = "fab ekjdnwdednkjwndkjw"
const cValidateFieldName01 = "my new field"
const cValidateFieldType01wrong = "kjenwjdnwkjdnwk"

func TestFieldValidator(t *testing.T) {
	field := BSField{Name: cValidateFieldName01, ValueType: VTText, Icon: cValidateFieldIcon01}
	err := ValidateField(field)
	if err != nil {
		t.Errorf("Expected no error, retrived: %s", err.Error())
	}
}

func TestFieldValidatorEmptyName(t *testing.T) {
	field := BSField{ValueType: VTText, Icon: cValidateFieldIcon01}
	err := ValidateField(field)
	if err == nil {
		t.Errorf("Expected validation error because of empty field name")
	}
}

func TestFieldValidatorWrongValueType(t *testing.T) {
	field := BSField{Name: cValidateFieldName01, ValueType: cValidateFieldType01wrong, Icon: cValidateFieldIcon01}
	err := ValidateField(field)
	if err == nil {
		t.Errorf("Expected validation error because of not supported value type")
	}
}

func TestFieldValidatorNotExistingIcon(t *testing.T) {
	field := BSField{Name: cValidateFieldName01, ValueType: VTText, Icon: cValidateFieldIcon02wrong}
	err := ValidateField(field)
	if err == nil {
		t.Errorf("Expected error because of non existing icon name ")
	}
}
