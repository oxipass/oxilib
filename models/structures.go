package models

// UpdateFieldForm - input structure to add or update the field
type UpdateFieldForm struct {
	ItemID int64 `json:"item_id"`
	OxiField
}

// UpdateItemForm - input structure to add the item
type UpdateItemForm struct {
	OxiItem
}

type UpdateTagForm struct {
	ItemID int64 `json:"item_id"`
	OxiTag
}

// InitStorageForm - initializing the db
type InitStorageForm struct {
	FileName   string `json:"filename"`
	Password   string `json:"password"`
	Encryption string `json:"encryption"`
}

// ReadAllForm parameters for reading from db
type ReadAllForm struct {
	ReadDeleted bool `json:"read_deleted"`
}

type TagTemplateJSON struct {
	ID    string `json:"id"`
	Color string `json:"color"`
}

type TagsTemplateJSON struct {
	Updated string            `json:"updated"`
	Tags    []TagTemplateJSON `json:"tags"`
}

type FieldTemplateJSON struct {
	ID        string `json:"id"`
	FieldType string `json:"type"`
	Icon      string `json:"icon"`
}

type FieldsTemplateJSON struct {
	Updated string              `json:"updated"`
	Fields  []FieldTemplateJSON `json:"fields"`
}

type ItemTemplateJSON struct {
	ID        string   `json:"id"`
	Icon      string   `json:"icon"`
	FieldsIds []string `json:"fields"`
}

type ItemsTemplateJSON struct {
	Updated string             `json:"updated"`
	Items   []ItemTemplateJSON `json:"items"`
}
