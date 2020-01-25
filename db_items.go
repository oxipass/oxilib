package bslib

import (
	"time"
)

const sqlCreateTableItems = `
	CREATE TABLE items (
   		item_id 		    CHAR PRIMARY KEY NOT NULL,
   		name 			    VARCHAR NOT NULL,
		icon             	VARCHAR NOT NULL,
		created  			DATETIME NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0'
	)
`

const sqlDeleteItem = `UPDATE items SET deleted=1, updated=? WHERE item_id=? `

func (sdb *storageDB) dbDeleteItem(itemID string) (err error) {
	if sdb.sTX == nil {
		return formError(BSERR00003DbTransactionFailed, "dbDeleteItem")
	}

	stmt, err := sdb.sTX.Prepare(sqlDeleteItem)
	if err != nil {
		return err
	}

	updateTime := prepareTimeForDb(time.Now())

	_, err = stmt.Exec(updateTime, itemID)
	if err != nil {
		return err
	}

	errClose := stmt.Close()
	if errClose != nil {
		err = formError(BSERR00016DbDeleteFailed, errClose.Error())
	}
	return nil
}

const sqlInsertItem = `
	INSERT 
		INTO items (item_id,name,icon,created,updated,deleted) 
		VALUES (?,?,?,?,?, 0)
`

func (sdb *storageDB) dbInsertItem(itemName string, itemIcon string) (itemID string, err error) {
	if sdb.sTX == nil {
		return "", formError(BSERR00003DbTransactionFailed, "dbInsertItem")
	}

	creationTime := prepareTimeForDb(time.Now())
	itemID = generateRandomString(8)

	stmt, err := sdb.sTX.Prepare(sqlInsertItem)
	if err != nil {
		return "", formError(BSERR00006DbInsertFailed, err.Error(), "dbInsertItem")
	}
	_, errStmt := stmt.Exec(itemID,
		itemName,
		itemIcon,
		creationTime,
		creationTime)

	if errStmt != nil {
		return "", formError(BSERR00006DbInsertFailed, errStmt.Error(), "dbInsertItem")
	}
	errClose := stmt.Close()
	if errClose != nil {
		return "", formError(BSERR00006DbInsertFailed, errClose.Error(), "dbInsertItem")
	}

	return itemID, nil
}

// List all non-deleted items
const sqlListItems = `
	SELECT item_id, name, icon, created, updated, deleted
		FROM items 
		WHERE deleted='0'
`

func (sdb *storageDB) dbSelectAllItems() (items []BSItem, err error) {

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

const sqlUpdateItemName = `UPDATE items SET name=?, updated=? WHERE item_id=? `

func (sdb *storageDB) dbUpdateItemName(itemID string, newName string) (err error) {
	if sdb.sTX == nil {
		return formError(BSERR00003DbTransactionFailed, "dbUpdateItemName")
	}

	stmt, err := sdb.sTX.Prepare(sqlUpdateItemName)
	if err != nil {
		return err
	}

	updateTime := prepareTimeForDb(time.Now())

	_, err = stmt.Exec(newName, updateTime, itemID)
	if err != nil {
		return err
	}

	errClose := stmt.Close()
	if errClose != nil {
		err = formError(BSERR00018DbItemNameUpdateFailed, errClose.Error())
	}
	return nil
}

// List all non-deleted items
const sqlGetItemById = `
	SELECT item_id, name, icon, created, updated, deleted
		FROM items 
		WHERE item_id=?
`

func (sdb *storageDB) dbGetItemById(itemId string) (item BSItem, err error) {
	stmt, err := sdb.sDB.Prepare(sqlGetItemById)
	if err != nil {
		return item, formError(BSERR00014ItemsReadFailed, err.Error())
	}
	rows, err := stmt.Query(itemId)
	if err != nil {
		return item, formError(BSERR00014ItemsReadFailed, err.Error())
	}

	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			err = formError(BSERR00014ItemsReadFailed, err.Error(), errClose.Error())
		}
		errClose = stmt.Close()
		if errClose != nil {
			err = formError(BSERR00014ItemsReadFailed, err.Error(), errClose.Error())
		}
	}()

	if rows.Next() {
		var bsItem BSItem
		err = rows.Scan(&bsItem.ID,
			&bsItem.Name,
			&bsItem.Icon,
			&bsItem.Created,
			&bsItem.Updated,
			&bsItem.Deleted)
		if err != nil {
			return item, err
		}
		return bsItem, nil
	}
	return item, formError(BSERR00019ItemNotFound)
}
