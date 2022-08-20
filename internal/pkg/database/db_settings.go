package database

import (
	"fmt"
	"github.com/oxipass/oxilib"
	"time"
)

type SettingInfo struct {
	DatabaseId string `json:"database_id"`
}

const sqlInsertInitialSettings = `
	INSERT INTO settings (database_id,keyword,crypt_id,language,database_version,update_timestamp,sync_timestamp)
		VALUES ('%s','%s','%s','%s','%d','%s','%s')
`

func (sdb *storageDB) initSettings(dbID, keyWord, cryptID, lang string) error {
	if sdb.sTX == nil {
		return oxilib.formError(oxilib.BSERR00003DbTransactionFailed, "initSettings")
	}

	lastUpdate := oxilib.prepareTimeForDb(time.Now())
	sqlQuery := fmt.Sprintf(sqlInsertInitialSettings, dbID, keyWord, cryptID, lang,
		oxilib.defaultDbVersion, lastUpdate, oxilib.constZeroTime)

	_, err := sdb.sTX.Exec(sqlQuery)

	if err != nil {
		return oxilib.formError(oxilib.BSERR00006DbInsertFailed, err.Error(), "initSettings")
	}

	return nil
}
