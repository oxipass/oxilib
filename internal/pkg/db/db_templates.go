package db

import (
	"github.com/oxipass/oxilib/internal/pkg/oxierr"
	"github.com/oxipass/oxilib/internal/pkg/utils"
	"github.com/oxipass/oxilib/models"
	"time"
)

const cSqlSelectItemTemplates = `
SELECT item_id, name, icon, created, updated, deleted
	FROM template_items
`

func (sdb *StorageDB) DbSelectItemTemplates(includeFields bool) (templates []models.OxiItemTemplate, err error) {

	rows, err := sdb.sDB.Query(cSqlSelectItemTemplates)
	if err != nil {
		return templates, oxierr.FormError(oxierr.BSERR00014ItemsReadFailed, err.Error())
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			err = oxierr.FormError(oxierr.BSERR00014ItemsReadFailed, err.Error(), errClose.Error())
		}
	}()

	var templateItem models.OxiItemTemplate

	for rows.Next() {
		err = rows.Scan(&templateItem.ID,
			&templateItem.Name,
			&templateItem.Icon,
			&templateItem.Created,
			&templateItem.Updated,
			&templateItem.Deleted)
		if err != nil {
			return templates, err
		}
		if includeFields {
			templateItem.Fields, err = sdb.DbSelectTemplateFields(templateItem.ID)
			if err != nil {
				err = oxierr.FormError(oxierr.BSERR00014ItemsReadFailed, err.Error())
			}
		}
		templates = append(templates, templateItem)
	}

	return templates, nil
}

const cSqlSelectItemTemplateFields = `
SELECT field_id, field_name, field_icon, value_type, created, updated, deleted
	FROM template_fields
	WHERE item_id = ?
`

func (sdb *StorageDB) DbSelectTemplateFields(itemId string) (fields []models.OxiFieldTemplate, err error) {

	stmtFields, err := sdb.sTX.Prepare(cSqlSelectItemTemplateFields)
	if err != nil {
		return fields, oxierr.FormError(oxierr.BSERR00028DbSelectFailed, err.Error(), "dbSelectTemplateFields:2")
	}

	rows, err := stmtFields.Query(itemId)
	if err != nil {
		return fields, oxierr.FormError(oxierr.BSERR00028DbSelectFailed, err.Error(), "dbSelectTemplateFields:3")
	}
	for rows.Next() {
		var field models.OxiFieldTemplate
		err = rows.Scan(&field.ID,
			&field.Name,
			&field.Icon,
			&field.ValueType,
			&field.Created,
			&field.Updated,
			&field.Deleted)
		if err != nil {
			return fields, oxierr.FormError(oxierr.BSERR00028DbSelectFailed, err.Error(), "dbSelectTemplateFields:4")
		}
		fields = append(fields, field)
	}
	errRows := rows.Close()
	if errRows != nil {
		return fields, oxierr.FormError(oxierr.BSERR00028DbSelectFailed, errRows.Error(), "dbSelectTemplateFields:5")
	}
	errClose := stmtFields.Close()
	if errClose != nil {
		return fields, oxierr.FormError(oxierr.BSERR00028DbSelectFailed, errClose.Error(), "dbSelectTemplateFields:6")
	}
	return fields, nil
}

const cSqlSelectItemTemplate = `
SELECT item_id, name, icon, created, updated, deleted
	FROM template_items
	WHERE item_id = ?
`

func (sdb *StorageDB) DbSelectItemTemplateById(itemId string) (models.OxiItemTemplate, error) {
	var item models.OxiItemTemplate
	//var field models.OxiFieldTemplate

	if sdb.sTX == nil {
		return item, oxierr.FormError(oxierr.BSERR00003DbTransactionFailed, "dbSelectItemTemplateById:1")
	}

	stmtItem, err := sdb.sTX.Prepare(cSqlSelectItemTemplate)
	if err != nil {
		return item, oxierr.FormError(oxierr.BSERR00028DbSelectFailed, err.Error(), "dbSelectItemTemplateById:2")
	}

	err = stmtItem.QueryRow(itemId).Scan(&item.ID,
		&item.Name,
		&item.Icon,
		&item.Created,
		&item.Updated,
		&item.Deleted)
	if err != nil {
		return item, oxierr.FormError(oxierr.BSERR00028DbSelectFailed, err.Error(), "dbSelectItemTemplateById:3")
	}
	errClose := stmtItem.Close()
	if errClose != nil {
		return item, oxierr.FormError(oxierr.BSERR00028DbSelectFailed, errClose.Error(), "dbSelectItemTemplateById:4")
	}
	item.Fields, err = sdb.DbSelectTemplateFields(item.ID)
	if err != nil {
		return item, oxierr.FormError(oxierr.BSERR00028DbSelectFailed, err.Error(), "dbSelectItemTemplateById:5")
	}
	/*
		stmtFields, err := sdb.sTX.Prepare(cSqlSelectItemTemplateFields)
		if err != nil {
			return item, oxierr.FormError(oxierr.BSERR00028DbSelectFailed, err.Error(), "dbSelectItemTemplateById:5")
		}

		rows, err := stmtFields.Query(itemId)
		if err != nil {
			return item, oxierr.FormError(oxierr.BSERR00028DbSelectFailed, err.Error(), "dbSelectItemTemplateById:6")
		}
		for rows.Next() {
			err = rows.Scan(&field.ID,
				&field.Name,
				&field.Icon,
				&field.ValueType,
				&field.Created,
				&field.Updated,
				&field.Deleted)
			if err != nil {
				return item, oxierr.FormError(oxierr.BSERR00028DbSelectFailed, err.Error(), "dbSelectItemTemplateById:7")
			}
			item.Fields = append(item.Fields, field)
		}
		errRows := rows.Close()
		if errRows != nil {
			return item, oxierr.FormError(oxierr.BSERR00028DbSelectFailed, errRows.Error(), "dbSelectItemTemplateById:8")
		}
		errClose = stmtFields.Close()
		if errClose != nil {
			return item, oxierr.FormError(oxierr.BSERR00028DbSelectFailed, errClose.Error(), "dbSelectItemTemplateById:9")
		}

	*/

	return item, nil
}

const cSqlInsertItemTemplate = `
	INSERT 
		INTO template_items (item_id, name,icon,created,updated,deleted) 
		VALUES (?,?,?,?,?,0)
`

func (sdb *StorageDB) DbInsertItemTemplate(itemId string, itemName string, itemIcon string) error {
	if sdb.sTX == nil {
		return oxierr.FormError(oxierr.BSERR00003DbTransactionFailed, "dbInsertItemTemplate")
	}

	creationTime := utils.PrepareTimeForDb(time.Now())

	stmt, err := sdb.sTX.Prepare(cSqlInsertItemTemplate)
	if err != nil {
		return oxierr.FormError(oxierr.BSERR00006DbInsertFailed, err.Error(), "dbInsertItemTemplate")
	}
	_, errStmt := stmt.Exec(itemId, itemName,
		itemIcon,
		creationTime,
		creationTime)

	if errStmt != nil {
		return oxierr.FormError(oxierr.BSERR00006DbInsertFailed, errStmt.Error(), "dbInsertItemTemplate")
	}

	errClose := stmt.Close()
	if errClose != nil {
		return oxierr.FormError(oxierr.BSERR00006DbInsertFailed, errClose.Error(), "dbInsertItemTemplate")
	}

	return nil
}

const cSqlInsertFieldTemplate = `
	REPLACE
		INTO template_fields (
			item_id,
		    field_id,
			field_icon,
			field_name,
			value_type,
			created,
			updated,
			deleted)
		VALUES (?,?,?,?,?,?,?,0)
`

func (sdb *StorageDB) DbInsertFieldTemplate(itemId string, fieldId string, field models.OxiField) error {

	if sdb.sTX == nil {
		return oxierr.FormError(oxierr.BSERR00003DbTransactionFailed, "dbInsertFieldTemplate:1")
	}

	creationTime := utils.PrepareTimeForDb(time.Now())

	stmt, err := sdb.sTX.Prepare(cSqlInsertFieldTemplate)
	if err != nil {
		return oxierr.FormError(oxierr.BSERR00006DbInsertFailed, err.Error(), "dbInsertFieldTemplate:2")
	}
	_, errStmt := stmt.Exec(itemId,
		fieldId,
		field.Icon,
		field.Name,
		field.ValueType,
		creationTime,
		creationTime)

	if errStmt != nil {
		return oxierr.FormError(oxierr.BSERR00006DbInsertFailed, errStmt.Error(), "dbInsertFieldTemplate:3")
	}
	errClose := stmt.Close()
	if errClose != nil {
		return oxierr.FormError(oxierr.BSERR00006DbInsertFailed, errClose.Error(), "dbInsertFieldTemplate:4")
	}

	return nil
}
