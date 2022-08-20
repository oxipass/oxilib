package oxilib

import (
	"github.com/oxipass/oxilib/models"
	"testing"
)

const cValidateFieldIcon01 = "brands/android"
const cValidateFieldIcon02wrong = "ekjdnwdednkjwndkjw"
const cValidateFieldName01 = "my new field"
const cValidateFieldType01wrong = "kjenwjdnwkjdnwk"

func TestFieldValidator(t *testing.T) {
	field := models.OxiField{Name: cValidateFieldName01, ValueType: VTText, Icon: cValidateFieldIcon01}
	err := ValidateField(field)
	if err != nil {
		t.Errorf("Expected no error, retrived: %s, field icon: %s", err.Error(), cValidateFieldIcon01)
	}
}

func TestFieldValidatorEmptyName(t *testing.T) {
	field := models.OxiField{ValueType: VTText, Icon: cValidateFieldIcon01}
	err := ValidateField(field)
	if err == nil {
		t.Errorf("Expected validation error because of empty field name")
	}
}

func TestFieldValidatorWrongValueType(t *testing.T) {
	field := models.OxiField{Name: cValidateFieldName01, ValueType: cValidateFieldType01wrong, Icon: cValidateFieldIcon01}
	err := ValidateField(field)
	if err == nil {
		t.Errorf("Expected validation error because of not supported value type")
	}
}

func TestFieldValidatorNotExistingIcon(t *testing.T) {
	field := models.OxiField{Name: cValidateFieldName01, ValueType: VTText, Icon: cValidateFieldIcon02wrong}
	err := ValidateField(field)
	if err == nil {
		t.Errorf("Expected error because of non existing icon name ")
	}
}

func TestValidateItemId(t *testing.T) {
	var itemForm models.UpdateItemForm
	itemForm.ID = 0
	err := ValidateItemBeforeUpdate(itemForm)
	if err == nil {
		t.Errorf("Expected error because of 0 id value ")
	}
	itemForm.ID = -1
	err = ValidateItemBeforeUpdate(itemForm)
	if err == nil {
		t.Errorf("Expected error because of negative id value ")
	}
}

func TestValidateItemName(t *testing.T) {
	var itemForm models.UpdateItemForm
	itemForm.ID = 1
	itemForm.Name = ""
	err := ValidateItemBeforeUpdate(itemForm)
	if err == nil {
		t.Errorf("Expected error because of empty name")
	}
}

const cFontAwesomeWrongSample = "fas fa-blah-blah"

func TestValidateItemIcon(t *testing.T) {
	var itemForm models.UpdateItemForm
	itemForm.ID = 1
	itemForm.Icon = ""
	err := ValidateItemBeforeUpdate(itemForm)
	if err == nil {
		t.Errorf("Expected error because of empty icon")
	}
	itemForm.Icon = cFontAwesomeWrongSample
	err = ValidateItemBeforeUpdate(itemForm)
	if err == nil {
		t.Errorf("Expected error because of non valid fant awesome value")
	}
}
