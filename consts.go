package oxilib

// !!! WARNING !!!
// Never reduce the length of the database fields as
// it can lead to data loss, they can be increased only
// Every field change requires also upgrade procedure to be developed

// DatabaseIDLength - default database ID length
const DatabaseIDLength int = 64

const defaultDbVersion int = 1

const constZeroTime = "0000-00-00 00:00:00"

const constDbFormat = "2006-01-02 15:04:05"

// ConstSuccessResponse - default success response
const ConstSuccessResponse = "success"

// CErrorResponse - default error response
const CErrorResponse = "error"

const cTempDBFile = "test.sqlite"

const cLangsFolder = "langs"
