package oxilib

import (
	"github.com/oxipass/oxilib/assets"
	"testing"
)

func TestDefaultItemsAvailble(t *testing.T) {
	templatesFromAssets, err := assets.GetItemsTemplate()
	if err != nil {
		t.Error(err)
	}
	defItemsLength := len(templatesFromAssets.Items)
	storage := GetInstance()
	templatesFromDb, err := storage.GetTemplatesItems()
	if err != nil {
		t.Error(err)
	}
	itemsMatches := 0
	for _, assetsItem := range templatesFromAssets.Items {
		matchFound := false
		for _, templItem := range templatesFromDb {
			if assetsItem.ID == templItem.ID {
				itemsMatches++
				matchFound = true
				break
			}
		}
		if !matchFound {
			t.Error("Asset item '" + assetsItem.ID + "' is not found in DB")
		}
	}
	if itemsMatches != defItemsLength {
		t.Error("Templates count in Assets and DB is different")
	}
}
