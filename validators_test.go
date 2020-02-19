package bslib

import "testing"

func TestFieldValidator(t *testing.T) {
	field := BSField{Name: "my field", ValueType: "text", Icon: "fab fa-android"}
	err := ValidateField(field)
	if err != nil {
		t.Errorf("Expected no error, retrived: %s", err.Error())
	}
}

func TestFieldValidatorEmptyName(t *testing.T) {
	field := BSField{ValueType: "text", Icon: "fab fa-android"}
	err := ValidateField(field)
	if err == nil {
		t.Errorf("Expected validation error because of empty field name")
	}
}

func TestFieldValidatorWrongValueType(t *testing.T) {
	field := BSField{Name: "my field", ValueType: "sdgshdgh", Icon: "fab fa-android"}
	err := ValidateField(field)
	if err == nil {
		t.Errorf("Expected validation error because of not supported value type")
	}
}

func TestFieldValidatorNotExistingIcon(t *testing.T) {
	field := BSField{Name: "my field", ValueType: "text", Icon: "fab fa-blah-blah"}
	err := ValidateField(field)
	if err == nil {
		t.Errorf("Expected error because of non existing icon name ")
	}
}
