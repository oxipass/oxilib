package db

import "database/sql"

// StorageDB is a class to access all storage functionality
type StorageDB struct {
	// instance & status
	sDB    *sql.DB
	dbOpen bool
	sTX    *sql.Tx

	// db settings
	DbVersion int
	DbID      string
	CryptID   string
	KeyWord   string
	DbLang    string
}
