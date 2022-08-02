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

// TODO: Store/update tags templates in database

// TODO: GetFieldsTemplate - returns embedded template data for fields
// TODO: Test fields translations
// TODO: Test fields icons
// TODO: Test fields types
// TODO: Store/update fields templates in database

// TODO: GetItemsTemplate - returns embedded template data for items
// TODO: Test items translations
// TODO: Test items icons
// TODO: Test fields templates availability
// TODO: Test tags templates availability
// TODO: Store/update items templates in database
