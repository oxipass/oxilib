package database

import (
	"errors"
	"github.com/oxipass/oxilib"
	"github.com/oxipass/oxilib/models"
	"time"
)

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

func (sdb *storageDB) dbInsertField(itemID int64, field models.OxiField) (fieldId int64, err error) {

	if sdb.sTX == nil {
		return 0, oxilib.formError(oxilib.BSERR00003DbTransactionFailed, "dbInsertField")
	}

	creationTime := oxilib.prepareTimeForDb(time.Now())

	stmt, err := sdb.sTX.Prepare(sqlInsertField)
	if err != nil {
		return 0, oxilib.formError(oxilib.BSERR00006DbInsertFailed, err.Error(), "dbInsertField")
	}
	res, errStmt := stmt.Exec(itemID,
		field.Icon,
		field.Name,
		field.ValueType,
		field.Value,
		creationTime,
		creationTime)

	if errStmt != nil {
		return 0, oxilib.formError(oxilib.BSERR00006DbInsertFailed, errStmt.Error(), "dbInsertField")
	}
	errClose := stmt.Close()
	if errClose != nil {
		return 0, oxilib.formError(oxilib.BSERR00006DbInsertFailed, errClose.Error(), "dbInsertField")
	}

	return res.LastInsertId()
}

// List all non-deleted items
const sqlListItemFields = `
	SELECT field_id, field_name, field_icon, field_value, value_type, created, updated, deleted
		FROM fields 
		WHERE deleted='0' and item_id=?
`

func (sdb *storageDB) dbSelectAllItemFields(itemId int64) (fields []models.OxiField, err error) {

	rows, err := sdb.sDB.Query(sqlListItemFields, itemId)
	if err != nil {
		return fields, oxilib.formError(oxilib.BSERR00021FieldsReadFailed, err.Error())
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			err = oxilib.formError(oxilib.BSERR00021FieldsReadFailed, err.Error(), errClose.Error())
		}
	}()

	var field models.OxiField

	for rows.Next() {
		err = rows.Scan(&field.ID,
			&field.Name,
			&field.Icon,
			&field.Value,
			&field.ValueType,
			&field.Created,
			&field.Updated,
			&field.Deleted)
		if err != nil {
			return fields, err
		}

		fields = append(fields, field)
	}
	return fields, nil
}

const sqlDeleteField = `UPDATE fields SET deleted=1, updated=? WHERE field_id=? `

func (sdb *storageDB) dbDeleteField(fieldID int64) (err error) {
	if sdb.sTX == nil {
		return oxilib.formError(oxilib.BSERR00003DbTransactionFailed, "dbDeleteField")
	}

	stmt, err := sdb.sTX.Prepare(sqlDeleteField)
	if err != nil {
		return err
	}

	updateTime := oxilib.prepareTimeForDb(time.Now())

	_, err = stmt.Exec(updateTime, fieldID)
	if err != nil {
		return err
	}

	errClose := stmt.Close()
	if errClose != nil {
		return oxilib.formError(oxilib.BSERR00016DbDeleteFailed, errClose.Error())
	}
	return nil
}

// List all non-deleted items
const sqlGetField = `
	SELECT field_id, field_name, field_icon, field_value, value_type, created, updated, deleted
		FROM fields 
		WHERE  field_id=?
`

func (sdb *storageDB) dbGetFieldById(fieldId int64) (field models.OxiField, err error) {

	rows, err := sdb.sDB.Query(sqlGetField, fieldId)
	if err != nil {
		return field, oxilib.formError(oxilib.BSERR00021FieldsReadFailed, err.Error())
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			err = oxilib.formError(oxilib.BSERR00021FieldsReadFailed, err.Error(), errClose.Error())
		}
	}()

	if rows.Next() {
		err = rows.Scan(&field.ID,
			&field.Name,
			&field.Icon,
			&field.Value,
			&field.ValueType,
			&field.Created,
			&field.Updated,
			&field.Deleted)
		return field, err
	}
	return field, errors.New(oxilib.BSERR00021FieldsReadFailed)
}
