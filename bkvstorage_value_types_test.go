package bslib

import (
	"errors"
	"testing"
)

func TestValueTypes(t *testing.T) {
	vTypes := GetValueTypes()
	if vTypes == nil {
		t.Error(errors.New("value types are empty"))
		t.FailNow()
		return
	}
	if len(vTypes) != 10 {
		t.Error(errors.New("wrong value types count or unhandled value type is added"))
		t.FailNow()
		return
	}
}
