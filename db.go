package bykovstorage

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3" // Needed to work correctly with database/sql
)

// StorageDB is a class to access all storage functionality
type storageDB struct {
	// instance & status
	sDB    *sql.DB
	dbOpen bool
	sTX    *sql.Tx

	// db settings
	dbVersion int
	dbID      string
	cryptID   string
	keyWord   string
}

func (sdb *storageDB) StartTransaction() (err error) {

	if sdb.sDB == nil {
		return formError(BSERR00012DbTxStartFailed, "Database is not open")
	}
	if sdb.sTX != nil {
		return formError(BSERR00012DbTxStartFailed, "Active transaction is already open, cannot open it twice, close  previous first")
	}
	sdb.sTX, err = sdb.sDB.Begin()
	if err != nil {
		return formError(BSERR00012DbTxStartFailed, err.Error())
	}

	return nil
}

func (sdb *storageDB) EndTransaction(commit bool) (err error) {
	if sdb.sTX == nil {
		return formError(BSERR00013DbTxEndFaild)
	}
	if commit {
		err = sdb.sTX.Commit()
	} else {
		err = sdb.sTX.Rollback()
	}
	sdb.sTX = nil
	return err
}

// Open method opens database in the provided file
func (sdb *storageDB) Open(filePath string) error {
	var err error
	sdb.dbOpen = false

	if _, err = os.Stat(filePath); err == nil {
		// File exists, open and check integrity
		sdb.sDB, err = sql.Open("sqlite3", filePath)
		if err != nil {
			return err
		}
		err = sdb.checkIntegrityGetSettings()
		if err != nil {
			return err
		}
		sdb.dbOpen = true
		return nil
	}

	// File does not exists
	sdb.sDB, err = sql.Open("sqlite3", filePath)
	if err != nil {
		return err
	}

	err = sdb.StartTransaction()
	if err != nil {
		return err
	}

	err = sdb.initDb()
	if err != nil {
		return err
	}

	err = sdb.EndTransaction(true)
	if err != nil {
		return err
	}

	sdb.dbOpen = true
	return nil
}

// IsOpen  - returns the flag if the database os already open
func (sdb *storageDB) IsOpen() bool {
	return sdb.dbOpen
}

// Version - returns current version of DB
func (sdb *storageDB) Version() int {
	return sdb.dbVersion
}

const constSelectTables = `
	SELECT name 
		FROM sqlite_master 
		WHERE type='table' AND name='settings' 
`

const constSelectVersion = `
	SELECT database_version, database_id, crypt_id, keyword
		FROM settings LIMIT 1
`

func (sdb *storageDB) checkIntegrityGetSettings() error {

	rows, err := sdb.sDB.Query(constSelectTables)
	if err != nil {
		return formError(BSERR00001DbIntegrityCheckFailed, err.Error())
	}
	defer rows.Close()
	if rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			return formError(BSERR00001DbIntegrityCheckFailed, err.Error())
		}

		if tableName == "settings" {
			rowsSet, errSet := sdb.sDB.Query(constSelectVersion)
			if errSet != nil {
				return formError(BSERR00001DbIntegrityCheckFailed, errSet.Error())
			}
			if rowsSet.Next() {
				errSet = rowsSet.Scan(&sdb.dbVersion, &sdb.dbID, &sdb.cryptID, &sdb.keyWord)
				if errSet != nil {
					return formError(BSERR00001DbIntegrityCheckFailed, errSet.Error())
				}
			}
		}
	}
	return nil
}

func (sdb *storageDB) initDb() (err error) {

	if sdb.sTX == nil {
		return formError(BSERR00003DbTransactionFailed)
	}

	err = sdb.createTableSettings()
	if err != nil {
		return err
	}

	err = sdb.createTableItems()
	if err != nil {
		return err
	}

	return nil
}

const sqlInsertInitialSettings = `
	INSERT INTO settings (database_id,keyword,crypt_id,database_version,update_timestamp,sync_timestamp)
		VALUES ('%s','%s','%s','%d','%s','%s')
`

func (sdb *storageDB) initSettings(dbID, keyWord, cryptID string) error {
	if sdb.sTX == nil {
		return formError(BSERR00003DbTransactionFailed, "initSettings")
	}

	lastUpdate := prepareTimeForDb(time.Now())
	sqlQuery := fmt.Sprintf(sqlInsertInitialSettings, dbID, keyWord, cryptID,
		defaultDbVersion, lastUpdate, constZeroTime)

	_, err := sdb.sTX.Exec(sqlQuery)

	if err != nil {
		return formError(BSERR00006DbInsertFailed, err.Error(), "initSettings")
	}

	return nil
}
