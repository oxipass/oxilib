package bslib

func (storage *StorageSingleton) GetValueTypes() (vTypes []string) {
	return []string{"text", "longtext", "password", "link", "email", "phone", "date", "expdate", "time", "2fa"}
}
