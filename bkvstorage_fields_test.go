package bslib

import (
	"errors"
	"testing"
)

func TestAddField(t *testing.T) {
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

	fieldName := generateRandomString(20)
	fieldValue := generateRandomString(20)
	fieldResult, errField := bsInstance.AddNewField(
		UpdateFieldForm{
			ItemID: response.ItemID,
			BSField: BSField{
				Name:      fieldName,
				Icon:      "fab fa-amazon",
				ValueType: VTText,
				Value:     fieldValue,
			},
		},
	)
	if errField != nil {
		t.Error(errField)
		t.FailNow()
		return
	}
	if fieldResult.Status != ConstSuccessResponse {
		t.Error(errors.New("response is not successful: " + fieldResult.MsgTxt))
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

	fields, errFields := bsInstance.ReadFieldsByItemID(response.ItemID)
	if errFields != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	for _, field := range fields {
		if field.ID == fieldResult.FieldID {
			if field.Icon == "fab fa-amazon" &&
				field.Name == fieldName &&
				field.Value == fieldValue {
				return
			}
			t.Error("ID is found but content is wrong")
			t.FailNow()
			return
		}
	}
	t.Error("Saved ID is not found")
	t.FailNow()
}

func TestDeleteField(t *testing.T) {
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
		t.Error(errors.New("response is not successful"))
		t.FailNow()
		return
	}

	fieldName := generateRandomString(20)
	fieldValue := generateRandomString(20)
	fieldResult, errField := bsInstance.AddNewField(
		UpdateFieldForm{
			ItemID: response.ItemID,
			BSField: BSField{
				Name:      fieldName,
				Icon:      "fab fa-amazon",
				ValueType: VTText,
				Value:     fieldValue,
			},
		},
	)

	if errField != nil {
		t.Error(errField)
		t.FailNow()
		return
	}

	if fieldResult.Status != ConstSuccessResponse {
		t.Error(errors.New("response is not successful: " + fieldResult.MsgTxt))
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

	respDel, errDel := bsInstance.DeleteField(UpdateFieldForm{
		BSField: BSField{
			ID: fieldResult.FieldID,
		},
	})
	if errDel != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	if respDel.Status != ConstSuccessResponse {
		t.Error(errors.New("response is not successful: " + fieldResult.MsgTxt))
		t.FailNow()
		return
	}

	field, errField := bsInstance.ReadFieldsByFieldID(fieldResult.FieldID)
	if errField != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	if field.ID == fieldResult.FieldID {
		if field.Deleted == false {
			t.Error("Field is not marked as deleted")
			t.FailNow()
			return
		}
		return
	}
	t.Error("Deleted field is not found")
	t.FailNow()
}
