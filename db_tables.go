package bykovstorage

import "fmt"

const sqlCreateTableSettings = `
	CREATE TABLE settings (
   		database_id 		CHAR(%d) PRIMARY KEY NOT NULL,
   		keyword 			CHAR(128) NOT NULL,
		crypt_id            CHAR(8) NOT NULL,
		database_version 	INT,
   		update_timestamp    DATETIME,
		sync_timestamp		DATETIME
	)
`

func (sdb *storageDB) createTableSettings() error {
	sqlQuery := fmt.Sprintf(sqlCreateTableSettings, DatabaseIDLength)

	_, err := sdb.sTX.Exec(sqlQuery)
	if err != nil {
		return formError(BSERR00002DbTableCreationFailed, err.Error())
	}

	return nil
}

const sqlCreateTableItems = `
	CREATE TABLE items (
   		item_id 		    CHAR(%d) PRIMARY KEY NOT NULL,
   		name 			    VARCHAR(%d) NOT NULL,
		icon_id             VARCHAR(%d) NOT NULL,
		creation_timestamp  DATETIME,
   		update_timestamp    DATETIME,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0'
	)
`

func (sdb *storageDB) createTableItems() error {
	sqlQuery := fmt.Sprintf(sqlCreateTableItems,
		DatabaseItemIDLength,
		DatabaseItemNameLength,
		DatabaseIconIDLength)

	_, err := sdb.sTX.Exec(sqlQuery)
	if err != nil {
		return formError(BSERR00002DbTableCreationFailed, err.Error())
	}

	return nil
}
