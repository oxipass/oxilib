package oxierr

import (
	"errors"
	"github.com/oxipass/oxilib/consts"

	"github.com/oxipass/oxilib/internal/pkg/utils"
	"github.com/oxipass/oxilib/models"
	"strings"
)

func FormError(errorID string, errorText ...string) error {
	var finalText string
	for i, errorStr := range errorText {
		if i == 0 {
			finalText = errorStr
		} else {
			finalText = finalText + ", " + errorStr
		}
	}
	if finalText == "" {
		return errors.New(errorID)
	}
	return errors.New(errorID + ": " + finalText)
}

func FormErrorResponse(err error) string {
	var errorResponse models.CommonResponse
	strItems := strings.Split(err.Error(), ": ")
	if len(strItems) > 0 && len(strItems[0]) > 11 && strings.Contains(strItems[0], "BSERR") {
		errorResponse.Status = consts.CErrorResponse
		errorResponse.MsgNum = strItems[0]
		if len(strItems) > 1 && len(strItems[1]) > 3 {
			errorResponse.MsgTxt = err.Error()[len(strItems[0])+2:]
		} else {
			errorResponse.MsgTxt = err.Error()[len(strItems[0]):]
		}
		jsonStr, jsonErr := utils.EncodeJSON(errorResponse)
		if jsonErr == nil {
			return jsonStr
		}
	}

	errorResponse.Status = consts.CErrorResponse
	errorResponse.MsgNum = BSERR00015UnknownError
	errorResponse.MsgTxt = err.Error()

	jsonStr, jsonErr := utils.EncodeJSON(errorResponse)
	if jsonErr == nil {
		return jsonStr
	}

	return jsonErr.Error() + ", " + err.Error()
}
