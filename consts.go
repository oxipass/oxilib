package bykovstorage

// !!! WARNING !!!
// Never reduce the length of the database fields as
// it can lead to data loss, they can be increased only
// Every field change requires also upgrade procedure to be developed

// DatabaseIDLength - default database ID length
const DatabaseIDLength int = 64

// DatabaseItemIDLength - default item ID length
const DatabaseItemIDLength int = 8

// DatabaseItemNameLength - default item name length
const DatabaseItemNameLength int = 512

// DatabaseIconIDLength - default icon ID (name) length
const DatabaseIconIDLength int = 32

const defaultDbVersion int = 1

const constZeroTime = "0000-00-00 00:00:00"

const constDbFormat = "2006-01-02 15:04:05"

// ConstSuccessResponse - default success response
const ConstSuccessResponse = "success"

// ConstErrorResponse - default error response
const ConstErrorResponse = "error"
