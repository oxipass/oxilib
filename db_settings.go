package bslib

import (
	"fmt"
	"time"
)

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
