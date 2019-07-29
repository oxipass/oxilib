package bykovstorage

import "fmt"

const sqlCreateTableSettings = `
	CREATE TABLE settings (
   		database_id 		CHAR PRIMARY KEY NOT NULL,
   		keyword 			CHAR NOT NULL,
		crypt_id            CHAR NOT NULL,
		database_version 	INT NOT NULL,
   		update_timestamp    DATETIME NOT NULL,
		sync_timestamp		DATETIME NOT NULL
	)
`

func (sdb *storageDB) createTableSettings() error {
	sqlQuery := fmt.Sprintf(sqlCreateTableSettings)

	_, err := sdb.sTX.Exec(sqlQuery)
	if err != nil {
		return formError(BSERR00002DbTableCreationFailed, err.Error())
	}

	return nil
}

const sqlCreateTableItems = `
	CREATE TABLE items (
   		item_id 		    CHAR PRIMARY KEY NOT NULL,
   		name 			    VARCHAR NOT NULL,
		icon_id             VARCHAR NOT NULL,
		creation_timestamp  DATETIME NOT NULL,
   		update_timestamp    DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0'
	)
`

func (sdb *storageDB) createTableItems() error {

	_, err := sdb.sTX.Exec(sqlCreateTableItems)
	if err != nil {
		return formError(BSERR00002DbTableCreationFailed, err.Error())
	}

	return nil
}
