package oxilib

import (
	"fmt"
	"time"
)

type SettingInfo struct {
	DatabaseId string `json:"database_id"`
}

const sqlCreateTableSettings = `
	CREATE TABLE IF NOT EXISTS settings (
   		database_id 		CHAR PRIMARY KEY NOT NULL,
   		keyword 			CHAR NOT NULL,
		crypt_id            CHAR NOT NULL,
		database_version 	INT NOT NULL,
   		update_timestamp    DATETIME NOT NULL,
		sync_timestamp		DATETIME NOT NULL
	)
`

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
