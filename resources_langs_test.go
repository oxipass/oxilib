package oxilib

import (
	"encoding/json"
	"github.com/oxipass/oxilib/models"
	"strings"
	"testing"
)

func TestWrongLanguage(t *testing.T) {
	lang, err := initLang("wrong")
	if err == nil && lang != "en" {
		t.Error("English should be loaded instead of wrong language")
		t.FailNow()
	}
}

func TestSetWrongLanguage(t *testing.T) {
	err := setLang("wrong")
	if err != nil {
		t.Error("English should be loaded instead of wrong language, no error expected")
		t.FailNow()
	}
}

func TestInitLang(t *testing.T) {
	langs := getLangs()
	for _, lang := range langs {
		setLang, errLang := initLang(lang.Code)
		if errLang != nil || setLang != lang.Code {
			if errLang == nil {
				t.Error(errLang)
			} else {
				t.Error("error while init lang: " + lang.Code)
			}
		}
	}
}

func TestSetLang(t *testing.T) {
	langs := getLangs()
	for _, lang := range langs {
		errLang := setLang(lang.Code)
		if errLang != nil {
			t.Error(errLang)
		}
	}
}

func TestGetLangs(t *testing.T) {
	langs := getLangs()

	errHappened := false
	for _, lang := range langs {

		if lang.Code == "" {
			t.Error("lang code is empty for " + lang.Name)
			errHappened = true
			continue
		}
		if lang.Name == "" {
			t.Error("lang name is empty for " + lang.Code)
			errHappened = true
			continue
		}
		if lang.NativeName == "" {
			t.Error("lang native name is empty for " + lang.Code)
			errHappened = true
			continue
		}
		if lang.Locales == nil {
			t.Error("lang locales is empty for " + lang.Code)
			errHappened = true
			continue
		}
	}
	if errHappened {
		t.FailNow()
	}
}

func TestLangsContent(t *testing.T) {
	files, errDir := getLangsFiles()
	if errDir != nil {
		t.Error(errDir)
		t.FailNow()
		return
	}

	errHappened := false
	for _, file := range files {
		var translations models.Translations
		fileBytes, errFile := langsResources.ReadFile(cLangsFolder + "/" + file.Name())
		if errFile != nil {
			t.Error(errFile)
			errHappened = true
			continue
		}
		errUnmarshal := json.Unmarshal(fileBytes, &translations)
		if errUnmarshal != nil {
			t.Error(errUnmarshal, "for file: "+file.Name()+" and lang "+translations.Name)
			errHappened = true
			continue
		}
	}
	if errHappened {
		t.FailNow()
	}
}

func TestLangsAvailability(t *testing.T) {
	files, errDir := getLangsFiles()
	if errDir != nil {
		t.Error(errDir)
		t.FailNow()
		return
	}
	errHappened := false
	for _, file := range files {
		if file.IsDir() {
			t.Error("langs directory contains directory: " + file.Name())
			errHappened = true
			continue
		}
		if !strings.HasSuffix(file.Name(), ".json") {
			t.Error("langs directory contains file with wrong extension: " + file.Name())
			errHappened = true
			continue
		}

	}
	if errHappened {
		t.FailNow()
	}

}

// TestLangsHaveEngKeys tests that all langs have keys as en.json
func TestLangsHaveEngKeys(t *testing.T) {
	files, errDir := getLangsFiles()
	if errDir != nil {
		t.Error(errDir)
		t.FailNow()
		return
	}

	var engTranslations models.Translations
	engBytes, errEngFile := langsResources.ReadFile(cLangsFolder + "/" + "en.json")
	if errEngFile != nil {
		t.Error(errEngFile)
		t.FailNow()
		return
	}
	errUnmarshalEng := json.Unmarshal(engBytes, &engTranslations)
	if errUnmarshalEng != nil {
		t.Error(errUnmarshalEng, "for file:en.json and lang "+engTranslations.Name)
		t.FailNow()
		return
	}

	errHappened := false
	for _, file := range files {
		var translations models.Translations
		fileBytes, errFile := langsResources.ReadFile(cLangsFolder + "/" + file.Name())
		if errFile != nil {
			t.Error(errFile)
			errHappened = true
			continue
		}
		errUnmarshal := json.Unmarshal(fileBytes, &translations)
		if errUnmarshal != nil {
			t.Error(errUnmarshal, "for file: "+file.Name()+" and lang "+translations.Name)
			errHappened = true
			continue
		}
		for tagKey, _ := range engTranslations.Translations {
			if transValue, ok := translations.Translations[tagKey]; !ok {
				t.Error("key " + tagKey + " is missing in " + file.Name())
				errHappened = true
				continue
			} else if transValue == "" {
				t.Error("key " + tagKey + " is empty in " + file.Name())
				errHappened = true
				continue
			}
		}
	}
	if errHappened {
		t.FailNow()
	}
}

// TestLangsKeysHaveEngKeys tests that all keys in langs have keys in en.json
func TestLangsKeysHaveEngKey(t *testing.T) {
	files, errDir := getLangsFiles()
	if errDir != nil {
		t.Error(errDir)
		t.FailNow()
		return
	}

	var engTranslations models.Translations
	engBytes, errEngFile := langsResources.ReadFile(cLangsFolder + "/" + "en.json")
	if errEngFile != nil {
		t.Error(errEngFile)
		t.FailNow()
		return
	}
	errUnmarshalEng := json.Unmarshal(engBytes, &engTranslations)
	if errUnmarshalEng != nil {
		t.Error(errUnmarshalEng, "for file:en.json and lang "+engTranslations.Name)
		t.FailNow()
		return
	}

	errHappened := false
	for _, file := range files {
		var translations models.Translations
		fileBytes, errFile := langsResources.ReadFile(cLangsFolder + "/" + file.Name())
		if errFile != nil {
			t.Error(errFile)
			errHappened = true
			continue
		}
		errUnmarshal := json.Unmarshal(fileBytes, &translations)
		if errUnmarshal != nil {
			t.Error(errUnmarshal, "for file: "+file.Name()+" and lang "+translations.Name)
			errHappened = true
			continue
		}
		for tagKey, _ := range translations.Translations {
			if _, ok := engTranslations.Translations[tagKey]; !ok {
				t.Error("key '" + tagKey + "' in file '" + file.Name() + "' is missing in en.json")
				errHappened = true
				continue
			}
		}
	}
	if errHappened {
		t.FailNow()
	}
}
