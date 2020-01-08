package bslib

import (
	"errors"
	"log"
	"testing"
)

func TestWrapperServiceFuncNotExists(t *testing.T) {
	log.Println("Checking not existing service function")
	const cNotExistingFunction = "blahblahfunc"
	var messageResponse JSONResponseCommon

	response := BSLibService(cNotExistingFunction, "")

	err := DecodeJSON(response, &messageResponse)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}

	if messageResponse.Status != CErrorResponse {
		t.Error("Should return error, here")
		return
	}
}

func TestServiceLogin(t *testing.T) {
	filename := generateTempFilename()
	dbPass = generateRandomString(12)

	encodedJson, err := EncodeJSON(JSONInputInitStorage{
		FileName:   filename,
		Encryption: "AES256",
		Password:   dbPass})
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	response := ServiceInitNewStorage(encodedJson)

	var messageResponse JSONResponseCommon
	err = DecodeJSON(response, &messageResponse)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	if messageResponse.Status == CErrorResponse {
		t.Error(errors.New(messageResponse.MsgTxt))
		t.FailNow()
		return
	}

	response = ServiceLockStorage()
	err = DecodeJSON(response, &messageResponse)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}

	if messageResponse.Status == CErrorResponse {
		t.Error(errors.New(messageResponse.MsgTxt))
		t.FailNow()
		return
	}
}
