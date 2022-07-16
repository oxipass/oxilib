package oxilib

import "time"

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

func (sdb *storageDB) dbInsertFieldTemplate(itemID int64, field OxiField) (fieldId int64, err error) {

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
