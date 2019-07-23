package bykovstorage

// BSItem - bykovstorage item structure
type BSItem struct {
	ID      string `json:"item_id"`
	Name    string `json:"item_name"`
	Icon    string `json:"item_icon"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Deleted bool   `json:"deleted"`
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

// JSONInputAddItem - input structure to add the item
type JSONInputAddItem struct {
	ItemName string `json:"item_name"`
	ItemIcon string `json:"item_icon"`
}

// JSONResponseAddItem - response structure for adding item
type JSONResponseAddItem struct {
	JSONResponseCommon
	ItemID string `json:"item_id"`
}

// JSONInputReadAll parameters for reading from database
type JSONInputReadAll struct {
	ReadDeleted bool `json:"read_deleted"`
}
