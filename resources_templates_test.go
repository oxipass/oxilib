package oxilib

import (
	"testing"
)

func TestTagsTemplates(t *testing.T) {
	tagsTemplate, err := GetTagsTemplate()
	if err != nil {
		t.Error(err)
	}
	if tagsTemplate.Updated == "" {
		t.Error("tags template update is empty")
	}
	if tagsTemplate.Tags == nil {
		t.Error("tags template array is nil")
	}
	if len(tagsTemplate.Tags) == 0 {
		t.Error("tags template tags is empty")
	}
	for _, tag := range tagsTemplate.Tags {
		if tag.ID == "" {
			t.Error("tag id is empty")
		}
		if tag.Color == "" {
			t.Error("tag color is empty for tag " + tag.ID)
		}
		if len(tag.Color) != 7 {
			t.Error("tag code is not 7 chars long by template '#AABBCC'  for tag id: " + tag.ID)
		} else if tag.Color[0] != '#' {
			t.Error("tag color is not started with '#' for tag id: " + tag.ID)
		}
	}
}

func TestTagsTemplatesTranslations(tst *testing.T) {
	tagsTemplate, err := GetTagsTemplate()
	if err != nil {
		tst.Error(err)
	}
	_, err = initLang("en")
	if err != nil {
		tst.Error(err)
	}
	for _, tag := range tagsTemplate.Tags {
		transl := t(tag.ID)
		if transl == "" {
			tst.Error("tag translation is empty for tag id: " + tag.ID)
		}
		if transl == tag.ID {
			tst.Error("tag translation is equal to tag id for tag id(meaning no translations): " + tag.ID)
		}
	}

}

func TestFieldsTemplates(t *testing.T) {
	fieldsTemplate, err := GetFieldsTemplate()
	if err != nil {
		t.Error(err)
	}
	if fieldsTemplate.Updated == "" {
		t.Error("fields template update is empty")
	}
	if fieldsTemplate.Fields == nil {
		t.Error("fields template array is nil")
	}
	if len(fieldsTemplate.Fields) == 0 {
		t.Error("fields template fields is empty")
	}
	for _, field := range fieldsTemplate.Fields {
		if field.ID == "" {
			t.Error("field id is empty")
		}
		if field.FieldType == "" {
			t.Error("field type is empty for field " + field.ID)
		}
		if field.Icon == "" {
			t.Error("field icon is empty for field " + field.ID)
		}

	}
}

func TestFieldsTemplatesTranslations(tst *testing.T) {
	fieldsTemplate, err := GetFieldsTemplate()
	if err != nil {
		tst.Error(err)
	}
	_, err = initLang("en")
	if err != nil {
		tst.Error(err)
	}
	for _, field := range fieldsTemplate.Fields {
		transl := t(field.ID)
		if transl == "" {
			tst.Error("field translation is empty for field id: " + field.ID)
		}
		if transl == field.ID {
			tst.Error("field translation is equal to field id for field id(meaning no translations): " + field.ID)
		}
	}

}

func TestFieldsTemplatesTypes(tst *testing.T) {
	fieldsTemplate, err := GetFieldsTemplate()
	if err != nil {
		tst.Error(err)
	}
	for _, field := range fieldsTemplate.Fields {

		if CheckValueType(field.FieldType) == false {
			tst.Error("field type is not valid for field " + field.ID)
		}
	}
}

func TestFieldsTemplatesIcon(tst *testing.T) {
	fieldsTemplate, err := GetFieldsTemplate()
	if err != nil {
		tst.Error(err)
	}
	for _, field := range fieldsTemplate.Fields {
		if field.Icon == "" {
			tst.Error("field icon is empty for the field " + field.ID)
		}

		if !CheckIfExistsInFontAwesome(field.Icon) {
			tst.Error("field icon '" + field.Icon + "' does not exist in icons for the field " + field.ID)
		}
	}

}

func TestItemsTemplatesTranslations(tst *testing.T) {
	itemsTemplate, err := GetItemsTemplate()
	if err != nil {
		tst.Error(err)
	}
	_, err = initLang("en")
	if err != nil {
		tst.Error(err)
	}
	for _, item := range itemsTemplate.Items {
		transl := t(item.ID)
		if transl == "" {
			tst.Error("item translation is empty for item id: " + item.ID)
		}
		if transl == item.ID {
			tst.Error("item translation is equal to item id for item id(meaning no translations): " + item.ID)
		}
	}

}

func TestItemsTemplatesIcon(tst *testing.T) {
	itemsTemplate, err := GetItemsTemplate()
	if err != nil {
		tst.Error(err)
	}
	for _, item := range itemsTemplate.Items {
		if item.Icon == "" {
			tst.Error("item icon is empty for the item " + item.ID)
		}

		if !CheckIfExistsInFontAwesome(item.Icon) {
			tst.Error("item icon '" + item.Icon + "' does not exist in icons for the item " + item.ID)
		}
	}

}

func TestItemsTemplatesFields(tst *testing.T) {
	itemsTemplate, err := GetItemsTemplate()
	if err != nil {
		tst.Error(err)
	}
	fieldsTemplate, err := GetFieldsTemplate()
	if err != nil {
		tst.Error(err)
	}

	for _, item := range itemsTemplate.Items {
		if item.FieldsIds == nil {
			tst.Error("item fields is nil for item " + item.ID)
			continue
		}
		if len(item.FieldsIds) == 0 {
			tst.Error("item fields is empty for item " + item.ID)
			continue
		}
		for _, field := range item.FieldsIds {
			if field == "" {
				tst.Error("field id is empty for item " + item.ID)
				continue
			}
			foundTemplate := false
			for _, fTemplate := range fieldsTemplate.Fields {
				if fTemplate.ID == field {
					foundTemplate = true
					break
				}
			}
			if foundTemplate == false {
				tst.Error("field id '" + field + "' does not exist in fields template for item " + item.ID)
			}
		}
	}
}

func TestItemsTemplatesTags(tst *testing.T) {
	itemsTemplate, err := GetItemsTemplate()
	if err != nil {
		tst.Error(err)
	}
	tagsTemplate, err := GetTagsTemplate()
	if err != nil {
		tst.Error(err)
	}

	for _, item := range itemsTemplate.Items {
		if item.TagsIds == nil {
			tst.Error("item tags is nil for item " + item.ID)
			continue
		}
		if len(item.TagsIds) == 0 {
			tst.Error("item tags is empty for item " + item.ID)
			continue
		}
		for _, tag := range item.TagsIds {
			if tag == "" {
				tst.Error("tag id is empty for item " + item.ID)
				continue
			}
			foundTemplate := false
			for _, tTemplate := range tagsTemplate.Tags {
				if tTemplate.ID == tag {
					foundTemplate = true
					break
				}
			}
			if foundTemplate == false {
				tst.Error("tag id '" + tag + "' does not exist in tags template for item " + item.ID)
			}
		}
	}
}
