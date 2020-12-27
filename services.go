package bslib

import "errors"

// All public services to be used by mobile or web app are here

func BSLibService(functionName string, inputJSON string) string {
	switch functionName {
	case "lock":
		return ServiceLockStorage()
	case "init":
		return ServiceInitStorage(inputJSON)
	case "initnew":
		return ServiceInitNewStorage(inputJSON)
	case "additem":
		return ServiceAddNewItem(inputJSON)
	case "readall":
		return ServiceReadAllItems(inputJSON)
	default:
		return formErrorResponse(errors.New(BSERR00020FunctionNotFound + ", " + functionName))
	}
}

func ServiceLockStorage() string {
	storage := GetInstance()
	err := storage.Lock()
	if err != nil {
		return formErrorResponse(err)
	}
	jsonResponse, err := EncodeJSON(CommonResponse{Status: ConstSuccessResponse})
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	return jsonResponse
}

func ServiceInitStorage(jsonInputParams string) string {
	var isForm InitStorageForm
	err := DecodeJSON(jsonInputParams, isForm)
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	storage := GetInstance()
	err = storage.Open(isForm.FileName)
	if err != nil {
		return formErrorResponse(err)
	}
	jsonResponse, err := EncodeJSON(CommonResponse{Status: ConstSuccessResponse})
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	return jsonResponse
}

func ServiceInitNewStorage(jsonInputParms string) string {
	var isForm InitStorageForm
	err := DecodeJSON(jsonInputParms, isForm)
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	storage := GetInstance()
	err = storage.Open(isForm.FileName)
	if err != nil {
		return formErrorResponse(err)
	}

	errSetPassword := storage.SetNewPassword(isForm.Password, isForm.Encryption)
	if errSetPassword != nil {
		return formErrorResponse(errSetPassword)
	}

	jsonResponse, err := EncodeJSON(CommonResponse{Status: ConstSuccessResponse})
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	return jsonResponse
}

// ServiceAddNewItem - wrapper for adding item with JSON for using it with mobile
func ServiceAddNewItem(jsonInputParms string) string {
	var addItemForm UpdateItemForm
	err := DecodeJSON(jsonInputParms, addItemForm)
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	storage := GetInstance()
	response, err := storage.AddNewItem(addItemForm)
	if err != nil {
		return formErrorResponse(err)
	}
	jsonResponse, err := EncodeJSON(response)
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	return jsonResponse
}

// ServiceReadAllItems - wrapper for reading all the items
func ServiceReadAllItems(inputParams string) string {
	var jsonAddItem ReadAllForm
	err := DecodeJSON(inputParams, jsonAddItem)
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	storage := GetInstance()
	allItems, err := storage.ReadAllItems(true, false)
	if err != nil {
		return formErrorResponse(err)
	}
	jsonResponse, err := EncodeJSON(allItems)
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	return jsonResponse
}
