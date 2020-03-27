package bslib

import (
	"errors"
	"time"
)

const sqlCreateTableFields = `
	CREATE TABLE IF NOT EXISTS fields (
   		field_id 		    INTEGER PRIMARY KEY AUTOINCREMENT,
		item_id             INTEGER NOT NULL,
		field_icon          VARCHAR NOT NULL,
   		field_name 			VARCHAR NOT NULL,
		value_type          VARCHAR NOT NULL,
		field_value         VARCHAR NOT NULL,
		created  			DATETIME NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0',
		FOREIGN KEY (item_id) REFERENCES items(item_id)
	)
`

const sqlInsertField = `
	INSERT
		INTO fields (
			item_id,
			field_icon,
			field_name,
			value_type,
			field_value,
			created,
			updated,
			deleted)
		VALUES (?,?,?,?,?,?,?,0)
`

func (sdb *storageDB) dbInsertField(itemID int64, field BSField) (fieldId int64, err error) {

	if sdb.sTX == nil {
		return 0, formError(BSERR00003DbTransactionFailed, "dbInsertField")
	}

	creationTime := prepareTimeForDb(time.Now())

	stmt, err := sdb.sTX.Prepare(sqlInsertField)
	if err != nil {
		return 0, formError(BSERR00006DbInsertFailed, err.Error(), "dbInsertField")
	}
	res, errStmt := stmt.Exec(itemID,
		field.Icon,
		field.Name,
		field.ValueType,
		field.Value,
		creationTime,
		creationTime)

	if errStmt != nil {
		return 0, formError(BSERR00006DbInsertFailed, errStmt.Error(), "dbInsertField")
	}
	errClose := stmt.Close()
	if errClose != nil {
		return 0, formError(BSERR00006DbInsertFailed, errClose.Error(), "dbInsertField")
	}

	return res.LastInsertId()
}

// List all non-deleted items
const sqlListItemFields = `
	SELECT field_id, field_name, field_icon, field_value, value_type, created, updated, deleted
		FROM fields 
		WHERE deleted='0' and item_id=?
`

func (sdb *storageDB) dbSelectAllItemFields(itemId int64) (fields []BSField, err error) {

	rows, err := sdb.sDB.Query(sqlListItemFields, itemId)
	if err != nil {
		return fields, formError(BSERR00021FieldsReadFailed, err.Error())
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			err = formError(BSERR00021FieldsReadFailed, err.Error(), errClose.Error())
		}
	}()

	var bsField BSField

	for rows.Next() {
		err = rows.Scan(&bsField.ID,
			&bsField.Name,
			&bsField.Icon,
			&bsField.Value,
			&bsField.ValueType,
			&bsField.Created,
			&bsField.Updated,
			&bsField.Deleted)
		if err != nil {
			return fields, err
		}

		fields = append(fields, bsField)
	}
	return fields, nil
}

const sqlDeleteField = `UPDATE fields SET deleted=1, updated=? WHERE field_id=? `

func (sdb *storageDB) dbDeleteField(fieldID int64) (err error) {
	if sdb.sTX == nil {
		return formError(BSERR00003DbTransactionFailed, "dbDeleteField")
	}

	stmt, err := sdb.sTX.Prepare(sqlDeleteField)
	if err != nil {
		return err
	}

	updateTime := prepareTimeForDb(time.Now())

	_, err = stmt.Exec(updateTime, fieldID)
	if err != nil {
		return err
	}

	errClose := stmt.Close()
	if errClose != nil {
		return formError(BSERR00016DbDeleteFailed, errClose.Error())
	}
	return nil
}

// List all non-deleted items
const sqlGetField = `
	SELECT field_id, field_name, field_icon, field_value, value_type, created, updated, deleted
		FROM fields 
		WHERE  field_id=?
`

func (sdb *storageDB) dbGetFieldById(fieldId int64) (field BSField, err error) {

	rows, err := sdb.sDB.Query(sqlGetField, fieldId)
	if err != nil {
		return field, formError(BSERR00021FieldsReadFailed, err.Error())
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			err = formError(BSERR00021FieldsReadFailed, err.Error(), errClose.Error())
		}
	}()

	var bsField BSField

	if rows.Next() {
		err = rows.Scan(&bsField.ID,
			&bsField.Name,
			&bsField.Icon,
			&bsField.Value,
			&bsField.ValueType,
			&bsField.Created,
			&bsField.Updated,
			&bsField.Deleted)
		return bsField, err
	}
	return field, errors.New(BSERR00021FieldsReadFailed)
}
