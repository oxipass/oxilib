package models

// OxiItem - item structure
type OxiItem struct {
	ID      int64      `json:"item_id"`
	Name    string     `json:"item_name"`
	Icon    string     `json:"item_icon"`
	Created string     `json:"created"`
	Updated string     `json:"updated"`
	Deleted bool       `json:"deleted"`
	Fields  []OxiField `json:"fields"`
	Tags    []OxiTag   `json:"tags"`
}
