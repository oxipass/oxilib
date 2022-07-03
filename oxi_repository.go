package oxilib

import "strings"

const (
	VTText             = "text"
	VTLongText         = "longtext"
	VTCard             = "card"
	VTPassword         = "password"
	VTLink             = "link"
	VTEmail            = "email"
	VTPhone            = "phone"
	VTDate             = "date"
	VTExpirationDate   = "expdate"
	VTTime             = "time"
	VTOneTimePassword  = "otp"
	VTPreviousPassword = "prevpass"
)

func GetValueTypes() (vTypes []string) {
	return []string{
		VTText,
		VTLongText,
		VTCard,
		VTPassword,
		VTLink,
		VTEmail,
		VTPhone,
		VTDate,
		VTExpirationDate,
		VTTime,
		VTOneTimePassword,
		VTPreviousPassword,
	}
}

func CheckValueType(vType string) bool {
	for _, v := range GetValueTypes() {
		if v == vType {
			return true
		}
	}
	return false
}

var cachedFA []string

func GetFontAwesomeList() (faValues []string) {
	if cachedFA != nil {
		return cachedFA
	}
	for k := range faCachedValues {
		cachedFA = append(cachedFA, k)
	}
	return cachedFA
}

func CheckIfExistsFontAwesome(faCheckValue string) bool {
	for faKey := range faCachedValues {
		if faKey == faCheckValue {
			return true
		}
	}
	return false
}

func SearchFontAwesomeList(term string) (faValues []string) {
	lowerTerm := strings.ToLower(term)
	for faKey, faSearchTerms := range faCachedValues {
		if strings.Contains(faKey, lowerTerm) || strings.Contains(faSearchTerms, lowerTerm) {
			faValues = append(faValues, faKey)
		}
	}
	return faValues
}
