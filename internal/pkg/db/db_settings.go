package db

import (
	"fmt"
	"github.com/oxipass/oxilib/consts"
	"github.com/oxipass/oxilib/internal/pkg/oxierr"
	"github.com/oxipass/oxilib/internal/pkg/utils"
	"time"
)

type SettingInfo struct {
	DatabaseId string `json:"database_id"`
}

const sqlInsertInitialSettings = `
	INSERT INTO settings (database_id,keyword,crypt_id,language,database_version,update_timestamp,sync_timestamp)
		VALUES ('%s','%s','%s','%s','%d','%s','%s')
`

func (sdb *StorageDB) InitSettings(dbID, keyWord, cryptID, lang string) error {
	if sdb.sTX == nil {
		return oxierr.FormError(oxierr.BSERR00003DbTransactionFailed, "initSettings")
	}

	lastUpdate := utils.PrepareTimeForDb(time.Now())
	sqlQuery := fmt.Sprintf(sqlInsertInitialSettings, dbID, keyWord, cryptID, lang,
		consts.CDbVersion, lastUpdate, consts.CZeroTime)

	_, err := sdb.sTX.Exec(sqlQuery)

	if err != nil {
		return oxierr.FormError(oxierr.BSERR00006DbInsertFailed, err.Error(), "initSettings")
	}

	return nil
}
