package bykovstorage

import "errors"
import "strings"

func formError(errorID string, errorText ...string) error {
	var finalText string
	for i, errorStr := range errorText {
		if i == 0 {
			finalText = errorStr
		} else {
			finalText = finalText + ", " + errorStr
		}
	}
	return errors.New(errorID + ": " + finalText)
}

func formErrorResponse(err error) string {
	var errorResponse JSONResponseCommon
	strItems := strings.Split(err.Error(), ": ")
	if len(strItems) > 0 && len(strItems[0]) > 11 && strings.Contains(strItems[0], "BSERR") {
		errorResponse.Status = ConstErrorResponse
		errorResponse.MsgNum = strItems[0]
		errorResponse.MsgTxt = err.Error()[len(strItems[0]):]
		jsonStr, jsonErr := EncodeJSON(errorResponse)
		if jsonErr == nil {
			return jsonStr
		}
	}

	errorResponse.Status = ConstErrorResponse
	errorResponse.MsgNum = BSERR00015UnknownError
	errorResponse.MsgTxt = err.Error()

	jsonStr, jsonErr := EncodeJSON(errorResponse)
	if jsonErr == nil {
		return jsonStr
	}

	return jsonErr.Error() + ", " + err.Error()
}

// BSERR00001DbIntegrityCheckFailed - failure during integrity check of database
const BSERR00001DbIntegrityCheckFailed = "BSERR00001_DBINTEGRITYCHECK"

// BSERR00002DbTableCreationFailed - failed to create table during initialization
const BSERR00002DbTableCreationFailed = "BSERR00002_DBTABLECREATIONFAILED"

// BSERR00003DbTransactionFailed - failed to start new transaction
const BSERR00003DbTransactionFailed = "BSERR00003_TRANSACTIONFAILED"

// BSERR00004EncCypherNotExist - cypher is not found or not supported by storage
const BSERR00004EncCypherNotExist = "BSERR00004_CYPHERNOTEXIST"

// BSERR00005ParseTimeFailed - time parsing has failed
const BSERR00005ParseTimeFailed = "BSERR00005_PARSETIMEFAILED"

// BSERR00006DbInsertFailed - insert into database has failed
const BSERR00006DbInsertFailed = "BSERR00006_INSERTFAILED"

// BSERR00008EncEncryptionError - general encryption error
const BSERR00008EncEncryptionError = "BSERR00008_ENCRYPTIONERROR"

// BSERR00009DbNotOpen - operation on the datanase which is not open yet
const BSERR00009DbNotOpen = "BSERR00009_DBNOTOPEN"

// BSERR00010EncWrongPassword - wrong password
const BSERR00010EncWrongPassword = "BSERR00010_WRONGPASSWORD"

// BSERR00011EncCypherNotProvided - cypher algorythm is not provided
const BSERR00011EncCypherNotProvided = "BSERR00011_CYPHERNOTPROVICED"

// BSERR00012DbTxStartFailed - tx start failed
const BSERR00012DbTxStartFailed = "BSERR00012_TXSTARTFAILED"

// BSERR00013DbTxEndFaild - tx end failed
const BSERR00013DbTxEndFaild = "BSERR00013_TXENDFAILED"

// BSERR00014ItemsReadFailed - items read failure
const BSERR00014ItemsReadFailed = "BSERR00014_ITEMSREADFAILED"

// BSERR00015UnknownError - unknown error, details will be in error text
const BSERR00015UnknownError = "BSERR00015_UNKNOWN_ERROR"
