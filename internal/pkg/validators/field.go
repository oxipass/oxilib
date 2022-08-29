package validators

import (
	"errors"
	"github.com/oxipass/oxilib/assets"
	"github.com/oxipass/oxilib/consts"
	"github.com/oxipass/oxilib/internal/pkg/oxierr"
	"github.com/oxipass/oxilib/models"
)

func ValidateField(field models.OxiField) error {
	if !assets.CheckIfExistsInFontAwesome(field.Icon) {
		return errors.New(oxierr.BSERR00022ValidationFailed + ": icon not found")
	}
	if !consts.CheckValueType(field.ValueType) {
		return errors.New(oxierr.BSERR00022ValidationFailed + ": value type not found")
	}
	if field.Name == "" {
		return errors.New(oxierr.BSERR00022ValidationFailed + ": field name cannot be empty")
	}

	return nil
}
