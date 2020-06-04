package bslib

import (
	"errors"
	"testing"
)

const cTestItemName01 = "hjb cwec78hduycbwj dbwne w"
const cTestItemIcon01 = "fas fa-ambulance"

const cTestItemName02 = "98jmwhj2ndycwbcjdwlmdk"
const cTestItemIcon02 = "fab fa-linkedin"

func testHelperCreateItem(t *testing.T) (itemId int64, testPassed bool) {
	bsInstance := GetInstance()
	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return 0, false
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
		t.Error(err)
		t.FailNow()
		return 0, false
	}
	if response.Status != ConstSuccessResponse {
		t.Error(errors.New("response is not successful"))
		t.FailNow()
		return 0, false
	}
	errLock := bsInstance.Lock()
	if errLock != nil {
		t.Error(errLock)
		t.FailNow()
		return 0, false
	}
	return response.ItemID, false
}

func TestUpdateItemName(t *testing.T) {
	itemId, testPassed := testHelperCreateItem(t)

	if !testPassed {
		return
	}
	bsInstance := GetInstance()

	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	updateResponse, errUpdated := bsInstance.UpdateItem(
		UpdateItemForm{
			BSItem: BSItem{
				ID:   itemId,
				Name: cTestItemName02,
			},
		},
	)
	if errUpdated != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
	if updateResponse.Status != ConstSuccessResponse {
		t.Error(errors.New("update response is not successful"))
		t.FailNow()
		return
	}

	errLock := bsInstance.Lock()
	if errLock != nil {
		t.Error(errLock)
		t.FailNow()
		return
	}

	errPass = bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	item, respErr := bsInstance.ReadItemById(itemId)
	if respErr != nil {
		t.Error(respErr)
		t.FailNow()
		return
	}

	if item.ID != itemId {
		t.Error(errors.New("response item is wrong"))
		t.FailNow()
		return
	}

	if item.Name != cTestItemName02 {
		t.Errorf("Expected '%s' after update, retrieved '%s'", cTestItemName02, item.Name)
		t.FailNow()
		return
	}

}

func TestDeleteItem(t *testing.T) {
	itemId, testPassed := testHelperCreateItem(t)

	if !testPassed {
		return
	}
	bsInstance := GetInstance()

	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	delResponse, errDelete := bsInstance.DeleteItem(
		UpdateItemForm{
			BSItem: BSItem{
				ID: itemId,
			},
		},
	)
	if errDelete != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
	if delResponse.Status != ConstSuccessResponse {
		t.Error(errors.New("deletion response is not successful"))
		t.FailNow()
		return
	}

	errLock := bsInstance.Lock()
	if errLock != nil {
		t.Error(errLock)
		t.FailNow()
		return
	}

	errPass = bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	items, respErr := bsInstance.ReadAllItems(false)
	if respErr != nil {
		t.Error(respErr)
		t.FailNow()
		return
	}

	for _, item := range items {
		if item.ID == itemId && item.Deleted == false {
			t.Error(errors.New("deleted item is found as not deleted"))
			t.FailNow()
			return
		}
	}

}

const cTestNonExistingIcon = "djcndkcnkd"

func TestAddItemWithNonExistingIcon(t *testing.T) {
	bsInstance := GetInstance()
	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	response, err := bsInstance.AddNewItem(
		UpdateItemForm{
			BSItem: BSItem{
				Name: cTestItemName01,
				Icon: cTestNonExistingIcon,
			},
		},
	)
	if err == nil {
		t.Error(errors.New("no error returned in spite of the fact that icon is not existing"))
		t.FailNow()
		return
	}
	if response.Status == ConstSuccessResponse {
		t.Error(errors.New("item was added in spite of the fact that icon is not existing"))
		t.FailNow()
		return
	}
}

func TestAddItem(t *testing.T) {
	itemId, testPassed := testHelperCreateItem(t)

	if !testPassed {
		return
	}
	bsInstance := GetInstance()

	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
	item, respErr := bsInstance.ReadItemById(itemId)
	if respErr != nil {
		t.Error(respErr)
		t.FailNow()
		return
	}
	if item.ID == itemId && item.Name == cTestItemName01 {
		return
	}

	t.Error(errors.New("created item is not found"))
	t.FailNow()

}

func TestUpdateItemIcon(t *testing.T) {
	itemId, testPassed := testHelperCreateItem(t)

	if !testPassed {
		return
	}
	bsInstance := GetInstance()

	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	updateResponse, errUpdated := bsInstance.UpdateItem(
		UpdateItemForm{
			BSItem: BSItem{
				ID:   itemId,
				Icon: cTestItemIcon02,
			},
		},
	)
	if errUpdated != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
	if updateResponse.Status != ConstSuccessResponse {
		t.Error(errors.New("update response is not successful"))
		t.FailNow()
		return
	}

	errLock := bsInstance.Lock()
	if errLock != nil {
		t.Error(errLock)
		t.FailNow()
		return
	}

	errPass = bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	item, respErr := bsInstance.ReadItemById(itemId)
	if respErr != nil {
		t.Error(respErr)
		t.FailNow()
		return
	}

	if item.ID != itemId {
		t.Error(errors.New("response item is wrong"))
		t.FailNow()
		return
	}

	if item.Icon != cTestItemIcon02 {
		t.Errorf("Expected '%s' after update, retrieved '%s'", cTestItemIcon02, item.Icon)
		t.FailNow()
		return
	}
}
