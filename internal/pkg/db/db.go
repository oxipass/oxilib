package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Needed to work correctly with db/sql
	"github.com/oxipass/oxilib/internal/pkg/oxierr"
	"os"
)

func (sdb *StorageDB) StartTX() (err error) {

	if sdb.sDB == nil {
		return oxierr.FormError(oxierr.BSERR00012DbTxStartFailed, "Database is not open")
	}
	if sdb.sTX != nil {
		return oxierr.FormError(oxierr.BSERR00012DbTxStartFailed, "Active transaction is already open, cannot open it twice, close  previous first")
	}
	sdb.sTX, err = sdb.sDB.Begin()
	if err != nil {
		return oxierr.FormError(oxierr.BSERR00012DbTxStartFailed, err.Error())
	}

	return nil
}

func (sdb *StorageDB) CommitTX() (err error) {
	return sdb.EndTransaction(true)
}

func (sdb *StorageDB) RollbackTX() (err error) {
	return sdb.EndTransaction(false)
}

func (sdb *StorageDB) EndTransaction(commit bool) (err error) {
	if sdb.sTX == nil {
		return oxierr.FormError(oxierr.BSERR00013DbTxEndFailed)
	}
	if commit {
		err = sdb.sTX.Commit()
	} else {
		err = sdb.sTX.Rollback()
	}
	sdb.sTX = nil
	return err
}

const cDBOpenParms = "?_txlock=immediate"

// Open method opens db in the provided file
func (sdb *StorageDB) Open(filePath string) error {
	var err error
	sdb.dbOpen = false

	if _, err = os.Stat(filePath); err == nil {
		// File exists, open and check the integrity
		sdb.sDB, err = sql.Open("sqlite3", "file:"+filePath+cDBOpenParms)
		if err != nil {
			return err
		}
		err = sdb.GetSettings()
		if err != nil {
			return err
		}
		sdb.dbOpen = true
		return nil
	}

	// File does not exist, create it
	sdb.sDB, err = sql.Open("sqlite3", "file:"+filePath+cDBOpenParms)
	if err != nil {
		return err
	}
	sdb.sDB.SetMaxOpenConns(1) // trying to remove db is locked issue

	err = sdb.StartTX()
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

func (sdb *StorageDB) Close() error {
	if !sdb.dbOpen {
		return nil
	}
	if sdb.sDB != nil {
		err := sdb.sDB.Close()
		if err != nil {
			return err
		}
		sdb.dbOpen = false
		sdb.sDB = nil
	}
	return nil
}

// IsOpen  - returns the flag if the db os already open
func (sdb *StorageDB) IsOpen() bool {
	return sdb.dbOpen
}

// Version - returns current version of DB
func (sdb *StorageDB) Version() int {
	return sdb.DbVersion
}

const constSelectTables = `
	SELECT name 
		FROM sqlite_master 
		WHERE type='table' AND name='settings' 
`

func (sdb *StorageDB) checkIntegrity() (bool, error) {
	rows, err := sdb.sDB.Query(constSelectTables)
	if err != nil {
		return false, oxierr.FormError(oxierr.BSERR00001DbIntegrityCheckFailed, err.Error())
	}

	foundSettings := false

	if rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			errRowsClose := rows.Close()
			if errRowsClose != nil {
				return false, oxierr.FormError(oxierr.BSERR00001DbIntegrityCheckFailed, err.Error(), errRowsClose.Error())
			}
			return false, oxierr.FormError(oxierr.BSERR00001DbIntegrityCheckFailed, err.Error())
		}

		if tableName == "settings" {
			foundSettings = true

		}
	}
	err = rows.Close()
	if err != nil {
		return false, oxierr.FormError(oxierr.BSERR00001DbIntegrityCheckFailed, err.Error())
	}
	return foundSettings, nil
}

const constSelectVersion = `
	SELECT database_version, database_id, crypt_id, language, keyword
		FROM settings LIMIT 1
`

func (sdb *StorageDB) GetSettings() error {
	foundSettings, err := sdb.checkIntegrity()
	if err != nil {
		return oxierr.FormError(oxierr.BSERR00001DbIntegrityCheckFailed, err.Error())
	}

	if !foundSettings {
		return oxierr.FormError(oxierr.BSERR00001DbIntegrityCheckFailed, "settings table is not found")
	}

	rowsSet, errSet := sdb.sDB.Query(constSelectVersion)
	if errSet != nil {
		return oxierr.FormError(oxierr.BSERR00001DbIntegrityCheckFailed, errSet.Error())
	}
	if rowsSet.Next() {
		errSet = rowsSet.Scan(&sdb.DbVersion, &sdb.DbID, &sdb.CryptID, &sdb.DbLang, &sdb.KeyWord)
		if errSet != nil {
			errSetClose := rowsSet.Close()
			if errSetClose != nil {
				return oxierr.FormError(oxierr.BSERR00001DbIntegrityCheckFailed, errSet.Error(), errSetClose.Error())
			}
			return oxierr.FormError(oxierr.BSERR00001DbIntegrityCheckFailed, errSet.Error())
		}
	}
	return rowsSet.Close()
}

func (sdb *StorageDB) initDb() (err error) {

	if sdb.sTX == nil {
		return oxierr.FormError(oxierr.BSERR00003DbTransactionFailed)
	}

	err = sdb.createTable(sqlCreateTableSettings)
	if err != nil {
		return err
	}

	err = sdb.createTable(sqlCreateTableItems)
	if err != nil {
		return err
	}

	err = sdb.createTable(sqlCreateTableFields)
	if err != nil {
		return err
	}

	err = sdb.createTable(sqlCreateTableTemplateItems)
	if err != nil {
		return err
	}

	err = sdb.createTable(sqlCreateTableTemplateFields)
	if err != nil {
		return err
	}

	err = sdb.createTable(sqlCreateTableTags)
	if err != nil {
		return err
	}

	err = sdb.createTable(sqlCreateTableItemsTags)
	if err != nil {
		return err
	}

	return nil
}
