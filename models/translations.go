package models

type Lang struct {
	Code       string   `json:"code"`
	Locales    []string `json:"locales"`
	Name       string   `json:"name"`
	NativeName string   `json:"native"`
}
type Translations struct {
	Lang
	Translations map[string]string `json:"translations"`
}
