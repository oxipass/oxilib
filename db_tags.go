package bslib

import "time"

const sqlCreateTableTags = `
	CREATE TABLE IF NOT EXISTS tags (
   		tag_id 		   	    INTEGER PRIMARY KEY AUTOINCREMENT,
		name         		VARCHAR NOT NULL,
		created  			DATETIME NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0'
	)
`

const sqlCreateTableItemsTags = `
	CREATE TABLE IF NOT EXISTS items_tags (
        it_id				INTEGER PRIMARY KEY AUTOINCREMENT,
		item_id 		    INT NOT NULL,
		tag_id				INT NOT NULL,
   		updated    			DATETIME NOT NULL,
		deleted				BOOLEAN NOT NULL CHECK (deleted IN (0,1)) default '0',
		FOREIGN KEY (item_id) REFERENCES items(item_id),
		FOREIGN KEY (tag_id) REFERENCES tags(tag_id)
	)
`

const sqlInsertTag = `
	INSERT INTO tags (name,created,updated,deleted) 
		VALUES (?,?,?,0)
`

func (sdb *storageDB) dbInsertTag(tagName string) (tagId int64, err error) {
	if sdb.sTX == nil {
		return 0, formError(BSERR00003DbTransactionFailed, "dbInsertTag")
	}

	creationTime := prepareTimeForDb(time.Now())

	stmt, err := sdb.sTX.Prepare(sqlInsertTag)
	if err != nil {
		return 0, formError(BSERR00006DbInsertFailed, err.Error(), "dbInsertTag")
	}
	res, errStmt := stmt.Exec(tagName,
		creationTime,
		creationTime)

	if errStmt != nil {
		return 0, formError(BSERR00006DbInsertFailed, errStmt.Error(), "dbInsertTag")
	}
	tagId, err = res.LastInsertId()
	if err != nil {
		return 0, formError(BSERR00006DbInsertFailed, err.Error(), "dbInsertTag")
	}
	errClose := stmt.Close()
	if errClose != nil {
		return 0, formError(BSERR00006DbInsertFailed, errClose.Error(), "dbInsertTag")
	}

	return tagId, nil
}

const sqlAssignTagToItem = `
	INSERT INTO items_tags (item_id, tag_id, updated, deleted) 
		VALUES (?,?,?,0)
`

func (sdb *storageDB) dbAssignTag(tagId int64, itemId int64) (itId int64, err error) {
	if sdb.sTX == nil {
		return 0, formError(BSERR00003DbTransactionFailed, "dbAssignTag")
	}

	creationTime := prepareTimeForDb(time.Now())

	stmt, err := sdb.sTX.Prepare(sqlAssignTagToItem)
	if err != nil {
		return 0, formError(BSERR00006DbInsertFailed, err.Error(), "dbAssignTag")
	}
	res, errStmt := stmt.Exec(itemId, tagId,
		creationTime)

	if errStmt != nil {
		return 0, formError(BSERR00006DbInsertFailed, errStmt.Error(), "dbAssignTag")
	}
	itId, err = res.LastInsertId()
	if err != nil {
		return 0, formError(BSERR00006DbInsertFailed, err.Error(), "dbAssignTag")
	}
	errClose := stmt.Close()
	if errClose != nil {
		return 0, formError(BSERR00006DbInsertFailed, errClose.Error(), "dbAssignTag")
	}

	return itId, nil
}

// sqlListItemTags - List all non-deleted items
const sqlListItemTags = `
	SELECT tags.tag_id, tags.name,  tags.created, tags.updated, tags.deleted
		FROM tags 
		INNER JOIN items_tags it on tags.tag_id = it.tag_id
		WHERE tags.deleted='0' 
		  AND it.deleted='0'
		  AND it.item_id=?
`

// dbSelectItemTags - select tags assigned to requested the item
func (sdb *storageDB) dbSelectItemTags(itemId int64) (tags []BSTag, err error) {

	rows, err := sdb.sDB.Query(sqlListItemTags, itemId)
	if err != nil {
		return tags, formError(BSERR00021FieldsReadFailed, err.Error())
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			err = formError(BSERR00021FieldsReadFailed, err.Error(), errClose.Error())
		}
	}()

	var bsTag BSTag

	for rows.Next() {
		err = rows.Scan(&bsTag.ID,
			&bsTag.Name,
			&bsTag.Created,
			&bsTag.Updated,
			&bsTag.Deleted)
		if err != nil {
			return tags, err
		}

		tags = append(tags, bsTag)
	}
	return tags, nil
}

// sqlListTags - List all available tags (excluding deleted)
const sqlListTags = `
	SELECT tag_id, name, created, updated, deleted
		FROM tags 
		WHERE tags.deleted='0' 
`

// dbSelectItemTags - select tags assigned to requested the item
func (sdb *storageDB) dbSelectTags() (tags []BSTag, err error) {

	rows, err := sdb.sDB.Query(sqlListTags)
	if err != nil {
		return tags, formError(BSERR00021FieldsReadFailed, err.Error())
	}
	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			err = formError(BSERR00021FieldsReadFailed, err.Error(), errClose.Error())
		}
	}()

	var bsTag BSTag

	for rows.Next() {
		err = rows.Scan(&bsTag.ID,
			&bsTag.Name,
			&bsTag.Created,
			&bsTag.Updated,
			&bsTag.Deleted)
		if err != nil {
			return tags, err
		}

		tags = append(tags, bsTag)
	}
	return tags, nil
}
