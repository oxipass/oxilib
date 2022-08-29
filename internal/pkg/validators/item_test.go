package validators

import (
	"github.com/oxipass/oxilib/models"
	"testing"
)

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
