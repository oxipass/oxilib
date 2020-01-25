package bslib

// BSItem - bykovstorage item structure
type BSItem struct {
	ID      string `json:"item_id"`
	Name    string `json:"item_name"`
	Icon    string `json:"item_icon"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Deleted bool   `json:"deleted"`
}

// BSFieldDefinition - fields definitions
type BSFieldDefinition struct {
	ID        string `json:"field_type_id"`
	Name      string `json:"name"`
	Icon      string `json:"icon_id"`
	ValueType string `json:"value_type"`
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

// JSONInputUpdateItem - input structure to add the item
type JSONInputUpdateItem struct {
	ItemID   string `json:"item_id"`
	ItemName string `json:"item_name"`
	ItemIcon string `json:"item_icon"`
}

// JSONResponseItemAdded - response structure for adding item
type JSONResponseItemAdded struct {
	JSONResponseCommon
	ItemID string `json:"item_id"`
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
