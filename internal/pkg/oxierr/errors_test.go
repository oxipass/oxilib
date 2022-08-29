package oxierr

import (
	"github.com/oxipass/oxilib/consts"
	"github.com/oxipass/oxilib/internal/pkg/utils"
	"github.com/oxipass/oxilib/models"
	"testing"
)

const cError1Test = "error one"
const cError2text = "error two"
const cPrintExpectedVsRetrieved = "Expected '%s', retrieved '%s'"

func TestFormError(t *testing.T) {
	expectedString := BSERR00015UnknownError + ": " + cError1Test + ", " + cError2text
	newErr := FormError(BSERR00015UnknownError, cError1Test, cError2text)
	if newErr.Error() != BSERR00015UnknownError+": "+cError1Test+", "+cError2text {
		t.Errorf(cPrintExpectedVsRetrieved, expectedString, newErr.Error())
		t.FailNow()
		return
	}
}

func TestFormErrorResponse(t *testing.T) {
	newErr := FormError(BSERR00015UnknownError, cError1Test, cError2text)
	response := FormErrorResponse(newErr)
	var messageResponse models.CommonResponse
	err := utils.DecodeJSON(response, &messageResponse)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	if messageResponse.Status != consts.CErrorResponse {
		t.Errorf(cPrintExpectedVsRetrieved, consts.CErrorResponse, messageResponse.Status)
		return
	}
	if messageResponse.MsgTxt != cError1Test+", "+cError2text {
		t.Errorf(cPrintExpectedVsRetrieved, cError1Test+", "+cError2text, messageResponse.MsgTxt)
		return
	}
	if messageResponse.MsgNum != BSERR00015UnknownError {
		t.Errorf(cPrintExpectedVsRetrieved, BSERR00015UnknownError, messageResponse.MsgNum)
		return
	}
}

func TestFormErrorResponseShort(t *testing.T) {
	const unknownError = "unknown error"
	newErr := FormError(unknownError)
	response := FormErrorResponse(newErr)
	var messageResponse models.CommonResponse
	err := utils.DecodeJSON(response, &messageResponse)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	if messageResponse.Status != consts.CErrorResponse {
		t.Errorf(cPrintExpectedVsRetrieved, consts.CErrorResponse, messageResponse.Status)
		return
	}
	if messageResponse.MsgTxt != unknownError {
		t.Errorf(cPrintExpectedVsRetrieved, unknownError, messageResponse.MsgTxt)
		return
	}
	if messageResponse.MsgNum != BSERR00015UnknownError {
		t.Errorf(cPrintExpectedVsRetrieved, BSERR00015UnknownError, messageResponse.MsgNum)
		return
	}
}
