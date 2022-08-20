package models

// OxiItemTemplate - item templatestructure
type OxiItemTemplate struct {
	ID      string             `json:"item_id"`
	Name    string             `json:"item_name"`
	Icon    string             `json:"item_icon"`
	Created string             `json:"created"`
	Updated string             `json:"updated"`
	Deleted bool               `json:"deleted"`
	Fields  []OxiFieldTemplate `json:"fields"`
}
