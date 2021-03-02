package bslib

import (
	"errors"
	"strings"
)

func formError(errorID string, errorText ...string) error {
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

func formErrorResponse(err error) string {
	var errorResponse CommonResponse
	strItems := strings.Split(err.Error(), ": ")
	if len(strItems) > 0 && len(strItems[0]) > 11 && strings.Contains(strItems[0], "BSERR") {
		errorResponse.Status = CErrorResponse
		errorResponse.MsgNum = strItems[0]
		if len(strItems) > 1 && len(strItems[1]) > 3 {
			errorResponse.MsgTxt = err.Error()[len(strItems[0])+2:]
		} else {
			errorResponse.MsgTxt = err.Error()[len(strItems[0]):]
		}
		jsonStr, jsonErr := EncodeJSON(errorResponse)
		if jsonErr == nil {
			return jsonStr
		}
	}

	errorResponse.Status = CErrorResponse
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
const BSERR00006DbInsertFailed = "BSERR00006_DB_INSERT_FAILED"

// BSERR00008EncEncryptionError - general encryption error
const BSERR00008EncEncryptionError = "BSERR00008_ENCRYPTIONERROR"

// BSERR00009DbNotOpen - operation on the datanase which is not open yet
const BSERR00009DbNotOpen = "BSERR00009_DBNOTOPEN"

// BSERR00010EncWrongPassword - wrong password
const BSERR00010EncWrongPassword = "BSERR00010_WRONGPASSWORD"

// BSERR00011EncCypherNotProvided - cypher algorithm is not provided
const BSERR00011EncCypherNotProvided = "BSERR00011_CYPHER_NOT_PROVIDED"

// BSERR00012DbTxStartFailed - tx start failed
const BSERR00012DbTxStartFailed = "BSERR00012_TXSTARTFAILED"

// BSERR00013DbTxEndFaild - tx end failed
const BSERR00013DbTxEndFaild = "BSERR00013_TXENDFAILED"

// BSERR00014ItemsReadFailed - items read failure
const BSERR00014ItemsReadFailed = "BSERR00014_ITEMSREADFAILED"

// BSERR00015UnknownError - unknown error, details will be in error text
const BSERR00015UnknownError = "BSERR00015_UNKNOWN_ERROR"

// BSERR00016DbDeleteFailed - failed to delete item
const BSERR00016DbDeleteFailed = "BSERR00016_DB_DELETE_FAILED"

// BSERR00017JSONFailed - failed to process jSON operation
const BSERR00017JSONFailed = "BSERR00017_JSON_FAILED"

// BSERR00018DbItemNameUpdateFailed - failed to update item name
const BSERR00018DbItemNameUpdateFailed = "BSERR00018_DB_NAME_UPDATE_FAILED"

// BSERR00019ItemNotFound - item not found
const BSERR00019ItemNotFound = "BSERR00019_ITEMNOTFOUND"

// BSERR00020FunctionNotFound - library function requested by API is not found
const BSERR00020FunctionNotFound = "BSERR00020_FUNCTION_NOT_FOUND"

// BSERR00021FieldsReadFailed - fields read failure
const BSERR00021FieldsReadFailed = "BSERR00021_FIELDSREADFAILED"

// BSERR00022ValidationFailed - fields read failure
const BSERR00022ValidationFailed = "BSERR00022_VALIDATIONFAILED"

// BSERR00023UpdateFieldsEmpty - there are no fields to update, all of them are empty
const BSERR00023UpdateFieldsEmpty = "BSERR00023_UPDATE_FIELDS_EMPTY"

// BSERR00024FontAwesomeIconNotFound - icon name is not found in font awesome
const BSERR00024FontAwesomeIconNotFound = "BSERR00024_FA_ICON_NOT_FOUND"

// BSERR00025ItemIdEmptyOrWrong - item id is not provided or is wrong
const BSERR00025ItemIdEmptyOrWrong = "BSERR00025_ITEMID_EMPTY_OR_WRONG"

// BSERR00026DbItemIconUpdateFailed - failed to update item icon
const BSERR00026DbItemIconUpdateFailed = "BSERR00016_DB_ICON_UPDATE_FAILED"

// BSERR00027ItemValidationError - item data validation error
const BSERR00027ItemValidationError = "BSERR00027_ITEM_VALIDATION_FAILED"
