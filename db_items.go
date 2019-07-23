package bykovstorage

import (
	"fmt"
	"log"
	"time"
)

const sqlInsertItem = `
	INSERT INTO items (item_id,name,icon_id,creation_timestamp,update_timestamp,deleted)
		VALUES ('%s','%s','%s','%s','%s', 0)
`

func (sdb *storageDB) insertItem(itemName string, itemIcon string) (itemID string, err error) {
	if sdb.sTX == nil {
		return "", formError(BSERR00003DbTransactionFailed, "insertItem")
	}

	creationTime := prepareTimeForDb(time.Now())
	itemID = generateRandomString(8)

	sqlQuery := fmt.Sprintf(sqlInsertItem,
		itemID,
		itemName,
		itemIcon,
		creationTime,
		creationTime)

	log.Print(sqlQuery)

	_, err = sdb.sTX.Exec(sqlQuery) // FIXME: remove it after debugging

	if err != nil {
		return itemID, formError(BSERR00006DbInsertFailed, err.Error(), "insertItem")
	}

	return itemID, nil
}

// creation_timestamp, update_timestamp,
const sqlListItems = `
	SELECT item_id, name, icon_id, deleted
		FROM items 
		WHERE deleted='0'
`

func (sdb *storageDB) selectAllItems() (items []BSItem, err error) {
	rows, err := sdb.sDB.Query(sqlListItems)
	if err != nil {
		return items, formError(BSERR00014ItemsReadFailed, err.Error())
	}
	defer rows.Close()

	var bsItem BSItem

	if rows.Next() {
		err = rows.Scan(&bsItem.ID,
			&bsItem.Name,
			&bsItem.Icon,
			//&bsItem.Created,
			//&bsItem.Updated,
			&bsItem.Deleted)
		if err != nil {
			return items, err
		}

		items = append(items, bsItem)
	}
	return items, nil
}
