package bslib

// BSItem - item structure
type BSItem struct {
	ID      int64     `json:"item_id"`
	Name    string    `json:"item_name"`
	Icon    string    `json:"item_icon"`
	Created string    `json:"created"`
	Updated string    `json:"updated"`
	Deleted bool      `json:"deleted"`
	Fields  []BSField `json:"fields"`
}

// BSField - fields definitions
type BSField struct {
	ID        string `json:"field_id"`
	Name      string `json:"field_name"`
	Icon      string `json:"field_icon"`
	ValueType string `json:"value_type"`
	Value     string `json:"field_value"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
	Deleted   bool   `json:"deleted"`
}

// JSONResponseItem - response returning one item
type JSONResponseItem struct {
	JSONResponseCommon
	BSItem
}

// JSONResponseItems - response returning many items
type JSONResponseItems struct {
	JSONResponseCommon
	Items []BSItem `json:"items"`
}

// JSONResponseCommon - common response header structure
type JSONResponseCommon struct {
	Status string `json:"status"`
	MsgNum string `json:"msg_num"`
	MsgTxt string `json:"msg_text"`
}

// JSONInputUpdateField input structure to add or update the field
type JSONInputUpdateField struct {
	ItemID int64 `json:"item_id"`
	BSField
}

// JSONResponseFieldAdded - response structure for adding field
type JSONResponseFieldAdded struct {
	JSONResponseCommon
	FieldID int64 `json:"field_id"`
}

// JSONInputUpdateItem - input structure to add the item
type JSONInputUpdateItem struct {
	ItemID   int64  `json:"item_id"`
	ItemName string `json:"item_name"`
	ItemIcon string `json:"item_icon"`
}

// JSONResponseItemAdded - response structure for adding item
type JSONResponseItemAdded struct {
	JSONResponseCommon
	ItemID int64 `json:"item_id"`
}

type JSONInputInitStorage struct {
	FileName   string `json:"filename"`
	Password   string `json:"password"`
	Encryption string `json:"encryption"`
}

// JSONInputReadAll parameters for reading from database
type JSONInputReadAll struct {
	ReadDeleted bool `json:"read_deleted"`
}
