package bykovstorage

// All public services to be used by mobile or web app are here

// ServiceAddNewItem - wrapper for adding item with JSON for using it with mobile
func ServiceAddNewItem(jsonInputParms string) string {
	var jsonAddItem JSONInputAddItem
	err := DecodeJSON(jsonInputParms, jsonAddItem)
	if err != nil {
		return formErrorResponse(err)
	}
	storage := GetInstance()
	response, err := storage.AddNewItem(jsonAddItem)
	if err != nil {
		return formErrorResponse(err)
	}
	jsonResponse, err := EncodeJSON(response)
	if err != nil {
		return formErrorResponse(err)
	}
	return jsonResponse
}

// ServiceReadAllItems - wrapper for reading all the items
func ServiceReadAllItems(jsonInputParms string) string {
	var jsonAddItem JSONInputReadAll
	err := DecodeJSON(jsonInputParms, jsonAddItem)
	if err != nil {
		return formErrorResponse(err)
	}
	storage := GetInstance()
	allItems, err := storage.ReadAllItems()
	if err != nil {
		return formErrorResponse(err)
	}
	jsonResponse, err := EncodeJSON(allItems)
	if err != nil {
		return formErrorResponse(err)
	}
	return jsonResponse
}
