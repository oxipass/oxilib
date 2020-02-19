package bslib

//
//const sqlCreateTableFieldsDefinitions = `
//	CREATE TABLE IF NOT EXISTS fields_definitions (
//   		field_def_id 		CHAR PRIMARY KEY NOT NULL,
//   		name 			    VARCHAR NOT NULL,
//		icon             	VARCHAR NOT NULL,
//		value_type          VARCHAR NOT NULL,
//		created  			DATETIME NOT NULL,
//   		updated  			DATETIME NOT NULL,
//		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0'
//	)
//`
//
//const sqlInsertFieldDefinition = `
//	INSERT
//		INTO fields_definitions (
//			field_def_id,
//			name,
//			icon,
//			value_type,
//			created,
//			updated,
//			deleted)
//		VALUES (?,?,?,?,?,?,0)
//`
//
//func (sdb *storageDB) dbInsertFieldDefinition(fieldDefinition BSFieldDefinition) (err error) {
//	if sdb.sTX == nil {
//		return formError(BSERR00003DbTransactionFailed, "dbInsertFieldDefinition")
//	}
//
//	creationTime := prepareTimeForDb(time.Now())
//
//	stmt, err := sdb.sTX.Prepare(sqlInsertFieldDefinition)
//	if err != nil {
//		return formError(BSERR00006DbInsertFailed, err.Error(), "dbInsertFieldDefinition")
//	}
//	_, errStmt := stmt.Exec(fieldDefinition.ID,
//		fieldDefinition.Name,
//		fieldDefinition.Icon,
//		fieldDefinition.ValueType,
//		creationTime,
//		creationTime)
//
//	if errStmt != nil {
//		return formError(BSERR00006DbInsertFailed, errStmt.Error(), "dbInsertFieldDefinition")
//	}
//	errClose := stmt.Close()
//	if errClose != nil {
//		return formError(BSERR00006DbInsertFailed, errClose.Error(), "dbInsertFieldDefinition")
//	}
//
//	return nil
//}
