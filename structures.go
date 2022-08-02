package oxilib

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

// OxiField - fields definitions
type OxiField struct {
	ID        int64  `json:"field_id"`
	Name      string `json:"field_name"`
	Icon      string `json:"field_icon"`
	ValueType string `json:"value_type"`
	Value     string `json:"field_value"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
	Deleted   bool   `json:"deleted"`
}

// OxiTag - tags definitions
type OxiTag struct {
	ID      int64  `json:"tag_id"`
	Name    string `json:"tag_name"`
	Color   string `json:"color"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Deleted bool   `json:"deleted"`
}

// CommonResponse - common response header structure
type CommonResponse struct {
	Status string `json:"status"`
	MsgNum string `json:"msg_num"`
	MsgTxt string `json:"msg_text"`
}

// ItemResponse - response returning one item
type ItemResponse struct {
	CommonResponse
	OxiItem
}

// ItemsResponse - response returning many items
type ItemsResponse struct {
	CommonResponse
	Items []OxiItem `json:"items"`
}

// UpdateFieldForm - input structure to add or update the field
type UpdateFieldForm struct {
	ItemID int64 `json:"item_id"`
	OxiField
}

// FieldAddedResponse - response structure for adding field
type FieldAddedResponse struct {
	CommonResponse
	FieldID int64 `json:"field_id"`
}

// UpdateItemForm - input structure to add the item
type UpdateItemForm struct {
	OxiItem
}

// TagAddedResponse - response structure for adding item
type TagAddedResponse struct {
	CommonResponse
	TagId int64 `json:"tag_id"`
}

type UpdateTagForm struct {
	ItemID int64 `json:"item_id"`
	OxiTag
}

type TagAssignedResponse struct {
	ItemTagId int64 `json:"item_tag_id"`
	CommonResponse
}

// ItemAddedResponse - response structure for adding item
type ItemAddedResponse struct {
	CommonResponse
	ItemID int64 `json:"item_id"`
}

type ItemUpdatedResponse struct {
	CommonResponse
	UpdatedFields string `json:"updated_fields"`
}

// InitStorageForm - initializing the database
type InitStorageForm struct {
	FileName   string `json:"filename"`
	Password   string `json:"password"`
	Encryption string `json:"encryption"`
}

// ReadAllForm parameters for reading from database
type ReadAllForm struct {
	ReadDeleted bool `json:"read_deleted"`
}

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

type Tag struct {
	ID    string `json:"id"`
	Color string `json:"color"`
}

type TagsTemplate struct {
	Updated string `json:"updated"`
	Tags    []Tag  `json:"tags"`
}
