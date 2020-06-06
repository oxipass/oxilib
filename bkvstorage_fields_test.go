package bslib

import (
	"errors"
	"testing"
)

const cFieldName01 = "nwjkdq dwqjdnwqd dejwdjqkwdq"
const cFieldValue01 = "iueruiefr efkernf erferfe"
const cFieldIcon01 = "fab fa-fort-awesome"

func testHelperCreateField(itemId int64) (fieldId int64, err error) {
	bsInstance := GetInstance()
	err = bsInstance.Unlock(dbPass)
	if err != nil {
		return 0, err
	}
	fieldResult, errField := bsInstance.AddNewField(
		UpdateFieldForm{
			ItemID: itemId,
			BSField: BSField{
				Name:      cFieldName01,
				Icon:      cFieldIcon01,
				ValueType: VTText,
				Value:     cFieldValue01,
			},
		},
	)
	if errField != nil {
		return 0, errField
	}
	if fieldResult.Status != ConstSuccessResponse {
		return 0, errors.New("response is not successful: " + fieldResult.MsgTxt)
	}
	errLock := bsInstance.Lock()
	if errLock != nil {
		return 0, errLock
	}
	return fieldResult.FieldID, nil
}

func testHelperCreateItemAndField() (itemId int64, fieldId int64, err error) {
	itemId, err = testHelperCreateItem()
	if err != nil {
		return 0, 0, err
	}
	fieldId, err = testHelperCreateField(itemId)
	if err != nil {
		return 0, 0, err
	}
	return itemId, fieldId, nil
}

func TestAddField(t *testing.T) {
	itemId, fieldId, err := testHelperCreateItemAndField()
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

	fields, errFields := bsInstance.ReadFieldsByItemID(itemId)
	if errFields != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	for _, field := range fields {
		if field.ID == fieldId {
			if field.Icon == cFieldIcon01 &&
				field.Name == cFieldName01 &&
				field.Value == cFieldValue01 {
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
	_, fieldId, err := testHelperCreateItemAndField()
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

	respDel, errDel := bsInstance.DeleteField(UpdateFieldForm{
		BSField: BSField{
			ID: fieldId,
		},
	})
	if errDel != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	if respDel.Status != ConstSuccessResponse {
		t.Error(errors.New("response is not successful: " + respDel.MsgTxt))
		t.FailNow()
		return
	}

	field, errField := bsInstance.ReadFieldsByFieldID(fieldId)
	if errField != nil {
		t.Error(errPass)
		t.FailNow()
		return
	}

	if field.ID == fieldId {
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
