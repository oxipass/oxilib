package bslib

import "testing"

const cError1Test = "error one"
const cError2text = "error two"

func TestFormError(t *testing.T) {
	expectedString := BSERR00015UnknownError + ": " + cError1Test + ", " + cError2text
	newErr := formError(BSERR00015UnknownError, cError1Test, cError2text)
	if newErr.Error() != BSERR00015UnknownError+": "+cError1Test+", "+cError2text {
		t.Errorf("Expected '%s', retrieved '%s'", expectedString, newErr.Error())
		t.FailNow()
		return
	}
}

func TestFormErrorResponse(t *testing.T) {
	newErr := formError(BSERR00015UnknownError, cError1Test, cError2text)
	response := formErrorResponse(newErr)
	var messageResponse JSONResponseCommon
	err := DecodeJSON(response, &messageResponse)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	if messageResponse.Status != CErrorResponse {
		t.Errorf("Expected '%s', retrieved '%s'", CErrorResponse, messageResponse.Status)
		return
	}
	if messageResponse.MsgTxt != cError1Test+", "+cError2text {
		t.Errorf("Expected '%s', retrieved '%s'", cError1Test+", "+cError2text, messageResponse.MsgTxt)
		return
	}
	if messageResponse.MsgNum != BSERR00015UnknownError {
		t.Errorf("Expected '%s', retrieved '%s'", BSERR00015UnknownError, messageResponse.MsgNum)
		return
	}
}
