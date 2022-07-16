package oxilib

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestLangsContent(t *testing.T) {
	files, errDir := getLangsFiles()
	if errDir != nil {
		t.Error(errDir)
		t.FailNow()
		return
	}

	errHappened := false
	for _, file := range files {
		var translations Translations
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
