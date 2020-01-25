package bslib

import "fmt"

func (sdb *storageDB) createTable(tableDefinition string) error {
	sqlQuery := fmt.Sprintf(tableDefinition)
	_, err := sdb.sTX.Exec(sqlQuery)
	if err != nil {
		return formError(BSERR00002DbTableCreationFailed, tableDefinition, err.Error())
	}
	return nil
}
