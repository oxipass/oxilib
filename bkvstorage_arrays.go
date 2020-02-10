package bslib

import "strings"

func GetValueTypes() (vTypes []string) {
	return []string{"text", "longtext", "password", "link", "email", "phone", "date", "expdate", "time", "2fa"}
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
