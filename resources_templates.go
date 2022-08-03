package oxilib

import (
	"embed"
	"encoding/json"
)

var (

	//go:embed templates/*.json
	templatesResources embed.FS
)

// GetTagsTemplate - returns embedded template data for tags
func GetTagsTemplate() (tagsTemplate TagsTemplate, err error) {
	fileBytes, errFile := templatesResources.ReadFile(cTemplatesFolder + "/" + cTagsTemplates)
	if errFile != nil {
		return tagsTemplate, errFile
	}
	errUnmarshal := json.Unmarshal(fileBytes, &tagsTemplate)
	if errUnmarshal != nil {
		return tagsTemplate, errUnmarshal
	}
	return tagsTemplate, nil
}

func GetFieldsTemplate() (fieldsTemplate FieldsTemplate, err error) {
	fileBytes, errFile := templatesResources.ReadFile(cTemplatesFolder + "/" + cFieldsTemplates)
	if errFile != nil {
		return fieldsTemplate, errFile
	}
	errUnmarshal := json.Unmarshal(fileBytes, &fieldsTemplate)
	if errUnmarshal != nil {
		return fieldsTemplate, errUnmarshal
	}
	return fieldsTemplate, nil
}

// TODO: Store/update tags templates in database

// TODO: Store/update fields templates in database

// TODO: GetItemsTemplate - returns embedded template data for items
// TODO: Test items translations
// TODO: Test items icons
// TODO: Test fields templates availability
// TODO: Test tags templates availability

// TODO: Store/update items templates in database
