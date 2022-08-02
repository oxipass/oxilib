package oxilib

import "testing"

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


