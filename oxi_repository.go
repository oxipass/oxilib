package oxilib

const (
	VTText             = "text"     // simple test
	VTLongText         = "longtext" // long text (multi lines, scroll when viewing)
	VTCard             = "card"     // any kind of 16 digits card to show it like 1111 2222 3333 4444
	VTPassword         = "password" // any hidden field, password, pin etc. It will be masked with asterisks
	VTLink             = "link"     // link to any internet page, special logic for http/https
	VTEmail            = "email"    // email address
	VTPhone            = "phone"    // phone number, TODO: show country code separately
	VTDate             = "date"     // any date, show calendar as input
	VTExpirationDate   = "expdate"  // the same as date, but with expiration date
	VTTime             = "time"     // time, show clock as input
	VTOneTimePassword  = "otp"      // one time password, show input with 6 digits changing every 30 seconds
	VTPreviousPassword = "prevpass" // previous password, the password backed up before here changing it
	VTRecoveryPhrase   = "recovery" // recovery phrase, show words with numbers separated by spaces
	VTQRCode           = "qrcode"   // QR code, show QR code
	VTAddress          = "address"  // address, show address
	VTPin              = "pin"      // Pin code, password with digits only
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
		VTRecoveryPhrase,
		VTQRCode,
		VTAddress,
		VTPin,
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
