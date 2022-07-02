package oxilib

import (
	"testing"
)

type testStruct struct {
	CheckStr  string `json:"check_str"`
	CheckBool bool   `json:"check_bool"`
	CheckInt  int64  `json:"check_int"`
}

func TestEncodeEmptyJSON(t *testing.T) {
	_, err := EncodeJSON(make(chan int))
	if err == nil {
		t.Errorf("Expected error, retrieved nil")
	}
}

func TestEncodeJSON(t *testing.T) {
	var wStr testStruct
	wStr.CheckBool = true
	wStr.CheckInt = 55
	wStr.CheckStr = "TEST01"
	prepJSON := `{"check_str":"TEST01","check_bool":true,"check_int":55}`

	encJson, err := EncodeJSON(wStr)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}

	if encJson != prepJSON {
		t.Errorf("Expected %s, retrieved %s", prepJSON, encJson)
		t.FailNow()
		return
	}

}
