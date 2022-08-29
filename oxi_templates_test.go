package oxilib

import (
	"github.com/oxipass/oxilib/assets"
	"testing"
)

func TestDefaultTagsAvailable(t *testing.T) {
	storage := GetInstance()
	tags, err := storage.GetTags()
	if err != nil {
		t.Error(err)
	}
	tagsTemplate, errTempl := assets.GetTagsTemplate()
	if errTempl != nil {
		t.Error(errTempl)
	}

	for _, tagTemplate := range tagsTemplate.Tags {
		found := false
		for _, tag := range tags {
			if tag.ExtId == tagTemplate.ID {
				found = true
				break
			}
		}
		if !found {
			t.Error("tag " + tagTemplate.ID + " is not found as default tag in tags")
		}

	}
}
