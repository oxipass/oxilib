package bslib

import (
	"errors"
	"testing"
)

const cFieldName01 = "nwjkdq dwqjdnwqd dejwdjqkwdq"
const cFieldValue01 = "iueruiefr efkernf erferfe"
const cFieldIcon01 = "fab fa-fort-awesome"

func testHelperCreateField(t *testing.T, itenId int64) (fieldId int64, testPassed bool) {
	bsInstance := GetInstance()
	fieldResult, errField := bsInstance.AddNewField(
		UpdateFieldForm{
			ItemID: itenId,
			BSField: BSField{
				Name:      cFieldName01,
				Icon:      cFieldIcon01,
				ValueType: VTText,
				Value:     cFieldValue01,
			},
		},
	)
	if errField != nil {
		t.Error(errField)
		t.FailNow()
		return 0, false
	}
	if fieldResult.Status != ConstSuccessResponse {
		t.Error(errors.New("response is not successful: " + fieldResult.MsgTxt))
		t.FailNow()
		return 0, false
	}
	errLock := bsInstance.Lock()
	if errLock != nil {
		t.Error(errLock)
		t.FailNow()
		return 0, false
	}
	return fieldResult.FieldID, true
}

func testHelperCreateItemAndField(t *testing.T) (itemId int64, fieldId int64, testPassed bool) {
	itemId, testPassed = testHelperCreateItem(t)
	if !testPassed {
		return 0, 0, false
	}
	fieldId, testPassed = testHelperCreateField(t, itemId)
	if !testPassed {
		return 0, 0, false
	}
	return itemId, fieldId, true
}

func TestAddField(t *testing.T) {
	itemId, fieldId, testPassed := testHelperCreateItemAndField(t)
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
	_, fieldId, testPassed := testHelperCreateItemAndField(t)
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
