package bslib

// All public services to be used by mobile or web app are here

func ServiceLockStorage() string {
	storage := GetInstance()
	err := storage.Lock()
	if err != nil {
		return formErrorResponse(err)
	}
	jsonResponse, err := EncodeJSON(JSONResponseCommon{Status: ConstSuccessResponse})
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	return jsonResponse
}

func ServiceInitStorage(jsonInputParms string) string {
	var jsonInitStorage JSONInputInitStorage
	err := DecodeJSON(jsonInputParms, jsonInitStorage)
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	storage := GetInstance()
	err = storage.Open(jsonInitStorage.FileName)
	if err != nil {
		return formErrorResponse(err)
	}
	jsonResponse, err := EncodeJSON(JSONResponseCommon{Status: ConstSuccessResponse})
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	return jsonResponse
}

func ServiceInitNewStorage(jsonInputParms string) string {
	var jsonInitStorage JSONInputInitStorage
	err := DecodeJSON(jsonInputParms, jsonInitStorage)
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	storage := GetInstance()
	err = storage.Open(jsonInitStorage.FileName)
	if err != nil {
		return formErrorResponse(err)
	}

	errSetPassword := storage.SetNewPassword(jsonInitStorage.Password, jsonInitStorage.Encryption)
	if errSetPassword != nil {
		return formErrorResponse(errSetPassword)
	}

	jsonResponse, err := EncodeJSON(JSONResponseCommon{Status: ConstSuccessResponse})
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	return jsonResponse
}

// ServiceAddNewItem - wrapper for adding item with JSON for using it with mobile
func ServiceAddNewItem(jsonInputParms string) string {
	var jsonAddItem JSONInputUpdateItem
	err := DecodeJSON(jsonInputParms, jsonAddItem)
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	storage := GetInstance()
	response, err := storage.AddNewItem(jsonAddItem)
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
func ServiceReadAllItems(jsonInputParms string) string {
	var jsonAddItem JSONInputReadAll
	err := DecodeJSON(jsonInputParms, jsonAddItem)
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	storage := GetInstance()
	allItems, err := storage.ReadAllItems()
	if err != nil {
		return formErrorResponse(err)
	}
	jsonResponse, err := EncodeJSON(allItems)
	if err != nil {
		return formErrorResponse(formError(BSERR00017JSONFailed, err.Error()))
	}
	return jsonResponse
}
