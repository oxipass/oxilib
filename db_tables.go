package bslib

func (sdb *storageDB) createTable(tableDefinition string) error {
	_, err := sdb.sTX.Exec(tableDefinition)
	if err != nil {
		return formError(BSERR00002DbTableCreationFailed, tableDefinition, err.Error())
	}
	return nil
}
