package oxilib

import (
	"errors"
	"testing"
)

const cValueTypesCount = 16

func TestValueTypes(t *testing.T) {
	vTypes := GetValueTypes()
	if vTypes == nil {
		t.Error(errors.New("value types are empty"))
		t.FailNow()
		return
	}
	if len(vTypes) != cValueTypesCount {
		t.Error(errors.New("wrong value types count or unhandled value type is added"))
		t.FailNow()
		return
	}
}
