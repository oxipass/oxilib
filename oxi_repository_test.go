package oxilib

import (
	"errors"
	"strings"
	"testing"
)

const cValueTypesCount = 13

func TestValueTypes(t *testing.T) {
	vTypes := GetValueTypes()
	if vTypes == nil {
		t.Error(errors.New("value types are empty"))
		t.FailNow()
		return
	}
	if len(vTypes) != cValueTypesCount {
		t.Error(errors.New("wrong value types count or unhandled value type is added"))
		t.FailNow()
		return
	}
}

func TestFontAwesomeValues(t *testing.T) {
	// Only free will be used (fas & fab)
	for _, faValue := range GetFontAwesomeList() {
		if !(strings.HasPrefix(faValue, "fas") || strings.HasPrefix(faValue, "fab")) {
			t.Error(errors.New("font awesome can have only prefixes 'fas' & 'fab', found prefix: " + faValue))
			t.FailNow()
			return
		}
	}
}

func TestFontAwesomeSearch(t *testing.T) {
	const searchTerm = "imdb"

	searchResult := SearchFontAwesomeList(searchTerm)
	if len(searchResult) == 0 {
		t.Error(errors.New("search is not working properly, predefined value not found"))
		t.FailNow()
		return
	}
}

const cTestRandomIconName = "kjnfejkfnekwj"

func TestCheckIfExistsFontAwesome(t *testing.T) {

	if CheckIfExistsFontAwesome("fas fa-battery-empty") == false {
		t.Error(errors.New("existing value is not found"))
		t.FailNow()
		return
	}

	if CheckIfExistsFontAwesome(cTestRandomIconName) == true {
		t.Error(errors.New("non-existing value is found"))
		t.FailNow()
		return
	}

	for _, faValue := range GetFontAwesomeList() {
		if CheckIfExistsFontAwesome(faValue) == false {
			t.Error(errors.New("existing value is not found: " + faValue))
			t.FailNow()
			return
		}
	}
}
