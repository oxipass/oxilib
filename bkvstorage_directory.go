package bslib

import "strings"

func GetValueTypes() (vTypes []string) {
	return []string{
		"text",
		"longtext",
		"card",
		"password",
		"link",
		"email",
		"phone",
		"date",
		"expdate",
		"time",
		"2fa",
		"prevpass",
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
