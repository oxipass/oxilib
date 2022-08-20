package models

// OxiFieldTemplate - fields template definitions
type OxiFieldTemplate struct {
	ID        string `json:"field_id"`
	Name      string `json:"field_name"`
	Icon      string `json:"field_icon"`
	ValueType string `json:"value_type"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
	Deleted   bool   `json:"deleted"`
}
