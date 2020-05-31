package bslib

import "time"

const sqlCreateTableTemplateItems = `
	CREATE TABLE IF NOT EXISTS template_items (
   		item_id 		    INTEGER PRIMARY KEY AUTOINCREMENT,
   		name 			    VARCHAR NOT NULL,
		icon             	VARCHAR NOT NULL,
		created  			DATETIME NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0'
	)
`

const sqlCreateTableTemplateFields = `
	CREATE TABLE IF NOT EXISTS template_fields (
   		field_id 		    INTEGER PRIMARY KEY AUTOINCREMENT,
		item_id             INTEGER NOT NULL,
		field_icon          VARCHAR NOT NULL,
   		field_name 			VARCHAR NOT NULL,
		value_type          VARCHAR NOT NULL,
		created  			DATETIME NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0',
		FOREIGN KEY (item_id) REFERENCES templates_items(item_id)
	)
`

const sqlInsertItemTemplate = `
	INSERT 
		INTO template_items (name,icon,created,updated,deleted) 
		VALUES (?,?,?,?,0)
`

func (sdb *storageDB) dbInsertItemTemplate(itemName string, itemIcon string) (itemID int64, err error) {
	if sdb.sTX == nil {
		return 0, formError(BSERR00003DbTransactionFailed, "dbInsertItemTemplate")
	}

	creationTime := prepareTimeForDb(time.Now())

	stmt, err := sdb.sTX.Prepare(sqlInsertItemTemplate)
	if err != nil {
		return 0, formError(BSERR00006DbInsertFailed, err.Error(), "dbInsertItemTemplate")
	}
	res, errStmt := stmt.Exec(itemName,
		itemIcon,
		creationTime,
		creationTime)

	if errStmt != nil {
		return 0, formError(BSERR00006DbInsertFailed, errStmt.Error(), "dbInsertItemTemplate")
	}
	itemID, err = res.LastInsertId()
	if err != nil {
		return 0, formError(BSERR00006DbInsertFailed, err.Error(), "dbInsertItemTemplate")
	}
	errClose := stmt.Close()
	if errClose != nil {
		return 0, formError(BSERR00006DbInsertFailed, errClose.Error(), "dbInsertItemTemplate")
	}

	return itemID, nil
}

const sqlInsertFieldTemplate = `
	INSERT
		INTO template_fields (
			item_id,
			field_icon,
			field_name,
			value_type,
			created,
			updated,
			deleted)
		VALUES (?,?,?,?,?,?,0)
`

func (sdb *storageDB) dbInsertFieldTemplate(itemID int64, field BSField) (fieldId int64, err error) {

	if sdb.sTX == nil {
		return 0, formError(BSERR00003DbTransactionFailed, "dbInsertFieldTemplate")
	}

	creationTime := prepareTimeForDb(time.Now())

	stmt, err := sdb.sTX.Prepare(sqlInsertFieldTemplate)
	if err != nil {
		return 0, formError(BSERR00006DbInsertFailed, err.Error(), "dbInsertFieldTemplate")
	}
	res, errStmt := stmt.Exec(itemID,
		field.Icon,
		field.Name,
		field.ValueType,
		creationTime,
		creationTime)

	if errStmt != nil {
		return 0, formError(BSERR00006DbInsertFailed, errStmt.Error(), "dbInsertFieldTemplate")
	}
	errClose := stmt.Close()
	if errClose != nil {
		return 0, formError(BSERR00006DbInsertFailed, errClose.Error(), "dbInsertFieldTemplate")
	}

	return res.LastInsertId()
}
