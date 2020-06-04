package bslib

import (
	"errors"
	"testing"
)

func TestUpdateItemName(t *testing.T) {
	bsInstance := GetInstance()
	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
	itemName := generateRandomString(20)
	response, err := bsInstance.AddNewItem(
		UpdateItemForm{
			BSItem: BSItem{
				Name: itemName,
				Icon: "fas fa-ambulance"},
		},
	)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	if response.Status != ConstSuccessResponse {
		t.Error(errors.New("response is not successful"))
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

	itemName2 := generateRandomString(20)
	updateResponse, errUpdated := bsInstance.UpdateItem(
		UpdateItemForm{
			BSItem: BSItem{
				ID:   response.ItemID,
				Name: itemName2,
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

	errLock = bsInstance.Lock()
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

	item, respErr := bsInstance.ReadItemById(response.ItemID)
	if respErr != nil {
		t.Error(respErr)
		t.FailNow()
		return
	}

	if item.ID != response.ItemID {
		t.Error(errors.New("response item is wrong"))
		t.FailNow()
		return
	}

	if item.Name != itemName2 {
		t.Errorf("Expected '%s' after update, retrieved '%s'", itemName2, item.Name)
		t.FailNow()
		return
	}

}

func TestDeleteItem(t *testing.T) {
	bsInstance := GetInstance()
	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
	itemName := generateRandomString(20)
	response, err := bsInstance.AddNewItem(
		UpdateItemForm{
			BSItem: BSItem{
				Name: itemName,
				Icon: "fas fa-ambulance"},
		},
	)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	if response.Status != ConstSuccessResponse {
		t.Error(errors.New("response is not successful"))
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

	delResponse, errDelete := bsInstance.DeleteItem(
		UpdateItemForm{
			BSItem: BSItem{
				ID: response.ItemID,
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

	errLock = bsInstance.Lock()
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
		if item.ID == response.ItemID && item.Deleted == false {
			t.Error(errors.New("deleted item is found as not deleted"))
			t.FailNow()
			return
		}
	}

}

func TestAddItemWithNonExistingIcon(t *testing.T) {
	bsInstance := GetInstance()
	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
	itemName := generateRandomString(20)
	response, err := bsInstance.AddNewItem(
		UpdateItemForm{
			BSItem: BSItem{
				Name: itemName,
				Icon: "ewfdwejfnerfkj",
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
	bsInstance := GetInstance()
	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
	itemName := generateRandomString(20)
	response, err := bsInstance.AddNewItem(
		UpdateItemForm{
			BSItem: BSItem{
				Name: itemName,
				Icon: "fas fa-ambulance",
			},
		},
	)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	if response.Status != ConstSuccessResponse {
		t.Error(errors.New("response is not successfull"))
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
	item, respErr := bsInstance.ReadItemById(response.ItemID)
	if respErr != nil {
		t.Error(respErr)
		t.FailNow()
		return
	}
	if item.ID == response.ItemID && item.Name == itemName {
		return
	}

	t.Error(errors.New("created item is not found"))
	t.FailNow()

}

func TestUpdateItemIcon(t *testing.T) {
	bsInstance := GetInstance()
	errPass := bsInstance.Unlock(dbPass)
	if errPass != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}
	itemName := generateRandomString(20)
	response, err := bsInstance.AddNewItem(
		UpdateItemForm{
			BSItem: BSItem{
				Name: itemName,
				Icon: "fas fa-ambulance"},
		},
	)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	if response.Status != ConstSuccessResponse {
		t.Error(errors.New("response is not successful"))
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

	itemIcon := "fab fa-linkedin"
	updateResponse, errUpdated := bsInstance.UpdateItem(
		UpdateItemForm{
			BSItem: BSItem{
				ID:   response.ItemID,
				Icon: itemIcon,
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

	errLock = bsInstance.Lock()
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

	item, respErr := bsInstance.ReadItemById(response.ItemID)
	if respErr != nil {
		t.Error(respErr)
		t.FailNow()
		return
	}

	if item.ID != response.ItemID {
		t.Error(errors.New("response item is wrong"))
		t.FailNow()
		return
	}

	if item.Icon != itemIcon {
		t.Errorf("Expected '%s' after update, retrieved '%s'", itemIcon, item.Icon)
		t.FailNow()
		return
	}

}
