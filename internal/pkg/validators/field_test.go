package validators

import (
	"github.com/oxipass/oxilib/consts"
	"github.com/oxipass/oxilib/models"
	"testing"
)

const cValidateFieldIcon01 = "brands/android"
const cValidateFieldIcon02wrong = "ekjdnwdednkjwndkjw"
const cValidateFieldName01 = "my new field"
const cValidateFieldType01wrong = "kjenwjdnwkjdnwk"

func TestFieldValidator(t *testing.T) {
	field := models.OxiField{Name: cValidateFieldName01, ValueType: consts.VTText, Icon: cValidateFieldIcon01}
	err := ValidateField(field)
	if err != nil {
		t.Errorf("Expected no error, retrived: %s, field icon: %s", err.Error(), cValidateFieldIcon01)
	}
}

func TestFieldValidatorEmptyName(t *testing.T) {
	field := models.OxiField{ValueType: consts.VTText, Icon: cValidateFieldIcon01}
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
	field := models.OxiField{Name: cValidateFieldName01, ValueType: consts.VTText, Icon: cValidateFieldIcon02wrong}
	err := ValidateField(field)
	if err == nil {
		t.Errorf("Expected error because of non existing icon name ")
	}
}
