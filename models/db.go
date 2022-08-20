package models

import "database/sql"

// StorageDB is a class to access all storage functionality
type StorageDB struct {
	// instance & status
	sDB    *sql.DB
	dbOpen bool
	sTX    *sql.Tx

	// db settings
	dbVersion int
	dbID      string
	cryptID   string
	keyWord   string
	dbLang    string
}
