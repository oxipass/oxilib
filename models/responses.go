package models

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

// FieldAddedResponse - response structure for adding field
type FieldAddedResponse struct {
	CommonResponse
	FieldID int64 `json:"field_id"`
}

// TagAddedResponse - response structure for adding item
type TagAddedResponse struct {
	CommonResponse
	TagId int64 `json:"tag_id"`
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
