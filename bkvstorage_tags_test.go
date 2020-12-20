package bslib

import (
	"errors"
	"testing"
)

const cTestTag1 = "test_tag1"
const cTestTag2 = "test_tag2"

func testHelperCreateItemAndTag(testTag string) (itemId int64, tagId int64, err error) {
	bsInstance := GetInstance()
	err = bsInstance.Unlock(dbPass)
	if err != nil {
		return 0, 0, err
	}
	response, err := bsInstance.AddNewItem(
		UpdateItemForm{
			BSItem: BSItem{
				Name: cTestItemName01,
				Icon: cTestItemIcon01,
			},
		},
	)
	if err != nil {
		return 0, 0, err
	}
	if response.Status != ConstSuccessResponse {
		return 0, 0, errors.New("response is not successful")
	}

	responseTag, err := bsInstance.AddNewTag(
		UpdateTagForm{
			0,
			BSTag{
				Name: testTag,
			},
		},
	)
	if err != nil {
		return 0, 0, err
	}

	if responseTag.Status != ConstSuccessResponse {
		return 0, 0, errors.New("response is not successful for tag creation")
	}

	errLock := bsInstance.Lock()
	if errLock != nil {
		return 0, 0, errLock
	}
	return response.ItemID, responseTag.TagId, nil
}

func TestCreateItemAndTag(t *testing.T) {
	_, tagId, err := testHelperCreateItemAndTag(cTestTag1)

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
		if tag.ID == tagId && tag.Name == cTestTag1 && tag.Deleted == false {
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
	itemId1, tagId1, err := testHelperCreateItemAndTag(cTestTag1)
	_, tagId2, err := testHelperCreateItemAndTag(cTestTag2)

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
		UpdateTagForm{
			ItemID: itemId1,
			BSTag: BSTag{
				ID: tagId1,
			},
		},
	)
	if errTA != nil {
		t.Error(errTA)
		t.FailNow()
		return
	}

	if responseTA.Status != ConstSuccessResponse {
		t.Error(errors.New("response is not successful for tag assignment"))
		t.FailNow()
		return
	}

	responseTA2, errTA2 := bsInstance.AssignTag(
		UpdateTagForm{
			ItemID: itemId1,
			BSTag: BSTag{
				ID: tagId2,
			},
		},
	)
	if errTA2 != nil {
		t.Error(errTA2)
		t.FailNow()
		return
	}

	if responseTA2.Status != ConstSuccessResponse {
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
		if tag.ID == tagId1 && tag.Name == cTestTag1 {
			foundTags++
		}
		if tag.ID == tagId2 && tag.Name == cTestTag2 {
			foundTags++
		}
	}
	if foundTags != 2 {
		t.Error("something went wrong, assigned tags were not found")
		t.FailNow()
		return
	}
}
