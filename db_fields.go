package bslib

import "time"

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

func (sdb *storageDB) dbInsertField(fieldDefinition BSField) (err error) {
	if sdb.sTX == nil {
		return formError(BSERR00003DbTransactionFailed, "dbInsertFieldDefinition")
	}

	creationTime := prepareTimeForDb(time.Now())

	stmt, err := sdb.sTX.Prepare(sqlInsertField)
	if err != nil {
		return formError(BSERR00006DbInsertFailed, err.Error(), "dbInsertFieldDefinition")
	}
	_, errStmt := stmt.Exec(fieldDefinition.ID,
		fieldDefinition.Name,
		fieldDefinition.Icon,
		fieldDefinition.ValueType,
		creationTime,
		creationTime)

	if errStmt != nil {
		return formError(BSERR00006DbInsertFailed, errStmt.Error(), "dbInsertFieldDefinition")
	}
	errClose := stmt.Close()
	if errClose != nil {
		return formError(BSERR00006DbInsertFailed, errClose.Error(), "dbInsertFieldDefinition")
	}

	return nil
}
