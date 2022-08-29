package assets

import (
	"errors"
	"github.com/oxipass/oxilib/consts"
	"strings"
	"testing"
)

const cTestSVGsolidKiss = "solid/face-kiss-wink-heart"
const cTestSVGregularClock = "regular/clock"
const cTestSVGbrandGolang = "brands/golang"
const cTestNotExistingIcon = "not-existing-icon"
const cSVGPrefix = "<svg xmlns="

func TestLoadingDefaultIcon(t *testing.T) {
	icon, err := getSVGImageFromFilePath(consts.CIconDefaultItem)
	if err != nil {
		t.Error(err)
	}
	if icon == "" || !strings.HasPrefix(icon, cSVGPrefix) {
		t.Error("icon is empty or is not svg")
	}
}

func TestLoadingSolidIcon(t *testing.T) {
	icon, err := getSVGImageFromFilePath(cTestSVGsolidKiss)
	if err != nil {
		t.Error(err)
	}
	if icon == "" || !strings.HasPrefix(icon, cSVGPrefix) {
		t.Error("icon is empty or is not svg")
	}
}

func TestLoadingRegularIcon(t *testing.T) {
	icon, err := getSVGImageFromFilePath(cTestSVGregularClock)
	if err != nil {
		t.Error(err)
	}
	if icon == "" || !strings.HasPrefix(icon, cSVGPrefix) {
		t.Error("icon is empty or is not svg")
	}
}

func TestLoadingBrandIcon(t *testing.T) {
	icon, err := getSVGImageFromFilePath(cTestSVGbrandGolang)
	if err != nil {
		t.Error(err)
	}
	if icon == "" || !strings.HasPrefix(icon, cSVGPrefix) {
		t.Error("icon is empty or is not svg")
	}
}

func TestLoadingNotExistingIcon(t *testing.T) {
	icon, err := getSVGImageFromFilePath(cTestNotExistingIcon)
	if err != nil {
		t.Error(err)
	}
	if icon == "" || !strings.HasPrefix(icon, cSVGPrefix) {
		t.Error("icon is empty or is not svg")
	}
	// Icon should be stanard document by default
	iconDoc, errDoc := getSVGImageFromFilePath(consts.CIconDefaultItem)
	if errDoc != nil {
		t.Error(errDoc)
	}

	if icon != iconDoc {
		t.Error("icon is not standard document")
	}
}

const cExistingSearchTerm = "email"
const cKeyEnvelope = "envelope"

func TestFontAwesomeSearch(t *testing.T) {

	foundKeys := SearchFontAwesomeList(cExistingSearchTerm)
	if len(foundKeys) == 0 {
		t.Error(errors.New("search is not working properly, predefined value not found"))
		t.FailNow()
		return
	}

	for _, foundKey := range foundKeys {
		if foundKey == cKeyEnvelope {
			return
		}
	}
	t.Error(errors.New("search is not working properly, predefined value not found"))
	t.FailNow()
}

func TestCheckIfExistsInFontAwesome(t *testing.T) {

	if CheckIfExistsInFontAwesome(cTestSVGsolidKiss) == false {
		t.Error(errors.New("existing value is not found"))
		t.FailNow()
	}

	if CheckIfExistsInFontAwesome(cTestNotExistingIcon) == true {
		t.Error(errors.New("non-existing value is found"))
		t.FailNow()
	}

}
