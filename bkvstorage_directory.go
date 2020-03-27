package bslib

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

func GetFontAwesomeList() (faValues []string) {
	return faCachedValues
}

func CheckIfExistsFontAwesome(faCheckValue string) bool {
	for _, faFound := range faCachedValues {
		if faFound == faCheckValue {
			return true
		}
	}
	return false
}

func SearchFontAwesomeList(term string) (faValues []string) {
	lowerTerm := strings.ToLower(term)
	for _, faFound := range faCachedValues {
		if strings.Contains(faFound, lowerTerm) {
			faValues = append(faValues, faFound)
		}
	}
	return faValues
}
