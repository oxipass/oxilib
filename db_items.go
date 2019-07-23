package bykovstorage

import (
	"time"
)

const sqlInsertItem = `INSERT INTO items (item_id,name,icon_id,creation_timestamp,update_timestamp,deleted) VALUES (?,?,?,?,?, 0)
`

func (sdb *storageDB) insertItem(itemName string, itemIcon string) (itemID string, err error) {
	if sdb.sTX == nil {
		return "", formError(BSERR00003DbTransactionFailed, "insertItem")
	}

	creationTime := prepareTimeForDb(time.Now())
	itemID = generateRandomString(8)

	stmt, err := sdb.sTX.Prepare(sqlInsertItem)
	if err != nil {
		return "", formError(BSERR00006DbInsertFailed, err.Error(), "insertItem")
	}
	_, errStmt := stmt.Exec(itemID,
		itemName,
		itemIcon,
		creationTime,
		creationTime)

	if errStmt != nil {
		return "", formError(BSERR00006DbInsertFailed, errStmt.Error(), "insertItem")
	}
	errClose := stmt.Close()
	if errClose != nil {
		return "", formError(BSERR00006DbInsertFailed, errClose.Error(), "insertItem")
	}

	return itemID, nil
}

// creation_timestamp, update_timestamp,
const sqlListItems = `
	SELECT item_id, name, icon_id, creation_timestamp, update_timestamp, deleted
		FROM items 
		WHERE deleted='0'
`

func (sdb *storageDB) selectAllItems() (items []BSItem, err error) {
	rows, err := sdb.sDB.Query(sqlListItems)
	if err != nil {
		return items, formError(BSERR00014ItemsReadFailed, err.Error())
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			err = formError(BSERR00014ItemsReadFailed, err.Error(), errClose.Error())
		}

	}()

	var bsItem BSItem

	for rows.Next() {
		err = rows.Scan(&bsItem.ID,
			&bsItem.Name,
			&bsItem.Icon,
			&bsItem.Created,
			&bsItem.Updated,
			&bsItem.Deleted)
		if err != nil {
			return items, err
		}

		items = append(items, bsItem)
	}
	return items, nil
}
