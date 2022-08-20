package oxilib

import (
	"embed"
	"encoding/json"
	"github.com/oxipass/oxilib/models"
)

var (

	//go:embed assets/templates/*.json
	templatesResources embed.FS
)

// GetTagsTemplate - returns embedded template data for tags
func GetTagsTemplate() (tagsTemplate models.TagsTemplateJSON, err error) {
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

func GetFieldsTemplate() (fieldsTemplate models.FieldsTemplateJSON, err error) {
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

func GetItemsTemplate() (itemsTemplate models.ItemsTemplateJSON, err error) {
	fileBytes, errFile := templatesResources.ReadFile(cTemplatesFolder + "/" + cItemsTemplates)
	if errFile != nil {
		return itemsTemplate, errFile
	}
	errUnmarshal := json.Unmarshal(fileBytes, &itemsTemplate)
	if errUnmarshal != nil {
		return itemsTemplate, errUnmarshal
	}
	return itemsTemplate, nil
}

// TODO: Store/update tags templates in database

// TODO: Store/update fields templates in database

// TODO: Store/update items templates in database
