package oxilib

import (
	"errors"
	"github.com/oxipass/oxilib/consts"
	"github.com/oxipass/oxilib/models"
	"testing"
)

const cTestTag1 = "test_tag1"
const cTestTagColor1 = "#ffff00"
const cTestTag2 = "test_tag2"
const cTesttagColor2 = "#ff0000"

func testHelperCreateItemAndTag(testTag string, testColor string) (itemId int64, tagId int64, err error) {
	bsInstance := GetInstance()
	err = bsInstance.Unlock(dbPass)
	if err != nil {
		return 0, 0, err
	}
	response, err := bsInstance.AddNewItem(
		models.UpdateItemForm{
			OxiItem: models.OxiItem{
				Name: cTestItemName01,
				Icon: cTestItemIcon01,
			},
		},
	)
	if err != nil {
		return 0, 0, err
	}
	if response.Status != consts.CSuccessResponse {
		return 0, 0, errors.New("response is not successful")
	}

	responseTag, err := bsInstance.AddNewTag(
		models.UpdateTagForm{
			0,
			models.OxiTag{
				Name:  testTag,
				Color: testColor,
			},
		},
	)
	if err != nil {
		return 0, 0, err
	}

	if responseTag.Status != consts.CSuccessResponse {
		return 0, 0, errors.New("response is not successful for tag creation")
	}

	errLock := bsInstance.Lock()
	if errLock != nil {
		return 0, 0, errLock
	}
	return response.ItemID, responseTag.TagId, nil
}

func TestCreateItemAndTag(t *testing.T) {
	_, tagId, err := testHelperCreateItemAndTag(cTestTag1, cTestTagColor1)

	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	bsInstance := GetInstance()

	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	availableTags, errTags := bsInstance.GetTags()
	if errTags != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	found := false
	for _, tag := range availableTags {
		if tag.ID == tagId && tag.Name == cTestTag1 &&
			tag.Color == cTestTagColor1 && tag.Deleted == false {
			found = true
			break
		}
	}
	if found == false {
		t.Error("added tag was not found ")
		t.FailNow()
		return
	}
}

func TestAssignTagToItems(t *testing.T) {
	itemId1, tagId1, err := testHelperCreateItemAndTag(cTestTag1, cTestTagColor1)
	_, tagId2, err := testHelperCreateItemAndTag(cTestTag2, cTesttagColor2)

	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	bsInstance := GetInstance()

	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	responseTA, errTA := bsInstance.AssignTag(
		models.UpdateTagForm{
			ItemID: itemId1,
			OxiTag: models.OxiTag{
				ID: tagId1,
			},
		},
	)
	if errTA != nil {
		t.Error(errTA)
		t.FailNow()
		return
	}

	if responseTA.Status != consts.CSuccessResponse {
		t.Error(errors.New("response is not successful for tag assignment"))
		t.FailNow()
		return
	}

	responseTA2, errTA2 := bsInstance.AssignTag(
		models.UpdateTagForm{
			ItemID: itemId1,
			OxiTag: models.OxiTag{
				ID: tagId2,
			},
		},
	)
	if errTA2 != nil {
		t.Error(errTA2)
		t.FailNow()
		return
	}

	if responseTA2.Status != consts.CSuccessResponse {
		t.Error(errors.New("response is not successful for tag assignment"))
		t.FailNow()
		return
	}

	foundTags := 0
	tags, errTags := bsInstance.ReadTagsByItemID(itemId1)
	if errTags != nil {
		t.Error(errTags)
		t.FailNow()
		return
	}
	for _, tag := range tags {
		if tag.ID == tagId1 && tag.Name == cTestTag1 && tag.Color == cTestTagColor1 && tag.Deleted == false {
			foundTags++
		}
		if tag.ID == tagId2 && tag.Name == cTestTag2 && tag.Color == cTesttagColor2 && tag.Deleted == false {
			foundTags++
		}
	}
	if foundTags != 2 {
		t.Error("something went wrong, assigned tags were not found")
		t.FailNow()
		return
	}
}
