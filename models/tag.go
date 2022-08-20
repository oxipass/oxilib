package models

// OxiTag - tags definitions
type OxiTag struct {
	ID      int64  `json:"tag_id"`
	Name    string `json:"tag_name"`
	Color   string `json:"color"`
	ExtId   string `json:"extid"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Deleted bool   `json:"deleted"`
}
