package assets

import (
	"embed"
	"encoding/json"
	"github.com/oxipass/oxilib/consts"
	"github.com/oxipass/oxilib/models"
)

var (

	//go:embed templates/*.json
	templatesResources embed.FS
)

// GetTagsTemplate - returns embedded template data for tags
func GetTagsTemplate() (tagsTemplate models.TagsTemplateJSON, err error) {
	fileBytes, errFile := templatesResources.ReadFile(consts.CTemplatesFolder + "/" + consts.CTagsTemplates)
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
	fileBytes, errFile := templatesResources.ReadFile(consts.CTemplatesFolder + "/" + consts.CFieldsTemplates)
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
	fileBytes, errFile := templatesResources.ReadFile(consts.CTemplatesFolder + "/" + consts.CItemsTemplates)
	if errFile != nil {
		return itemsTemplate, errFile
	}
	errUnmarshal := json.Unmarshal(fileBytes, &itemsTemplate)
	if errUnmarshal != nil {
		return itemsTemplate, errUnmarshal
	}
	return itemsTemplate, nil
}

// TODO: Store/update tags templates in db

// TODO: Store/update fields templates in db

// TODO: Store/update items templates in db
