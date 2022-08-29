package validators

import (
	"github.com/oxipass/oxilib/assets"
	"github.com/oxipass/oxilib/internal/pkg/oxierr"
	"github.com/oxipass/oxilib/models"
)

func ValidateItemBeforeUpdate(updateItemParams models.UpdateItemForm) error {
	if updateItemParams.ID <= 0 {
		return oxierr.FormError(oxierr.BSERR00025ItemIdEmptyOrWrong)
	}

	if updateItemParams.Name == "" && updateItemParams.Icon == "" { // if both are empty then there is nothing to update
		return oxierr.FormError(oxierr.BSERR00027ItemValidationError)
	}
	if updateItemParams.Icon != "" {
		if !assets.CheckIfExistsInFontAwesome(updateItemParams.Icon) {
			return oxierr.FormError(oxierr.BSERR00024FontAwesomeIconNotFound)
		}
	}
	return nil
}
