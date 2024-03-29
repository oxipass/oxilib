package assets

import (
	"embed"
	"encoding/json"
	"errors"
	"github.com/oxipass/oxilib/consts"
	"github.com/oxipass/oxilib/models"
	"io/fs"
)

var (
	//go:embed langs/*.json
	langsResources embed.FS
)

var tr map[string]string
var currentLanguage string

func getLangsFiles() ([]fs.DirEntry, error) {
	return langsResources.ReadDir(consts.CLangsFolder)
}

func initLang(langCode string) (string, error) {
	files, err := getLangsFiles()
	if err != nil {
		return "", err
	}
	for _, file := range files {
		var lang models.Lang
		fileBytes, errFile := langsResources.ReadFile(consts.CLangsFolder + "/" + file.Name())
		if errFile != nil {
			return "", errFile
		}
		errUnmarshal := json.Unmarshal(fileBytes, &lang)
		if errUnmarshal != nil {
			return "", errUnmarshal
		}
		if lang.Code == langCode {
			var translations models.Translations
			transBytes, errTrans := langsResources.ReadFile(consts.CLangsFolder + "/" + file.Name())
			if errTrans != nil {
				return "", errTrans
			}
			errUnmarshalTrans := json.Unmarshal(transBytes, &translations)
			if errUnmarshalTrans != nil {
				return "", errUnmarshalTrans
			}
			tr = translations.Translations
			return langCode, nil
		}
	}
	if langCode != "en" {
		return initLang("en")
	}
	return "", errors.New("language not loaded")
}

func T(key string) string {
	var err error
	if tr == nil {
		currentLanguage, err = initLang("en")
		if err != nil {
			return key
		}
	}
	return tr[key]
}

func SetLang(langCode string) error {
	var err error
	currentLanguage, err = initLang(langCode)
	if err != nil {
		return err
	}
	return nil
}

func GetLangs() []models.Lang {
	files, err := getLangsFiles()
	if err != nil {
		return nil
	}
	langs := make([]models.Lang, len(files))
	for i, file := range files {
		var lang models.Lang
		fileBytes, errFile := langsResources.ReadFile(consts.CLangsFolder + "/" + file.Name())
		if errFile != nil {
			return nil
		}
		errUnmarshal := json.Unmarshal(fileBytes, &lang)
		if errUnmarshal != nil {
			return nil
		}
		langs[i] = lang
	}
	return langs
}
