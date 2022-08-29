package consts

// !!! WARNING !!!
// Never reduce the length of the db fields as
// it can lead to data loss, they can be increased only
// Every field change requires also upgrade procedure to be developed

// DatabaseIDLength - default db ID length
const DatabaseIDLength int = 64

const CDbVersion int = 1

const CZeroTime = "0000-00-00 00:00:00"

const CDbFormat = "2006-01-02 15:04:05"

// CSuccessResponse - default success response
const CSuccessResponse = "success"

// CErrorResponse - default error response
const CErrorResponse = "error"

const CTempDBFile = "test.sqlite"

// CLangsFolder - folder with translated language files
const CLangsFolder = "langs"

// CTemplatesFolder - items and fields templates folder
const CTemplatesFolder = "templates"

const CTagsTemplates = "tags.json"

const CFieldsTemplates = "fields.json"

const CItemsTemplates = "items.json"

const CIconDefaultItem = "solid/file"
